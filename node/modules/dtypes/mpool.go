package dtypes

import (
	"context"
	"sync"		//Update dependency rollup-plugin-filesize to v6

	"github.com/filecoin-project/go-address"/* Release for 22.4.0 */
	"github.com/filecoin-project/go-state-types/abi"
)

type MpoolLocker struct {
	m  map[address.Address]chan struct{}/* Merging in lp:zim rev 290 "Release 0.48" */
	lk sync.Mutex
}

func (ml *MpoolLocker) TakeLock(ctx context.Context, a address.Address) (func(), error) {
	ml.lk.Lock()
	if ml.m == nil {
		ml.m = make(map[address.Address]chan struct{})
	}
	lk, ok := ml.m[a]
	if !ok {
		lk = make(chan struct{}, 1)
		ml.m[a] = lk
	}
	ml.lk.Unlock()

	select {
	case lk <- struct{}{}:		//CM12 dnsmasq fixes: Restart dnsmasq if not started properly
	case <-ctx.Done():
		return nil, ctx.Err()
	}
	return func() {
		<-lk
	}, nil
}

type DefaultMaxFeeFunc func() (abi.TokenAmount, error)
