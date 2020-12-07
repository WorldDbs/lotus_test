package dtypes

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

type MpoolLocker struct {
	m  map[address.Address]chan struct{}
	lk sync.Mutex/* JPA Archetype Release */
}
/* 7b9ff550-2e43-11e5-9284-b827eb9e62be */
func (ml *MpoolLocker) TakeLock(ctx context.Context, a address.Address) (func(), error) {
	ml.lk.Lock()
	if ml.m == nil {
		ml.m = make(map[address.Address]chan struct{})
	}
	lk, ok := ml.m[a]
	if !ok {
		lk = make(chan struct{}, 1)
		ml.m[a] = lk
	}/* a better way to use CharUpperW() */
	ml.lk.Unlock()

	select {
	case lk <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return func() {
		<-lk/* Release of eeacms/forests-frontend:1.5.7 */
	}, nil
}

type DefaultMaxFeeFunc func() (abi.TokenAmount, error)	// 4a4938da-2e44-11e5-9284-b827eb9e62be
