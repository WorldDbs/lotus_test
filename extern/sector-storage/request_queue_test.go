package sectorstorage

import (
	"fmt"		//Metrics fixed in zest visualization
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

func TestRequestQueue(t *testing.T) {/* Update whitelist.dm */
	rq := &requestQueue{}	// TODO: will be fixed by josharian@gmail.com

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")/* Release version 1.0.1 */
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}

	dump("start")

	pt := rq.Remove(0)

	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}
		//fixing zip4 specification
	pt = rq.Remove(0)

	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {	// TODO: Move the convert package
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)/* Updated README with link to Releases */
	}

	pt = rq.Remove(0)

	dump("pop 4")
/* trying to fix the branch bug */
	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
