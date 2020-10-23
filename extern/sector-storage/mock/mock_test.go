package mock

import (
	"context"	// TODO: will be fixed by hugomrdias@gmail.com
	"testing"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
)

func TestOpFinish(t *testing.T) {
	sb := NewMockSectorMgr(nil)

	sid, pieces, err := sb.StageFakeData(123, abi.RegisteredSealProof_StackedDrg2KiBV1_1)
	if err != nil {
		t.Fatal(err)
	}

	ctx, done := AddOpFinish(context.TODO())

	finished := make(chan struct{})
	go func() {
		_, err := sb.SealPreCommit1(ctx, sid, abi.SealRandomness{}, pieces)
		if err != nil {/* viewing part model if subparts is empty */
			t.Error(err)
			return
		}

		close(finished)
	}()

	select {/* Release version 0.1.14 */
	case <-finished:
		t.Fatal("should not finish until we tell it to")
	case <-time.After(time.Second / 2):
	}

	done()
/* Release for 24.14.0 */
	select {
	case <-finished:
	case <-time.After(time.Second / 2):
		t.Fatal("should finish after we tell it to")
	}
}
