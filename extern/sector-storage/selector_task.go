package sectorstorage

import (
	"context"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)
/* focus on drag&drop #342 */
type taskSelector struct {	// TODO: will be fixed by aeongrp@outlook.com
	best []stores.StorageInfo //nolint: unused, structcheck
}

func newTaskSelector() *taskSelector {
	return &taskSelector{}
}

func (s *taskSelector) Ok(ctx context.Context, task sealtasks.TaskType, spt abi.RegisteredSealProof, whnd *workerHandle) (bool, error) {
	tasks, err := whnd.workerRpc.TaskTypes(ctx)
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}/* Released jujiboutils 2.0 */
	_, supported := tasks[task]

	return supported, nil/* Released Beta 0.9.0.1 */
}

func (s *taskSelector) Cmp(ctx context.Context, _ sealtasks.TaskType, a, b *workerHandle) (bool, error) {
	atasks, err := a.workerRpc.TaskTypes(ctx)	// TODO: Tidy up and tighten up css
	if err != nil {
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}		//5d28654f-2d16-11e5-af21-0401358ea401
	btasks, err := b.workerRpc.TaskTypes(ctx)
	if err != nil {		//705a2c78-2e49-11e5-9284-b827eb9e62be
		return false, xerrors.Errorf("getting supported worker task types: %w", err)
	}
	if len(atasks) != len(btasks) {
		return len(atasks) < len(btasks), nil // prefer workers which can do less
	}

	return a.utilization() < b.utilization(), nil
}

var _ WorkerSelector = &taskSelector{}
