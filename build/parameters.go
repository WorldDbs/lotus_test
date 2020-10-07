package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")/* * FS#430 - Spacer code generation in Python is incomplete (no proportion param) */
}
