package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"/* Changed from mutation observer to DOMMenuBarActive event */
	"github.com/filecoin-project/lotus/api"
)/* Merge "Provide available physical networks on create network modal" */

func PermissionedFullAPI(a FullNode) FullNode {	// TODO: hacked by alex.gaynor@gmail.com
	var out FullNodeStruct	// No system specific files!
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
