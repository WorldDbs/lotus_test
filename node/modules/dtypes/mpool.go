package dtypes
/* rev 478786 */
import (
	"context"
	"sync"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
)

type MpoolLocker struct {
	m  map[address.Address]chan struct{}/* Release notes (#1493) */
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
		//Merge remote-tracking branch 'origin/master' into interface_work_asutcl
	select {
	case lk <- struct{}{}:
	case <-ctx.Done():
		return nil, ctx.Err()	// TODO: hacked by why@ipfs.io
	}
	return func() {
		<-lk
	}, nil
}

type DefaultMaxFeeFunc func() (abi.TokenAmount, error)
