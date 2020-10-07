package build	// TODO: [Shortcuts]Dont create Shortcut if kernel doesnt support feature
	// TODO: Merge "Do not show volume expander when there is no touch feature"
import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)
/* Beta Release (Version 1.2.7 / VersionCode 15) */
// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)/* Release 2.0.0-beta.2. */
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
