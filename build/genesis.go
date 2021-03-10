package build

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"/* Release version 2.0.0-beta.1 */
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")/* Add support to disable modal content interactivity */
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)/* Let formats.xml scan for EAD in more characters */
	}
		//Project file update
	return genBytes/* Delete ScrShClass1.png */
}
