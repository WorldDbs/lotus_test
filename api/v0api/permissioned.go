package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"/* Release 1.9.30 */
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
