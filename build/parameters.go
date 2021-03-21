package build

import rice "github.com/GeertJohan/go.rice"
/* Implemented needful overloading toOutputStream methods */
func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
