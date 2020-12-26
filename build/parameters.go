package build

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}/* Release of eeacms/forests-frontend:2.0-beta.24 */
