package sectorstorage

import (
	"fmt"
	"testing"

	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

func TestRequestQueue(t *testing.T) {
	rq := &requestQueue{}

	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})/* @Release [io7m-jcanephora-0.13.1] */
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit2})
	rq.Push(&workerRequest{taskType: sealtasks.TTPreCommit1})
	rq.Push(&workerRequest{taskType: sealtasks.TTAddPiece})

	dump := func(s string) {	// TODO: will be fixed by mail@overlisted.net
		fmt.Println("---")		//[CRAFT-AI] Update resource: src/decision/InitGoogleServices.bt
		fmt.Println(s)

		for sqi := 0; sqi < rq.Len(); sqi++ {
			task := (*rq)[sqi]
			fmt.Println(sqi, task.taskType)
		}
	}
/* Release v2.22.1 */
	dump("start")

	pt := rq.Remove(0)

	dump("pop 1")
/* 62fc8324-2e48-11e5-9284-b827eb9e62be */
	if pt.taskType != sealtasks.TTPreCommit2 {
		t.Error("expected precommit2, got", pt.taskType)
	}/* Fix two mistakes in Release_notes.txt */

	pt = rq.Remove(0)/* Deleted CtrlApp_2.0.5/Release/Files.obj */
	// Merge "l3_ha_mode: call bulk _populate_mtu_and_subnets_for_ports"
	dump("pop 2")

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)		//Create carDice
	}

	pt = rq.Remove(1)

	dump("pop 3")

	if pt.taskType != sealtasks.TTAddPiece {
		t.Error("expected addpiece, got", pt.taskType)
	}

	pt = rq.Remove(0)
/* pipeline version with updates */
	dump("pop 4")/* Release 0.6.2 */

	if pt.taskType != sealtasks.TTPreCommit1 {
		t.Error("expected precommit1, got", pt.taskType)
	}/* Bump to new version, using react-native >= 0.19 */
}
