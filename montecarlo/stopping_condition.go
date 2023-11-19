package montecarlo

import (
	"sort"

	"github.com/domino14/macondo/stats"
	"github.com/domino14/macondo/tilemapping"
	"github.com/rs/zerolog/log"
)

const IterationsCutoff = 2000
const PerPlyStopScaling = 625
const SimilarPlaysIterationsCutoff = 750

const MinReasonableWProb = 0.005 // or 0.5%

// use stats to figure out when to stop simming.

func (s *Simmer) shouldStop(iterationCount uint64,
	playSimilarityCache map[string]bool) bool {
	// This function runs as the sim is ongoing. So we should be careful
	// what we do with memory here.
	plays := s.plays
	sc := s.stoppingCondition

	if len(plays) < 2 {
		return true
	}
	if int(iterationCount) > IterationsCutoff+(s.maxPlies*PerPlyStopScaling) {
		return true
	}
	// Otherwise, do some statistics.
	// shallow copy the array so we can sort it/play with it.
	c := make([]*SimmedPlay, len(plays))
	// count ignored plays
	ignoredPlays := 0
	bottomUnignoredWinPct := 0.0
	for i := range c {
		c[i] = plays[i]
		c[i].RLock()
		if c[i].ignore {
			ignoredPlays++
		}
		c[i].RUnlock()
	}
	if ignoredPlays >= len(c)-1 {
		// if there is only 1 unignored play, exit.
		return true
	}

	// sort copy by win pct.
	sort.Slice(c, func(i, j int) bool {
		c[i].RLock()
		c[j].RLock()
		defer c[j].RUnlock()
		defer c[i].RUnlock()
		if c[i].winPctStats.Mean() == c[j].winPctStats.Mean() {
			return c[i].equityStats.Mean() > c[j].equityStats.Mean()
		}
		return c[i].winPctStats.Mean() > c[j].winPctStats.Mean()
	})

	// find the bottom unignored win pct play
	for i := len(c) - 1; i >= 0; i-- {
		c[i].RLock()
		if !c[i].ignore {
			bottomUnignoredWinPct = c[i].winPctStats.Mean()
			c[i].RUnlock()
			break
		}
		c[i].RUnlock()
	}

	// we want to cut off plays that have no chance of winning.
	// assume the very top play is the winner, and then cut off plays that have
	// no chance of catching up.
	// "no chance" is of course defined by the stopping condition :)

	var ci float64
	switch sc {
	case Stop95:
		ci = stats.Z95
	case Stop98:
		ci = stats.Z98
	case Stop99:
		ci = stats.Z99
	}
	tiebreakByEquity := false
	tentativeWinner := c[0]
	tentativeWinner.RLock()
	μ := tentativeWinner.winPctStats.Mean()
	e := tentativeWinner.winPctStats.StandardError(ci)
	if μ <= MinReasonableWProb {
		// If the top play by win % has basically no win chance, tiebreak the whole
		// thing by equity.
		tiebreakByEquity = true
	} else if μ >= (1-MinReasonableWProb) && bottomUnignoredWinPct >= (1-MinReasonableWProb) {
		// If the top play by win % has basically no losing chance, check if the bottom
		// play also has no losing chance
		tiebreakByEquity = true
	}
	tentativeWinner.RUnlock()

	if tiebreakByEquity {
		// We may need to re-determine the tentative winner.
		highestEquity := -1000000.0
		highestEquityIdx := -1
		for idx, p := range c {
			p.RLock()
			eq := p.equityStats.Mean()
			if eq > highestEquity {
				highestEquityIdx = idx
				highestEquity = eq
			}
			p.RUnlock()
		}
		if highestEquityIdx != 0 {
			c[0], c[highestEquityIdx] = c[highestEquityIdx], c[0]
			tentativeWinner = c[0]
			log.Info().
				Str("old-tentative-winner", c[highestEquityIdx].play.ShortDescription()).
				Str("tentative-winner", tentativeWinner.play.ShortDescription()).
				Msg("tiebreaking by equity, re-determining tentative winner")
		}
		μ = tentativeWinner.equityStats.Mean()
		e = tentativeWinner.equityStats.StandardError(ci)
		log.Debug().Msg("stopping-condition-tiebreak-by-equity")
	}

	newIgnored := 0
	// assume standard normal distribution (?)
	for _, p := range c[1:] {
		p.RLock()
		if p.ignore {
			p.RUnlock()
			continue
		}
		μi := p.winPctStats.Mean()
		ei := p.winPctStats.StandardError(ci)
		if tiebreakByEquity {
			μi = p.equityStats.Mean()
			ei = p.equityStats.StandardError(ci)
		}
		p.RUnlock()
		if passTest(μ, e, μi, ei) {
			p.Ignore()
			newIgnored++
		} else if iterationCount > SimilarPlaysIterationsCutoff {
			if materiallySimilar(tentativeWinner, p, playSimilarityCache) {
				p.Ignore()
				newIgnored++
			}
		}
	}
	if newIgnored > 0 {
		log.Debug().Int("newIgnored", newIgnored).Msg("sim-cut-off")
	}
	if ignoredPlays+newIgnored >= len(c)-1 {
		// if there is only 1 unignored play, exit.
		return true
	}
	return false
}

// passTest: determine if a random variable X > Y with the given
// confidence level; return true if X > Y.
func passTest(μ, e, μi, ei float64) bool {
	// Z := zVal(μ, v, μi, vi)
	// X > Y if (μ - e) > (μi + ei)
	return (μ - e) > (μi + ei)
}

func materiallySimilar(p1, p2 *SimmedPlay, pcache map[string]bool) bool {

	p1ps := p1.play.ShortDescription()
	p2ps := p2.play.ShortDescription()
	if p1ps > p2ps {
		p1ps, p2ps = p2ps, p1ps
	}
	lookupstr := p1ps + "|" + p2ps
	if similar, ok := pcache[lookupstr]; ok {
		log.Trace().Str("lookupstr", lookupstr).
			Bool("similar", similar).
			Msg("in-similarity-cache")
		return similar
	}

	// two plays are "materially similar" if they use the same tiles and
	// start at the same square.
	p1r, p1c, p1v := p1.play.CoordsAndVertical()
	p2r, p2c, p2v := p2.play.CoordsAndVertical()

	if !(p1r == p2r && p1c == p2c && p1v == p2v) {
		pcache[lookupstr] = false
		return false
	}
	if p1.play.TilesPlayed() != p2.play.TilesPlayed() {
		pcache[lookupstr] = false
		return false
	}
	if len(p1.play.Tiles()) != len(p2.play.Tiles()) {
		pcache[lookupstr] = false
		return false
	}
	// these plays start at the same square and are the same length.
	// do they use the same tiles?
	a1 := make([]tilemapping.MachineLetter, len(p1.play.Tiles()))
	a2 := make([]tilemapping.MachineLetter, len(p2.play.Tiles()))
	copy(a1, p1.play.Tiles())
	copy(a2, p2.play.Tiles())
	sort.Slice(a1, func(i, j int) bool { return a1[i] < a1[j] })
	sort.Slice(a2, func(i, j int) bool { return a2[i] < a2[j] })
	for i := range a1 {
		if a1[i] != a2[i] {
			pcache[lookupstr] = false
			return false
		}
	}
	log.Debug().Str("lookupstr", lookupstr).Msg("materially-similar")
	pcache[lookupstr] = true
	return true
}

// func zVal(μ, v, μi, vi float64) float64 {
// 	// mean of X - Y = E(X-Y) = E(X) - E(Y)
// 	mean := μ - μi
// 	// variance of (X-Y) = V(X) + V(Y)
// 	variance := v + vi
// 	stdev := math.Sqrt(variance)
// 	// P(X > Y) = P(X - Y > 0)
// 	// let D = X - Y
// 	// then P(D > 0)
// 	// convert to standard normal variable (mean 0 stdev 1)
// 	// = P ((D - mean) / (stdev) > (0 - mean) / stdev)
// 	// then P(Z>(0 - mean)/stdev)
// 	// 95 percentile is Z 1.96
// 	// 99 percentile is Z 2.58
// 	return -mean / stdev
// }

// func zValStdev(μ, s, μi, si float64) float64 {
// 	return zVal(μ, s*s, μi, si*si)
// }
