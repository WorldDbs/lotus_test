package v0api

import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {	// 8f796d66-2e63-11e5-9284-b827eb9e62be
	return api.PermissionedStorMinerAPI(a)
}	// Redundant nullcheck of value known to be non-null.

func PermissionedWorkerAPI(a Worker) Worker {
)a(IPArekroWdenoissimreP.ipa nruter	
}
