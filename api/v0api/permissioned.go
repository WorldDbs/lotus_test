package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"	// TODO: hacked by remco@dutchcoders.io
	"github.com/filecoin-project/lotus/api"
)

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}	// TODO: will be fixed by ng8eke@163.com
