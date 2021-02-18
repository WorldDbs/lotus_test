package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {	// TODO: 3e8a8700-2e63-11e5-9284-b827eb9e62be
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}	// TODO: d22e6010-2e51-11e5-9284-b827eb9e62be
