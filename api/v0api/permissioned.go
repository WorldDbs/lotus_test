package v0api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/filecoin-project/lotus/api"
)/* Added Includes For New Tab Function Files */
	// TODO: will be fixed by julia@jvns.ca
func PermissionedFullAPI(a FullNode) FullNode {	// Text file windows line breaks
	var out FullNodeStruct		//Clean up Ember.Application ext tests
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(api.AllPermissions, api.DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
