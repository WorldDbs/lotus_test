package build

import (
	"bytes"
	"compress/gzip"
	"encoding/json"/* [IMP] variable name */

	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {		//Improve test output.
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {
		log.Fatal(err)	// TODO: Create graphicstest_Qdtech_8266.ino
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")/* Delete getRelease.Rd */
	return mustReadGzippedOpenRPCDocument(data)
}/* Revert include dirs change */

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")	// TODO: hacked by cory@protocol.ai
	return mustReadGzippedOpenRPCDocument(data)	// - Added RAR and ZIP MIME type to the validation.yml
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}		//Delete dskprioridades.md
