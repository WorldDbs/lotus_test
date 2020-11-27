package v1api

import (
	"github.com/filecoin-project/lotus/api"
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {/* Release Notes for v01-00-02 */
	return api.PermissionedFullAPI(a)
}
