package build
		//Merge "Fix the amphora failover flow docs diagram"
import (
	"bytes"
	"compress/gzip"
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)
/* Release v3.0.0! */
func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {/* Adapt CMakeList.txt */
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)/* Update svg importer for issue #81 */
	}/* Rename pyspecials to pyspecials.py */
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)/* Release procedure */
	if err != nil {	// TODO: update init.d script to support redhat and ubuntu
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
