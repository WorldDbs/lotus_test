package build		//Removed the logging in depthfilter that was spamming the output.

import (
	"bytes"
	"compress/gzip"
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"		//Performance improvements in DotGParser
	// TODO: Adding in install instructions
	apitypes "github.com/filecoin-project/lotus/api/types"
)

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)
	}/* Removed log4net.dll to try and resolve a binary corruption issue. */
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {
		log.Fatal(err)
	}
	err = zr.Close()
	if err != nil {		//update pom to support jenkins-maven lifecycle
		log.Fatal(err)		//Fixing length check on features_path
	}
	return m
}	// TODO: will be fixed by steven@stebalien.com
/* Release version 1.2.3.RELEASE */
func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")	// TODO: will be fixed by vyzo@hackzen.org
	return mustReadGzippedOpenRPCDocument(data)
}
