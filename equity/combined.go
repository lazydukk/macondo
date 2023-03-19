package equity

import (
	"github.com/domino14/macondo/alphabet"
	"github.com/domino14/macondo/board"
	"github.com/domino14/macondo/cache"
	"github.com/domino14/macondo/config"
	"github.com/domino14/macondo/move"
	"github.com/rs/zerolog/log"
)

// CombinedStaticCalculator is a redundant struct that combines
// the function of several calculators. It is only here for speed purposes.
type CombinedStaticCalculator struct {
	leaveValues                Leaves
	preEndgameAdjustmentValues []float64
}

func NewCombinedStaticCalculator(lexiconName string,
	cfg *config.Config, leaveFilename, pegfile string) (
	*CombinedStaticCalculator, error) {

	calc := &CombinedStaticCalculator{}
	if leaveFilename == "" {
		leaveFilename = LeaveFilename
	}
	if pegfile == "" {
		pegfile = PEGAdjustmentFilename
	}
	leaves, err := cache.Load(cfg, "leavefile:"+lexiconName+":"+leaveFilename, LeaveCacheLoadFunc)
	if err != nil {
		log.Err(err).Msg("loading-leaves")
	}
	pegValues, err := cache.Load(cfg, "pegfile:"+lexiconName+":"+pegfile, PEGCacheLoadFunc)
	if err != nil {
		log.Err(err).Msg("loading-peg-values")
	}
	var ok bool
	calc.leaveValues, ok = leaves.(*OldLeaves)
	if !ok {
		log.Info().Msg("no leaves found, will use greedy strategy")
		calc.leaveValues = &BlankLeaves{}
	}
	calc.preEndgameAdjustmentValues, ok = pegValues.([]float64)
	if !ok {
		log.Info().Msg("no peg values found, will use no pre-endgame strategy")
		calc.preEndgameAdjustmentValues = []float64{}
	}
	return calc, nil
}

func (csc CombinedStaticCalculator) Equity(play *move.Move, board *board.GameBoard,
	bag *alphabet.Bag, oppRack *alphabet.Rack) float64 {

	leave := play.Leave()
	score := play.Score()

	leaveAdjustment := 0.0
	otherAdjustments := 0.0

	if board.IsEmpty() {
		otherAdjustments += placementAdjustment(play, board)
	}

	if bag.TilesRemaining() > 0 {
		leaveAdjustment = csc.leaveValues.LeaveValue(leave)
		bagPlusSeven := bag.TilesRemaining() - play.TilesPlayed() + 7
		if bagPlusSeven < len(csc.preEndgameAdjustmentValues) {
			preEndgameAdjustment := csc.preEndgameAdjustmentValues[bagPlusSeven]
			// log.Debug().Float64("peg-adjust", preEndgameAdjustment).Int("bagPlusSeven", bagPlusSeven).Msg("equity calc")
			otherAdjustments += preEndgameAdjustment
		}
	} else {
		// The bag is empty.
		otherAdjustments += endgameAdjustment(play, oppRack, bag.LetterDistribution())
	}

	return float64(score) + leaveAdjustment + otherAdjustments
}

func (csc CombinedStaticCalculator) LeaveValue(leave alphabet.MachineWord) float64 {
	return csc.leaveValues.LeaveValue(leave)
}