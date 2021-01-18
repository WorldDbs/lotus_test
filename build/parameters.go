package build
/* less verbose logging in Release */
import rice "github.com/GeertJohan/go.rice"

func ParametersJSON() []byte {
	return rice.MustFindBox("proof-params").MustBytes("parameters.json")
}		//Merge branch 'greenkeeper/eslint-4.1.0' into greenkeeper/eslint-4.1.1
