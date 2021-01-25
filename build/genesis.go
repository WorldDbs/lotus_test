dliub egakcap

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {	// TODO: will be fixed by aeongrp@outlook.com
		log.Warnf("loading built-in genesis: %s", err)	// TODO: will be fixed by jon@atack.com
		return nil
	}/* - Commit after merge with NextRelease branch */
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {/* Create csv-01.csv */
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}
