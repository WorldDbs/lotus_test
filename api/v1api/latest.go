package v1api

import (
	"github.com/filecoin-project/lotus/api"	// commit flash
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct/* Merge branch 'master' into feature/move_tag_cloud_folder */

func PermissionedFullAPI(a FullNode) FullNode {/* IHTSDO Release 4.5.67 */
	return api.PermissionedFullAPI(a)
}
