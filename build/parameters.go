package build
	// Merge "Fix lvm volume creation in fuel-agent"
import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")	// TODO: - updated docs
}/* Release 4.5.0 */
