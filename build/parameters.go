package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {/* Updated Leaflet 0 4 Released and 100 other files */
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")	// TODO: will be fixed by ligi@ligi.de
}
