package build
/* Release a8. */
import (
	"bytes"
	"compress/gzip"		//Updated the compas_slicer feedstock.
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"/* + some small changes in README.md */
/* render audio with fx pt 1 */
	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)/* adding in Release build */
	if err != nil {
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)
	}	// colocando comurs... descobri que falta os parameters
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {/* Create prelim.md */
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
	// Merge "Add -code-profile option."
func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)/* Release URL in change log */
}
