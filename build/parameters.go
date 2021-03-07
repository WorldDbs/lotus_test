package build/* Improved purging of container pages and added logging. */

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {	// Texts and images for the upcoming update (pending registrations)
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}
