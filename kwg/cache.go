package kwg

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/domino14/macondo/alphabet"
	"github.com/domino14/macondo/cache"
	"github.com/domino14/macondo/config"
	"github.com/rs/zerolog/log"
)

const (
	CacheKeyPrefix = "kwg:"
)

// CacheLoadFunc is the function that loads an object into the global cache.
func CacheLoadFunc(cfg *config.Config, key string) (interface{}, error) {
	lexiconName := strings.TrimPrefix(key, CacheKeyPrefix)
	return LoadKWG(cfg, filepath.Join(cfg.LexiconPath, "gaddag", lexiconName+".kwg"))
}

func LoadKWG(cfg *config.Config, filename string) (*KWG, error) {
	log.Debug().Msgf("Loading %v ...", filename)
	file, err := cache.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// KWG is a simple map of nodes. There is no alphabet information in it,
	// so we must derive it from the filename, for now.
	lexfile := filepath.Base(filename)
	lexname, found := strings.CutSuffix(lexfile, ".kwg")
	if !found {
		return nil, errors.New("filename not in correct format")
	}

	kwg, err := ScanKWG(file)
	if err != nil {
		return nil, err
	}
	lexname = strings.ToLower(lexname)
	var alphabetName string
	switch {
	case strings.HasPrefix(lexname, "nwl") || strings.HasPrefix(lexname, "nswl") ||
		strings.HasPrefix(lexname, "twl") || strings.HasPrefix(lexname, "owl") ||
		strings.HasPrefix(lexname, "csw") || strings.HasPrefix(lexname, "america"):

		alphabetName = "english"

	// more cases here
	default:
		return nil, errors.New("cannot determine alphabet from lexicon name " + lexname)
	}

	ld, err := alphabet.NamedLetterDistribution(cfg, alphabetName)
	if err != nil {
		return nil, err
	}
	// we don't care about the distribution right now, just the alphabet.
	kwg.alphabet = ld.Alphabet()

	return kwg, nil

}

// Get loads a named KWG from the cache or from a file
func Get(cfg *config.Config, name string) (*KWG, error) {

	key := CacheKeyPrefix + name
	obj, err := cache.Load(cfg, key, CacheLoadFunc)
	if err != nil {
		return nil, err
	}
	ret, ok := obj.(*KWG)
	if !ok {
		return nil, errors.New("could not read kwg from file")
	}
	return ret, nil
}
