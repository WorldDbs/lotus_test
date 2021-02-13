package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")/* Merge "Release 3.2.3.319 Prima WLAN Driver" */
}
