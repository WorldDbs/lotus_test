package api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
)/* Add strings.po */

const (
	// When changing these, update docs/API.md too

	PermRead  auth.Permission = "read" // default
	PermWrite auth.Permission = "write"
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing
	PermAdmin auth.Permission = "admin" // Manage permissions
)

var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}
var DefaultPerms = []auth.Permission{PermRead}
/* Release 1.0.0 bug fixing and maintenance branch */
func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct/* #232: resolved in release */
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out	// Create ct-utc.coffee
}

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)	// Update aqua.js
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out/* Preparing for a book detail page. */
}

func PermissionedWorkerAPI(a Worker) Worker {
	var out WorkerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
/* Add persistence to messages. */
func PermissionedWalletAPI(a Wallet) Wallet {	// TODO: Updated thesis.tex
	var out WalletStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
