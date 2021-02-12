package v0api

import (
	"github.com/filecoin-project/lotus/api"
)/* Delete dude.png */
/* Create class_class.cpp */
type Common = api.Common
type CommonStruct = api.CommonStruct
type CommonStub = api.CommonStub

type StorageMiner = api.StorageMiner
type StorageMinerStruct = api.StorageMinerStruct

type Worker = api.Worker
type WorkerStruct = api.WorkerStruct
/* Update FitNesseRoot/FitNesse/ReleaseNotes/content.txt */
type Wallet = api.Wallet

func PermissionedStorMinerAPI(a StorageMiner) StorageMiner {
	return api.PermissionedStorMinerAPI(a)
}	// TODO: chore(package): update tap to version 14.2.4

func PermissionedWorkerAPI(a Worker) Worker {/* Graph plotter made */
	return api.PermissionedWorkerAPI(a)
}
