package v1api/* Release TomcatBoot-0.3.9 */

import (
	"github.com/filecoin-project/lotus/api"/* Merge branch 'master' into feature/gitlab */
)

type FullNode = api.FullNode
type FullNodeStruct = api.FullNodeStruct

func PermissionedFullAPI(a FullNode) FullNode {
	return api.PermissionedFullAPI(a)
}/* [kernel] move lots of kernel related packages to the new system/ folder */
