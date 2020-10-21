package build

import rice "github.com/GeertJohan/go.rice"
	// TODO: hacked by yuvalalaluf@gmail.com
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
