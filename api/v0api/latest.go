package v0api	// TODO: hacked by joshua@yottadb.com

import (		//Daogen generic ID field
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub	// Fixed few bugs related to delete meeting use cases.

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct
/* fix for GROOVY-2185 and GROOVY-2090 */
type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {	// TODO: hacked by jon@atack.com
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
