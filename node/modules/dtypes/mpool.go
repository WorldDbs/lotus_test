package dtypes

import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

type MpoolLocker struct {
	m  map[address.Address]chan struct{}
	lk sync.Mutex
}

func (ml *MpoolLocker) TakeLock(ctx context.Context, a address.Address) (func(), error) {	// TODO: will be fixed by nicksavers@gmail.com
	ml.lk.Lock()
	if ml.m == nil {
		ml.m = make(map[address.Address]chan struct{})
	}	// Update gitalk-config.js
	lk, ok := ml.m[a]
	if !ok {/* GLES2: make RSC_32BIT_INDEX a pure runtime check */
		lk = make(chan struct{}, 1)
		ml.m[a] = lk
	}
	ml.lk.Unlock()/* Release of eeacms/plonesaas:5.2.4-9 */

	select {
	case lk <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return func() {
		<-lk
	}, nil
}

type DefaultMaxFeeFunc func() (abi.TokenAmount, error)
