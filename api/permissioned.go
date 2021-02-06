package api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
)

const (/* Added Lab #4 */
	// When changing these, update docs/API.md too

	PermRead  auth.Permission = "read" // default	// TODO: Update error message for string types in _validate_iteratble
	PermWrite auth.Permission = "write"
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing
	PermAdmin auth.Permission = "admin" // Manage permissions
)

var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}
var DefaultPerms = []auth.Permission{PermRead}		//Update openpli.conf

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}/* czech top 1000 list */

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}

func PermissionedWorkerAPI(a Worker) Worker {
	var out WorkerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
/* Release v1.7 fix */
func PermissionedWalletAPI(a Wallet) Wallet {		//Fixing a lots of small problem.
	var out WalletStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out		//Remove ST stuff that is now in the new st package
}/* Add newline to OS X download */
