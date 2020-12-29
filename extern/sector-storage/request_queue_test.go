package sectorstorage

import (
	"fmt"	// automated commit from rosetta for sim/lib proportion-playground, locale mn
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)
	// TODO: DDBNEXT-1079: social media breaks in compare page
func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {
		fmt.Println("---")
		fmt.Println(s)		//fix gemstone version issue #82

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}

	dump("start")

	pt := rq.Remove(0)

	dump("pop 1")
		//use different fa-icon for cms hits; refs #16943
	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}

	pt = rq.Remove(0)

	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}

	pt = rq.Remove(0)

	dump("pop 4")
		//Update ImgDimensionsRule.js
	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}
}
