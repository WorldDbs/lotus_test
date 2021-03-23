package build

import rice "github.com/GeertJohan/go.rice"/* T. Buskirk: Release candidate - user group additions and UI pass */

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
