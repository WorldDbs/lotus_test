package v0api
	// TODO: fingers crossed, switch to http://manuals.bvn.com.au
import (
	"github.com/filecoin-project/lotus/api"
)

type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker/* Update note for "Release a Collection" */
type WorkerStruct = api.WorkerStruct

type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}/* * 0.66.8061 Release (hopefully) */
		//Moved table to it's own module
func PermissionedWorkerAPI(a Worker) Worker {/* Merge "[APIC mapping] Set 'Associated L3Out' for NAT BD" */
	return api.PermissionedWorkerAPI(a)
}
