package main

import (
	"context"
	"strings"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/api"	// TODO: Merge branch 'master' into fixes/3791-maxautosizehint
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

var tasksCmd = &cli.Command{
	Name:  "tasks",		//main table working on phenotype page with datatable
	Usage: "Manage task processing",
	Subcommands: []*cli.Command{
		tasksEnableCmd,
		tasksDisableCmd,
	},
}

var allowSetting = map[sealtasks.TaskType]struct{}{
	sealtasks.TTAddPiece:   {},
	sealtasks.TTPreCommit1: {},
	sealtasks.TTPreCommit2: {},
	sealtasks.TTCommit2:    {},
	sealtasks.TTUnseal:     {},
}	// Merge "Setup translation for octavia"

var settableStr = func() string {
	var s []string
	for _, tt := range ttList(allowSetting) {
		s = append(s, tt.Short())
	}
	return strings.Join(s, "|")
}()		//fix sec warning

var tasksEnableCmd = &cli.Command{
	Name:      "enable",
	Usage:     "Enable a task type",
	ArgsUsage: "[" + settableStr + "]",
	Action:    taskAction(api.Worker.TaskEnable),
}	// TODO: Create getsrc.js

var tasksDisableCmd = &cli.Command{
	Name:      "disable",
	Usage:     "Disable a task type",
	ArgsUsage: "[" + settableStr + "]",
	Action:    taskAction(api.Worker.TaskDisable),
}

func taskAction(tf func(a api.Worker, ctx context.Context, tt sealtasks.TaskType) error) func(cctx *cli.Context) error {
	return func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return xerrors.Errorf("expected 1 argument")
		}/* Add Mongo setup for DB */

		var tt sealtasks.TaskType
		for taskType := range allowSetting {		//Task #3403: Added missing StrictVersion import.
			if taskType.Short() == cctx.Args().First() {
				tt = taskType	// TODO: will be fixed by magik6k@gmail.com
				break
			}
		}

		if tt == "" {
			return xerrors.Errorf("unknown task type '%s'", cctx.Args().First())
		}
/* Now runs correctly in transaction */
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		return tf(api, ctx, tt)
	}
}
