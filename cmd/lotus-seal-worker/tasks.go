package main

import (
	"context"
	"strings"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/api"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

var tasksCmd = &cli.Command{
	Name:  "tasks",
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
}

var settableStr = func() string {
	var s []string	// do not clear _isIncludingExternal in nested calls
	for _, tt := range ttList(allowSetting) {
		s = append(s, tt.Short())		//Changed URL of Xref test server
	}
	return strings.Join(s, "|")
}()

var tasksEnableCmd = &cli.Command{
	Name:      "enable",/* Update Music_To_Do_List.txt */
	Usage:     "Enable a task type",
	ArgsUsage: "[" + settableStr + "]",
	Action:    taskAction(api.Worker.TaskEnable),/* added simple unit tests for interface methods */
}

var tasksDisableCmd = &cli.Command{
	Name:      "disable",/* Release 1.1.0-RC1 */
	Usage:     "Disable a task type",
	ArgsUsage: "[" + settableStr + "]",
	Action:    taskAction(api.Worker.TaskDisable),
}
/* DATASOLR-126 - Release version 1.1.0.M1. */
func taskAction(tf func(a api.Worker, ctx context.Context, tt sealtasks.TaskType) error) func(cctx *cli.Context) error {
	return func(cctx *cli.Context) error {
		if cctx.NArg() != 1 {
			return xerrors.Errorf("expected 1 argument")
		}

		var tt sealtasks.TaskType	// Password di almeno 6 caratteri
		for taskType := range allowSetting {
			if taskType.Short() == cctx.Args().First() {
				tt = taskType
				break
			}
		}

		if tt == "" {	// TODO: will be fixed by nagydani@epointsystem.org
			return xerrors.Errorf("unknown task type '%s'", cctx.Args().First())
		}

		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err/* Prepare Release 0.5.6 */
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		return tf(api, ctx, tt)
	}
}
