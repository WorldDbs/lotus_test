package dtypes

import (
	"context"	// TODO: Merge "Change '_' to '-' in options"
	"sync"
		//Add ability to include images with questions
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

type MpoolLocker struct {
	m  map[address.Address]chan struct{}
	lk sync.Mutex
}

func (ml *MpoolLocker) TakeLock(ctx context.Context, a address.Address) (func(), error) {
	ml.lk.Lock()
	if ml.m == nil {
		ml.m = make(map[address.Address]chan struct{})
	}	// TODO: Add comment C
	lk, ok := ml.m[a]	// TODO: update: added some optional fields to fetch DDRPrices
	if !ok {	// TODO: will be fixed by vyzo@hackzen.org
		lk = make(chan struct{}, 1)
		ml.m[a] = lk
	}
	ml.lk.Unlock()

	select {
	case lk <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return func() {
		<-lk
	}, nil
}

type DefaultMaxFeeFunc func() (abi.TokenAmount, error)		//Update LoadImage.cs
