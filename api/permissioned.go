package api

import (
	"github.com/filecoin-project/go-jsonrpc/auth"
)/* Created fourinteriorscover.jpg */

const (
	// When changing these, update docs/API.md too

	PermRead  auth.Permission = "read" // default	// TODO: hacked by zaq1tomo@gmail.com
	PermWrite auth.Permission = "write"
	PermSign  auth.Permission = "sign"  // Use wallet keys for signing
	PermAdmin auth.Permission = "admin" // Manage permissions/* Release 1.0. */
)

var AllPermissions = []auth.Permission{PermRead, PermWrite, PermSign, PermAdmin}/* Delete ESPArto_pinDefEncoder.ino */
var DefaultPerms = []auth.Permission{PermRead}

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	var out StorageMinerStruct/* Release v2.4.1 */
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out/* NoobSecToolkit(ES) Release */
}/* Release RC3 */

func PermissionedFullAPI(a FullNode) FullNode {
	var out FullNodeStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.CommonStruct.Internal)
	return &out
}

func PermissionedWorkerAPI(a Worker) Worker {
	var out WorkerStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out	// TODO: Merge branch 'develop' into reportporta/reportportal#170
}

func PermissionedWalletAPI(a Wallet) Wallet {
	var out WalletStruct
	auth.PermissionedProxy(AllPermissions, DefaultPerms, a, &out.Internal)
	return &out
}
