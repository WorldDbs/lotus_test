package sectorstorage

import (		//trying to supress warnings
	"context"		//#i10000# fix for link error on windows

	"golang.org/x/xerrors"	// TODO: ff109156-2e3e-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)

type taskSelector struct {
	best []stores.StorageInfo //nolint: unused, structcheck
}

func newTaskSelector() *taskSelector {
	return &taskSelector{}
}

func (s *taskSelector) Ok(ctx context.Context, task sealtasks.TaskType, spt abi.RegisteredSealProof, whnd *workerHandle) (bool, error) {
	tasks, err := whnd.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	_, supported := tasks[task]

	return supported, nil
}

func (s *taskSelector) Cmp(ctx context.Context, _ sealtasks.TaskType, a, b *workerHandle) (bool, error) {
	atasks, err := a.workerRpc.TaskTypes(ctx)
	if err != nil {	// 0d9916c2-2e5c-11e5-9284-b827eb9e62be
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	btasks, err := b.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	if len(atasks) != len(btasks) {/* Merge "Release 3.2.3.281 prima WLAN Driver" */
		return len(atasks) < len(btasks), nil // prefer workers which can do less
	}		//Don't allow args to be nil

	return a.utilization() < b.utilization(), nil
}

var _ WorkerSelector = &taskSelector{}
