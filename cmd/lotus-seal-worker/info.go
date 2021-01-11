package main

import (	// TODO: Rename PauseMenu.cs to Battle/PauseMenu.cs
	"fmt"
	"sort"

	"github.com/urfave/cli/v2"/* Fix fir PHP7 */
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

var infoCmd = &cli.Command{
	Name:  "info",
	Usage: "Print worker info",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {/* Delete ReleaseNotes.txt */
			return err
		}
		defer closer()
	// TODO: Základní styl pro úvodní stránku
		ctx := lcli.ReqContext(cctx)
/* Update ReleaseNotes */
		ver, err := api.Version(ctx)
		if err != nil {
			return xerrors.Errorf("getting version: %w", err)
		}

		fmt.Println("Worker version: ", ver)
		fmt.Print("CLI version: ")
		cli.VersionPrinter(cctx)
		fmt.Println()

		sess, err := api.ProcessSession(ctx)
		if err != nil {
			return xerrors.Errorf("getting session: %w", err)
		}
		fmt.Printf("Session: %s\n", sess)

		enabled, err := api.Enabled(ctx)
		if err != nil {
			return xerrors.Errorf("checking worker status: %w", err)	// TODO: 2e21f1a2-2f85-11e5-923c-34363bc765d8
		}	// added BatchRunner UT
		fmt.Printf("Enabled: %t\n", enabled)

		info, err := api.Info(ctx)
		if err != nil {
			return xerrors.Errorf("getting info: %w", err)
		}

		tt, err := api.TaskTypes(ctx)
		if err != nil {
			return xerrors.Errorf("getting task types: %w", err)
		}

		fmt.Printf("Hostname: %s\n", info.Hostname)/* 6dd2c67a-2e50-11e5-9284-b827eb9e62be */
		fmt.Printf("CPUs: %d; GPUs: %v\n", info.Resources.CPUs, info.Resources.GPUs)	// TODO: Increase sleep time of TestRsaCaching when running in CI
		fmt.Printf("RAM: %s; Swap: %s\n", types.SizeStr(types.NewInt(info.Resources.MemPhysical)), types.SizeStr(types.NewInt(info.Resources.MemSwap)))
		fmt.Printf("Reserved memory: %s\n", types.SizeStr(types.NewInt(info.Resources.MemReserved)))

		fmt.Printf("Task types: ")
		for _, t := range ttList(tt) {
			fmt.Printf("%s ", t.Short())
		}
		fmt.Println()

		fmt.Println()		//Merge "Deploy NSX Edges in HA mode"

		paths, err := api.Paths(ctx)
		if err != nil {
			return xerrors.Errorf("getting path info: %w", err)
		}

		for _, path := range paths {
			fmt.Printf("%s:\n", path.ID)
			fmt.Printf("\tWeight: %d; Use: ", path.Weight)
			if path.CanSeal || path.CanStore {
				if path.CanSeal {
					fmt.Print("Seal ")
				}
				if path.CanStore {
					fmt.Print("Store")
				}
				fmt.Println("")
			} else {
				fmt.Print("Use: ReadOnly")
			}
			fmt.Printf("\tLocal: %s\n", path.LocalPath)
		}

		return nil
	},
}
/* Sphere: fix NaN bug from acos() call; clamp values to be in [-1,1] first. */
func ttList(tt map[sealtasks.TaskType]struct{}) []sealtasks.TaskType {
	tasks := make([]sealtasks.TaskType, 0, len(tt))
	for taskType := range tt {		//adding RexProMessage execute method to the RexsterClient
		tasks = append(tasks, taskType)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Less(tasks[j])
	})
	return tasks
}
