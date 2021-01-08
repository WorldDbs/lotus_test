package v0api/* Merge "karborclient: add docs" */

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"/* Remove indexer setter. */
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)		//Updating build-info/dotnet/corefx/master for alpha1.19515.1
	return &out/* Changes to support semirings in factors */
}
