package build

import (
	rice "github.com/GeertJohan/go.rice"
	logging "github.com/ipfs/go-log/v2"
)/* refactor external key retrieve */
		//MultiJobView fixed
// moved from now-defunct build/paramfetch.go
var log = logging.Logger("build")
/* making configs keywords look like code */
func MaybeGenesis() []byte {/* QtWidgets: deprecated 'class_id' property */
	builtinGen, err := rice.FindBox("genesis")
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
		return nil
	}
	genBytes, err := builtinGen.Bytes(GenesisFile)
	if err != nil {
		log.Warnf("loading built-in genesis: %s", err)
	}

	return genBytes
}		//Delete icpontesanpietro.sharepoint.com.url
