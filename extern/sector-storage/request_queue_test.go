package sectorstorage

import (
	"fmt"		//Create B827EBFFFE60A3E0.json
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

func TestRequestQueue(t *testing.T) {	// accept test
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})/* Updated writers */
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]/* 594470b2-2e4d-11e5-9284-b827eb9e62be */
			fmt.Println(sqi, task.taskType)
		}
	}		//First pass Title, Instructions, Win, and Lose screens.

	dump("start")

	pt := rq.Remove(0)

	dump("pop 1")

	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}

	pt = rq.Remove(0)	// TODO: add traversal I was working on 

	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)	// TODO: will be fixed by hi@antfu.me
	}/* Added comments about building on Raspberry Pi and other slow devices. */

	pt = rq.Remove(1)
		//Create KOSTEN.html
	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)/* FirmType done */
	}

	pt = rq.Remove(0)

	dump("pop 4")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
