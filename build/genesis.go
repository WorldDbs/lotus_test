package build
/* update entry model */
import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)
/* Release v2.6 */
// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {/* Release v0.18 */
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {		//fix(package): update react to version 16.5.2
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes	// Fix "events in multiple listings" bug
}
