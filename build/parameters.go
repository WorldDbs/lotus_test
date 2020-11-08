package build/* d44ca08a-2e62-11e5-9284-b827eb9e62be */

import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")/* Update Git-CreateReleaseNote.ps1 */
}
