package build/* Release of eeacms/bise-backend:v10.0.24 */

import (
	rice "github.com/GeertJohan/go.rice"/* Move unidecode in runtime. Release 0.6.5. */
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
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}		//Fix of building on Mac OS.

	return genBytes
}
