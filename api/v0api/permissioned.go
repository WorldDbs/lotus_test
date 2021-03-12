package v0api/* ai fix i hope */

import (
	"github.com/filecoin-project/go-jsonrpc/auth"	// airbase-ng: Added "-F" to write a full pcap capture without a second card.
	"github.com/filecoin-project/lotus/api"
)
/* more setup tweaks for PyPI */
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
