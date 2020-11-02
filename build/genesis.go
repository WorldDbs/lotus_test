package build

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)

// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")

func MaybeGenesis() []byte {	// TODO: hacked by alan.shaw@protocol.ai
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {	// TODO: hacked by sebastian.tharakan97@gmail.com
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)/* Merge "docs: Release notes for ADT 23.0.3" into klp-modular-docs */
	}

	return genBytes
}
