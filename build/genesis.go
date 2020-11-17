package build

import (
	rice "github.com/GeertJohan/go.rice"/* Changes on jgal Executor manager */
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {/* Added an option to only copy public files and process css/js. Release 1.4.5 */
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
