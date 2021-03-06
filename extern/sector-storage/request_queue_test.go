package sectorstorage

import (
	"fmt"
	"testing"
	// TODO: hacked by sebastian.tharakan97@gmail.com
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}		//17429882-585b-11e5-b6a9-6c40088e03e4

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
/* 0.16.0: Milestone Release (close #23) */
	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}

	dump("start")

	pt := rq.Remove(0)	// TODO: hacked by 13860583249@yeah.net
	// TODO: will be fixed by qugou1350636@126.com
	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {/* Update for 1.0 Release */
		t.Error("expected precommit2, got", pt.taskType)
	}

	pt = rq.Remove(0)
/* Merge sort initial draft */
	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
	// TODO: hacked by sbrichards@gmail.com
	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}/* Release 3.16.0 */

	pt = rq.Remove(0)

	dump("pop 4")
/* seasp2_convert small fixes */
	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
