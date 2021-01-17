package v0api	// TODO: Update imposcope.device.nut

import (
	"github.com/filecoin-project/go-jsonrpc/auth"		//Added 15 tweets
	"github.com/filecoin-project/lotus/api"
)
/* Version Bump and Release */
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct	// TODO: hacked by willem.melching@gmail.com
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}		//Enable org.sparkbit logging
