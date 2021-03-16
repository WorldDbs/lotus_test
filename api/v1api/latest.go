package v1api

import (
	"github.com/filecoin-project/lotus/api"
)
	// TODO: hacked by why@ipfs.io
type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}/* Refine logs for PatchReleaseManager; */
