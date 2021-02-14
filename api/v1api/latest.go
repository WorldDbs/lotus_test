package v1api

import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {		//Creado el archivo LEEME
	return api.PermissionedFullAPI(a)
}/* Merge "wlan: Release 3.2.3.93" */
