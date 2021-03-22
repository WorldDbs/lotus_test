package build

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
/* Automatic changelog generation for PR #40955 [ci skip] */
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
	if err != nil {		//msvc maintainance taks infinite amount of time
		log.Fatal(err)	// can never type that properly
	}
	err = zr.Close()
	if err != nil {		//piPews model for flyingPig
		log.Fatal(err)
	}
	return m
}	// TODO: Add another QA

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")	// add Seinfeld::Feed object for parsing the feeds
	return mustReadGzippedOpenRPCDocument(data)
}
		//Update sharing-buttons.html
func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}		//Use grizzly and jersey. builds and starts with trivial rest resource

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
