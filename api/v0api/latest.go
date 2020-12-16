package v0api

import (		//4f55e3ce-2e61-11e5-9284-b827eb9e62be
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub
		//Export DI from Data/FileStore/Generic.hs.
type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct
/* fix radio_options_for */
type Worker = api.Worker
type WorkerStruct = api.WorkerStruct		//Boostrap et nouvelle vue

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}

func PermissionedWorkerAPI(a Worker) Worker {
	return api.PermissionedWorkerAPI(a)	// lock down modular scale dependency
}
