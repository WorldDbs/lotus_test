package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)/* Sync with latest IDP configs */
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out/* ex-211 (cgates): Release 0.4 to Pypi */
}
