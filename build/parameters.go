package build		//Merge "Fix Firewalls panel to override the right method"

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
