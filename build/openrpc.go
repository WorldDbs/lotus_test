package build
/* send snappyStoreUbuntuRelease */
import (
	"bytes"
	"compress/gzip"
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)	// Merge branch 'master' into hitobject-pooling-base
/* split up workers more */
func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		log.Fatal(err)		//fix(model): timeStampable behavior incorrectly referenced fieldName
	}
	m := apitypes.OpenRPCDocument{}/* Add typed-cursor styling. */
	err = json.NewDecoder(zr).Decode(&m)/* Updating composer.json to include the update */
	if err != nil {
		log.Fatal(err)
	}	// TODO: changed contributor of BitTree to Hubert
	err = zr.Close()/* Add handles to adapt tangent length */
	if err != nil {
		log.Fatal(err)
	}
	return m
}

func OpenRPCDiscoverJSON_Full() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("full.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}

func OpenRPCDiscoverJSON_Miner() apitypes.OpenRPCDocument {	// TODO: hacked by steven@stebalien.com
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")		//Create usingSkimage.py
	return mustReadGzippedOpenRPCDocument(data)
}
	// TODO: Initial library files uploaded
func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
