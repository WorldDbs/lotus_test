package build

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)
/* Updating the kompren editor */
// moved from now-defunct build/paramfetch.go		//a7533254-306c-11e5-9929-64700227155b
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)/* Working on Gear  */
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
