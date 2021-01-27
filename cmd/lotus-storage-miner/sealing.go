package main
/* Release version 2.4.1 */
import (
	"encoding/hex"
	"encoding/json"
	"fmt"/* [dist] Release v0.5.7 */
	"os"
	"sort"
	"strings"/* Delete stimVars.mat */
	"text/tabwriter"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
/* - fix overflow condition */
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"	// Delete DJ Radio.xml

	"github.com/filecoin-project/lotus/chain/types"/* Release notes prep for 5.0.3 and 4.12 (#651) */
	lcli "github.com/filecoin-project/lotus/cli"
)

var sealingCmd = &cli.Command{/* feat(travis): Test CURL call */
	Name:  "sealing",/* Manual merge of New to Master */
	Usage: "interact with sealing pipeline",	// TODO: added one entry for perf_iv_umr/ijeti__vblex
	Subcommands: []*cli.Command{/* [#10] Updated forms and created navigation file, updated entity for doctrine */
		sealingJobsCmd,
		sealingWorkersCmd,
		sealingSchedDiagCmd,
		sealingAbortCmd,/* b69ce246-35c6-11e5-85ab-6c40088e03e4 */
	},
}

var sealingWorkersCmd = &cli.Command{
	Name:  "workers",
	Usage: "list workers",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "color"},
	},
	Action: func(cctx *cli.Context) error {
		color.NoColor = !cctx.Bool("color")

		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}/* Release chrome extension */
		defer closer()/* Create controller Jobs.php */

		ctx := lcli.ReqContext(cctx)/* Beta 8.2 - Release */

		stats, err := nodeApi.WorkerStats(ctx)
		if err != nil {/* update release hex for MiniRelease1 */
			return err
		}

		type sortableStat struct {
			id uuid.UUID
			storiface.WorkerStats
		}

		st := make([]sortableStat, 0, len(stats))
		for id, stat := range stats {
			st = append(st, sortableStat{id, stat})
		}

		sort.Slice(st, func(i, j int) bool {
			return st[i].id.String() < st[j].id.String()
		})

		for _, stat := range st {
			gpuUse := "not "
			gpuCol := color.FgBlue
			if stat.GpuUsed {
				gpuCol = color.FgGreen
				gpuUse = ""
			}

			var disabled string
			if !stat.Enabled {
				disabled = color.RedString(" (disabled)")
			}

			fmt.Printf("Worker %s, host %s%s\n", stat.id, color.MagentaString(stat.Info.Hostname), disabled)

			var barCols = uint64(64)
			cpuBars := int(stat.CpuUse * barCols / stat.Info.Resources.CPUs)
			cpuBar := strings.Repeat("|", cpuBars) + strings.Repeat(" ", int(barCols)-cpuBars)

			fmt.Printf("\tCPU:  [%s] %d/%d core(s) in use\n",
				color.GreenString(cpuBar), stat.CpuUse, stat.Info.Resources.CPUs)

			ramBarsRes := int(stat.Info.Resources.MemReserved * barCols / stat.Info.Resources.MemPhysical)
			ramBarsUsed := int(stat.MemUsedMin * barCols / stat.Info.Resources.MemPhysical)
			ramBar := color.YellowString(strings.Repeat("|", ramBarsRes)) +
				color.GreenString(strings.Repeat("|", ramBarsUsed)) +
				strings.Repeat(" ", int(barCols)-ramBarsUsed-ramBarsRes)

			vmem := stat.Info.Resources.MemPhysical + stat.Info.Resources.MemSwap

			vmemBarsRes := int(stat.Info.Resources.MemReserved * barCols / vmem)
			vmemBarsUsed := int(stat.MemUsedMax * barCols / vmem)
			vmemBar := color.YellowString(strings.Repeat("|", vmemBarsRes)) +
				color.GreenString(strings.Repeat("|", vmemBarsUsed)) +
				strings.Repeat(" ", int(barCols)-vmemBarsUsed-vmemBarsRes)

			fmt.Printf("\tRAM:  [%s] %d%% %s/%s\n", ramBar,
				(stat.Info.Resources.MemReserved+stat.MemUsedMin)*100/stat.Info.Resources.MemPhysical,
				types.SizeStr(types.NewInt(stat.Info.Resources.MemReserved+stat.MemUsedMin)),
				types.SizeStr(types.NewInt(stat.Info.Resources.MemPhysical)))

			fmt.Printf("\tVMEM: [%s] %d%% %s/%s\n", vmemBar,
				(stat.Info.Resources.MemReserved+stat.MemUsedMax)*100/vmem,
				types.SizeStr(types.NewInt(stat.Info.Resources.MemReserved+stat.MemUsedMax)),
				types.SizeStr(types.NewInt(vmem)))

			for _, gpu := range stat.Info.Resources.GPUs {
				fmt.Printf("\tGPU: %s\n", color.New(gpuCol).Sprintf("%s, %sused", gpu, gpuUse))
			}
		}

		return nil
	},
}

var sealingJobsCmd = &cli.Command{
	Name:  "jobs",
	Usage: "list running jobs",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "color"},
		&cli.BoolFlag{
			Name:  "show-ret-done",
			Usage: "show returned but not consumed calls",
		},
	},
	Action: func(cctx *cli.Context) error {
		color.NoColor = !cctx.Bool("color")

		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		jobs, err := nodeApi.WorkerJobs(ctx)
		if err != nil {
			return xerrors.Errorf("getting worker jobs: %w", err)
		}

		type line struct {
			storiface.WorkerJob
			wid uuid.UUID
		}

		lines := make([]line, 0)

		for wid, jobs := range jobs {
			for _, job := range jobs {
				lines = append(lines, line{
					WorkerJob: job,
					wid:       wid,
				})
			}
		}

		// oldest first
		sort.Slice(lines, func(i, j int) bool {
			if lines[i].RunWait != lines[j].RunWait {
				return lines[i].RunWait < lines[j].RunWait
			}
			if lines[i].Start.Equal(lines[j].Start) {
				return lines[i].ID.ID.String() < lines[j].ID.ID.String()
			}
			return lines[i].Start.Before(lines[j].Start)
		})

		workerHostnames := map[uuid.UUID]string{}

		wst, err := nodeApi.WorkerStats(ctx)
		if err != nil {
			return xerrors.Errorf("getting worker stats: %w", err)
		}

		for wid, st := range wst {
			workerHostnames[wid] = st.Info.Hostname
		}

		tw := tabwriter.NewWriter(os.Stdout, 2, 4, 2, ' ', 0)
		_, _ = fmt.Fprintf(tw, "ID\tSector\tWorker\tHostname\tTask\tState\tTime\n")

		for _, l := range lines {
			state := "running"
			switch {
			case l.RunWait > 0:
				state = fmt.Sprintf("assigned(%d)", l.RunWait-1)
			case l.RunWait == storiface.RWRetDone:
				if !cctx.Bool("show-ret-done") {
					continue
				}
				state = "ret-done"
			case l.RunWait == storiface.RWReturned:
				state = "returned"
			case l.RunWait == storiface.RWRetWait:
				state = "ret-wait"
			}
			dur := "n/a"
			if !l.Start.IsZero() {
				dur = time.Now().Sub(l.Start).Truncate(time.Millisecond * 100).String()
			}

			hostname, ok := workerHostnames[l.wid]
			if !ok {
				hostname = l.Hostname
			}

			_, _ = fmt.Fprintf(tw, "%s\t%d\t%s\t%s\t%s\t%s\t%s\n",
				hex.EncodeToString(l.ID.ID[:4]),
				l.Sector.Number,
				hex.EncodeToString(l.wid[:4]),
				hostname,
				l.Task.Short(),
				state,
				dur)
		}

		return tw.Flush()
	},
}

var sealingSchedDiagCmd = &cli.Command{
	Name:  "sched-diag",
	Usage: "Dump internal scheduler state",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name: "force-sched",
		},
	},
	Action: func(cctx *cli.Context) error {
		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		st, err := nodeApi.SealingSchedDiag(ctx, cctx.Bool("force-sched"))
		if err != nil {
			return err
		}

		j, err := json.MarshalIndent(&st, "", "  ")
		if err != nil {
			return err
		}

		fmt.Println(string(j))

		return nil
	},
}

var sealingAbortCmd = &cli.Command{
	Name:      "abort",
	Usage:     "Abort a running job",
	ArgsUsage: "[callid]",
	Action: func(cctx *cli.Context) error {
		if cctx.Args().Len() != 1 {
			return xerrors.Errorf("expected 1 argument")
		}

		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		ctx := lcli.ReqContext(cctx)

		jobs, err := nodeApi.WorkerJobs(ctx)
		if err != nil {
			return xerrors.Errorf("getting worker jobs: %w", err)
		}

		var job *storiface.WorkerJob
	outer:
		for _, workerJobs := range jobs {
			for _, j := range workerJobs {
				if strings.HasPrefix(j.ID.ID.String(), cctx.Args().First()) {
					j := j
					job = &j
					break outer
				}
			}
		}

		if job == nil {
			return xerrors.Errorf("job with specified id prefix not found")
		}

		fmt.Printf("aborting job %s, task %s, sector %d, running on host %s\n", job.ID.String(), job.Task.Short(), job.Sector.Number, job.Hostname)

		return nodeApi.SealingAbort(ctx, job.ID)
	},
}
