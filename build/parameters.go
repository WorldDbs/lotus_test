package build

import rice "github.com/GeertJohan/go.rice"	// TODO: hacked by alan.shaw@protocol.ai

func ParametersJSON() []byte {/* Release the kraken! :octopus: */
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
