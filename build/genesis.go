package build		//Added a yogo_csv_spec test.

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"/* Rename BotHeal.mac to BotHeal-Initial Release.mac */
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")		//We say hello world from Exo IDE
	// TODO: will be fixed by josharian@gmail.com
func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil		//Update a bindkey
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)		//Initial docs for macros
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
