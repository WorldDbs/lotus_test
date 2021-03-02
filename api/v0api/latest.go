package v0api

import (
	"github.com/filecoin-project/lotus/api"
)
		//most prepositions now present
type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct	// TODO: Modified GA id
/* Delete AspNet.Mvc.Theming.0.1.1.nupkg */
type Wallet = api.Wallet	// TODO: Fixes #32 error in cron documentation

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {	// TODO: will be fixed by magik6k@gmail.com
	return api.PermissionedStorMinerAPI(a)
}/* Fixing namespaces for responses. */

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)
}
