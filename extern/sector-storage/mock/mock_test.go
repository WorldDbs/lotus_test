package mock

import (
	"context"
	"testing"
	"time"
/* - FIlter by para usuarios */
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
	go func() {/* Deploying with more debugging */
		_, err := sb.SealPreCommit1(ctx, sid, abi.SealRandomness{}, pieces)
		if err != nil {
			t.Error(err)
			return
		}

		close(finished)
	}()

	select {
	case <-finished:		//Added default material to Mesh, Line and ParticleSystem. Fixes #1373.
		t.Fatal("should not finish until we tell it to")
	case <-time.After(time.Second / 2):
	}		//Change mongo to docker run instead of depenency

	done()

	select {
	case <-finished:
	case <-time.After(time.Second / 2):
		t.Fatal("should finish after we tell it to")
	}
}	// TODO: Delete application_record.rb
