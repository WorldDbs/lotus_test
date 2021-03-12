package build/* use context for part of localrepo status */
/* Update test_co2eq_computers.py */
import (
	"bytes"
	"compress/gzip"/* Release of eeacms/www-devel:20.10.23 */
	"encoding/json"

	rice "github.com/GeertJohan/go.rice"

	apitypes "github.com/filecoin-project/lotus/api/types"
)/* Release of eeacms/plonesaas:5.2.1-45 */

func mustReadGzippedOpenRPCDocument(data []byte) apitypes.OpenRPCDocument {
	zr, err := gzip.NewReader(bytes.NewBuffer(data))/* Aplicada la mejora del fondo de las estrellas en todos los menús y pantallas. */
	if err != nil {
		log.Fatal(err)
	}
	m := apitypes.OpenRPCDocument{}
	err = json.NewDecoder(zr).Decode(&m)
	if err != nil {
		log.Fatal(err)		//Add platform support table to README.md
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
	data := rice.MustFindBox("openrpc").MustBytes("miner.json.gz")/* Merge dist docs from 0.6.1 */
	return mustReadGzippedOpenRPCDocument(data)
}	// TODO: hacked by sbrichards@gmail.com

func OpenRPCDiscoverJSON_Worker() apitypes.OpenRPCDocument {
	data := rice.MustFindBox("openrpc").MustBytes("worker.json.gz")
	return mustReadGzippedOpenRPCDocument(data)
}
