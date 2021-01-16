package build
/* add version_requirement to dependencies */
import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}/* 8d28d394-2e5e-11e5-9284-b827eb9e62be */
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {/* Released v2.1.3 */
		log.Warnf("loading built-in genesis: %s", err)
	}/* Release 1.0.3 for Bukkit 1.5.2-R0.1 and ByteCart 1.5.0 */

	return genBytes
}
