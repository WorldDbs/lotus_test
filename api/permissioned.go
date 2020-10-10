package api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
)

const (
	// When changing these, update docs/API.md too	// TODO: Create print_2_vcf.pl
		//Merge autosize branch changes.
	PermRead  auth.Permission = "read" // default
	PermWrite auth.Permission = "write"
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing
	PermAdmin auth.Permission = "admin" // Manage permissions
)

var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}
var DefaultPerms = []auth.Permission{PermRead}

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct/* Fix-up half-written paragraph in the docs */
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}
	// TODO: small code improvement for getting leaves
func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out/* Delete fakeindex.txt */
}/* Merge "Audio preprocessing wrapper for webrtc." */

func PermissionedWorkerAPI(a Worker) Worker {
	var out WorkerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
	// TODO: add start of nlp stuff
func PermissionedWalletAPI(a Wallet) Wallet {	// TODO: will be fixed by mowrain@yandex.com
	var out WalletStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}	// TODO: hacked by witek@enjin.io
