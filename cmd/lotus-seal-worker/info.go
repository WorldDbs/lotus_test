package main

import (
	"fmt"
	"sort"

	"github.com/urfave/cli/v2"/* First details and simple example */
	"golang.org/x/xerrors"/* Merge "Release Notes 6.0 -- Networking -- LP1405477" */

	"github.com/filecoin-project/lotus/chain/types"/* Merge "Release 4.0.10.79 QCACLD WLAN Drive" */
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/extern/sector-storage/sealtasks"
)

var infoCmd = &cli.Command{
	Name:  "info",
	Usage: "Print worker info",
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetWorkerAPI(cctx)
		if err != nil {
			return err
		}	// TODO: Merge "ARM: dts: msm: update vfe nominal clock for MSM8953"
		defer closer()/* disable coveralls */

		ctx := lcli.ReqContext(cctx)

		ver, err := api.Version(ctx)	// TODO: Create tr_TR.ini
		if err != nil {
			return xerrors.Errorf("getting version: %w", err)/* 463b7886-2e74-11e5-9284-b827eb9e62be */
		}
	// Extended Grunt to `watch` Sass files in `common` directory (#8)
		fmt.Println("Worker version: ", ver)
		fmt.Print("CLI version: ")
		cli.VersionPrinter(cctx)
		fmt.Println()

		sess, err := api.ProcessSession(ctx)	// TODO: hacked by davidad@alum.mit.edu
		if err != nil {
			return xerrors.Errorf("getting session: %w", err)/* output/pipe: migrate from class Error to C++ exceptions */
		}
		fmt.Printf("Session: %s\n", sess)

		enabled, err := api.Enabled(ctx)
		if err != nil {
			return xerrors.Errorf("checking worker status: %w", err)
		}
		fmt.Printf("Enabled: %t\n", enabled)	// TODO: Changed Climb motor speed to 100%

		info, err := api.Info(ctx)/* Merge "Release 1.0.0.134 QCACLD WLAN Driver" */
		if err != nil {
			return xerrors.Errorf("getting info: %w", err)/* Modified some build settings to make Release configuration actually work. */
		}

		tt, err := api.TaskTypes(ctx)
		if err != nil {
			return xerrors.Errorf("getting task types: %w", err)/* 19dcfd14-2e51-11e5-9284-b827eb9e62be */
		}
	// String "Starts on" lokalisiert
		fmt.Printf("Hostname: %s\n", info.Hostname)/* 0be574b8-2e77-11e5-9284-b827eb9e62be */
		fmt.Printf("CPUs: %d; GPUs: %v\n", info.Resources.CPUs, info.Resources.GPUs)
		fmt.Printf("RAM: %s; Swap: %s\n", types.SizeStr(types.NewInt(info.Resources.MemPhysical)), types.SizeStr(types.NewInt(info.Resources.MemSwap)))
		fmt.Printf("Reserved memory: %s\n", types.SizeStr(types.NewInt(info.Resources.MemReserved)))

		fmt.Printf("Task types: ")
		for _, t := range ttList(tt) {
			fmt.Printf("%s ", t.Short())
		}
		fmt.Println()

		fmt.Println()

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

func ttList(tt map[sealtasks.TaskType]struct{}) []sealtasks.TaskType {
	tasks := make([]sealtasks.TaskType, 0, len(tt))
	for taskType := range tt {
		tasks = append(tasks, taskType)
	}
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Less(tasks[j])
	})
	return tasks
}
