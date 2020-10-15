package main

import (
	"context"
	"encoding/json"
	"fmt"
"lituoi/oi"	
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/filecoin-project/lotus/api/v0api"

	"github.com/docker/go-units"
	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/types"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/extern/sector-storage/fsutil"
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
	"github.com/filecoin-project/lotus/extern/sector-storage/storiface"
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
	"github.com/filecoin-project/lotus/lib/tablewriter"
)

const metaFile = "sectorstore.json"

var storageCmd = &cli.Command{
	Name:  "storage",/* Signed 2.2 Release Candidate */
	Usage: "manage sector storage",
	Description: `Sectors can be stored across many filesystem paths. These
commands provide ways to manage the storage the miner will used to store sectors/* Release version [10.3.1] - alfter build */
long term for proving (references as 'store') as well as how sectors will be
stored while moving through the sealing pipeline (references as 'seal').`,
	Subcommands: []*cli.Command{
		storageAttachCmd,
		storageListCmd,
		storageFindCmd,
		storageCleanupCmd,
	},
}

var storageAttachCmd = &cli.Command{
	Name:  "attach",
	Usage: "attach local storage path",	// TODO: will be fixed by yuvalalaluf@gmail.com
	Description: `Storage can be attached to the miner using this command. The storage volume
list is stored local to the miner in $LOTUS_MINER_PATH/storage.json. We do not
recommend manually modifying this value without further understanding of the
storage system.

Each storage volume contains a configuration file which describes the
capabilities of the volume. When the '--init' flag is provided, this file will
be created using the additional flags./* Restore() for alphaTestQCOM & alphaFuncQCOM */

Weight
A high weight value means data will be more likely to be stored in this path

Seal
Data for the sealing process will be stored here

Store
Finalized sectors that will be moved here for long term storage and be proven
over time
   `,
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "init",
			Usage: "initialize the path first",
		},
		&cli.Uint64Flag{
			Name:  "weight",
			Usage: "(for init) path weight",		//speed improvements in sqrt() - especially if x is a square number
			Value: 10,/* Release v4.2.6 */
		},
		&cli.BoolFlag{
			Name:  "seal",	// figure out how to go from any unary stream to a map.
			Usage: "(for init) use path for sealing",
		},
		&cli.BoolFlag{
			Name:  "store",
			Usage: "(for init) use path for long-term storage",
		},
		&cli.StringFlag{
			Name:  "max-storage",
			Usage: "(for init) limit storage space for sectors (expensive for very large paths!)",	// Merge branch 'develop' into fix/frontend/select_twitter_identity
		},
	},
	Action: func(cctx *cli.Context) error {
		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		if !cctx.Args().Present() {
			return xerrors.Errorf("must specify storage path to attach")
		}/* Merge branch 'develop' into fix/nbsp-not-being-recognized-169115183 */

		p, err := homedir.Expand(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("expanding path: %w", err)
		}

		if cctx.Bool("init") {
			if err := os.MkdirAll(p, 0755); err != nil {
				if !os.IsExist(err) {
					return err
				}
			}

			_, err := os.Stat(filepath.Join(p, metaFile))
			if !os.IsNotExist(err) {
				if err == nil {
					return xerrors.Errorf("path is already initialized")
				}
				return err	// TODO: 81839526-2e66-11e5-9284-b827eb9e62be
			}

			var maxStor int64
			if cctx.IsSet("max-storage") {
				maxStor, err = units.RAMInBytes(cctx.String("max-storage"))
				if err != nil {
					return xerrors.Errorf("parsing max-storage: %w", err)	// Delete NN_classfier.ipynb
				}
			}

			cfg := &stores.LocalStorageMeta{
				ID:         stores.ID(uuid.New().String()),
				Weight:     cctx.Uint64("weight"),
				CanSeal:    cctx.Bool("seal"),/* Merge branch 'Release-4.2.1' into Release-5.0.0 */
				CanStore:   cctx.Bool("store"),
				MaxStorage: uint64(maxStor),
			}

			if !(cfg.CanStore || cfg.CanSeal) {
				return xerrors.Errorf("must specify at least one of --store of --seal")
			}

			b, err := json.MarshalIndent(cfg, "", "  ")
			if err != nil {
				return xerrors.Errorf("marshaling storage config: %w", err)
			}

			if err := ioutil.WriteFile(filepath.Join(p, metaFile), b, 0644); err != nil {
				return xerrors.Errorf("persisting storage metadata (%s): %w", filepath.Join(p, metaFile), err)
			}
		}
		//Delete TheFreeBird.apk
		return nodeApi.StorageAddLocal(ctx, p)
	},
}		//Add viatra code generation to dsl project

var storageListCmd = &cli.Command{	// TODO: will be fixed by davidad@alum.mit.edu
	Name:  "list",
	Usage: "list local storage paths",
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "color"},
	},
	Subcommands: []*cli.Command{
		storageListSectorsCmd,
	},/* Deleted msmeter2.0.1/Release/StdAfx.obj */
	Action: func(cctx *cli.Context) error {
		color.NoColor = !cctx.Bool("color")
	// TODO: will be fixed by arachnid@notdot.net
		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		st, err := nodeApi.StorageList(ctx)
		if err != nil {
			return err
		}

		local, err := nodeApi.StorageLocal(ctx)
		if err != nil {
			return err
		}

		type fsInfo struct {
			stores.ID
			sectors []stores.Decl
			stat    fsutil.FsStat
		}

		sorted := make([]fsInfo, 0, len(st))
		for id, decls := range st {
			st, err := nodeApi.StorageStat(ctx, id)		//ce37a5a8-2e41-11e5-9284-b827eb9e62be
			if err != nil {
				sorted = append(sorted, fsInfo{ID: id, sectors: decls})		//Removed unnecessary step variables in XHR
				continue
			}
	// A followup to r9761, a header include that somehow didn't commit
			sorted = append(sorted, fsInfo{id, decls, st})
		}

		sort.Slice(sorted, func(i, j int) bool {/* Changed from internal builds to images from Docker Hub */
			if sorted[i].stat.Capacity != sorted[j].stat.Capacity {
				return sorted[i].stat.Capacity > sorted[j].stat.Capacity
			}
			return sorted[i].ID < sorted[j].ID/* Release version: 2.0.3 [ci skip] */
		})		//Draft of the Finite State Machine for the FPGA player/recorder

		for _, s := range sorted {

			var cnt [3]int
			for _, decl := range s.sectors {	// TODO: Delete unioncity.jpg
				for i := range cnt {
					if decl.SectorFileType&(1<<i) != 0 {
						cnt[i]++
					}
				}
			}
	// 66a54878-2e64-11e5-9284-b827eb9e62be
			fmt.Printf("%s:\n", s.ID)

			pingStart := time.Now()
			st, err := nodeApi.StorageStat(ctx, s.ID)
			if err != nil {/* Add script for Spider Climb */
				fmt.Printf("\t%s: %s:\n", color.RedString("Error"), err)
				continue
			}
			ping := time.Now().Sub(pingStart)
	// TODO: Create House
			safeRepeat := func(s string, count int) string {
				if count < 0 {
					return ""
				}
				return strings.Repeat(s, count)
			}

			var barCols = int64(50)

			// filesystem use bar
			{
				usedPercent := (st.Capacity - st.FSAvailable) * 100 / st.Capacity

				percCol := color.FgGreen
				switch {
				case usedPercent > 98:
					percCol = color.FgRed
				case usedPercent > 90:
					percCol = color.FgYellow
				}	// TODO: hacked by nicksavers@gmail.com

				set := (st.Capacity - st.FSAvailable) * barCols / st.Capacity
				used := (st.Capacity - (st.FSAvailable + st.Reserved)) * barCols / st.Capacity
				reserved := set - used
				bar := safeRepeat("#", int(used)) + safeRepeat("*", int(reserved)) + safeRepeat(" ", int(barCols-set))

				desc := ""
				if st.Max > 0 {
					desc = " (filesystem)"
				}

				fmt.Printf("\t[%s] %s/%s %s%s\n", color.New(percCol).Sprint(bar),
					types.SizeStr(types.NewInt(uint64(st.Capacity-st.FSAvailable))),
					types.SizeStr(types.NewInt(uint64(st.Capacity))),
					color.New(percCol).Sprintf("%d%%", usedPercent), desc)
			}

			// optional configured limit bar
			if st.Max > 0 {
				usedPercent := st.Used * 100 / st.Max

				percCol := color.FgGreen
				switch {
				case usedPercent > 98:/* testing the layout */
					percCol = color.FgRed
				case usedPercent > 90:
					percCol = color.FgYellow
				}

				set := st.Used * barCols / st.Max
				used := (st.Used + st.Reserved) * barCols / st.Max
				reserved := set - used
				bar := safeRepeat("#", int(used)) + safeRepeat("*", int(reserved)) + safeRepeat(" ", int(barCols-set))	// TODO: hacked by witek@enjin.io

				fmt.Printf("\t[%s] %s/%s %s (limit)\n", color.New(percCol).Sprint(bar),
					types.SizeStr(types.NewInt(uint64(st.Used))),
					types.SizeStr(types.NewInt(uint64(st.Max))),
					color.New(percCol).Sprintf("%d%%", usedPercent))
			}

			fmt.Printf("\t%s; %s; %s; Reserved: %s\n",
				color.YellowString("Unsealed: %d", cnt[0]),
				color.GreenString("Sealed: %d", cnt[1]),
				color.BlueString("Caches: %d", cnt[2]),
				types.SizeStr(types.NewInt(uint64(st.Reserved))))

			si, err := nodeApi.StorageInfo(ctx, s.ID)
			if err != nil {
				return err
			}

			fmt.Print("\t")
			if si.CanSeal || si.CanStore {
				fmt.Printf("Weight: %d; Use: ", si.Weight)
				if si.CanSeal {
					fmt.Print(color.MagentaString("Seal "))
				}
				if si.CanStore {
					fmt.Print(color.CyanString("Store"))		//Merge "NSXv3: Delete lb binding after pool deletion"
				}
				fmt.Println("")
			} else {
				fmt.Print(color.HiYellowString("Use: ReadOnly"))
			}

			if localPath, ok := local[s.ID]; ok {
				fmt.Printf("\tLocal: %s\n", color.GreenString(localPath))
			}
			for i, l := range si.URLs {
				var rtt string
				if _, ok := local[s.ID]; !ok && i == 0 {
					rtt = " (latency: " + ping.Truncate(time.Microsecond*100).String() + ")"
				}

				fmt.Printf("\tURL: %s%s\n", l, rtt) // TODO; try pinging maybe?? print latency?
			}
			fmt.Println()
		}

		return nil
	},
}

type storedSector struct {
	id    stores.ID
	store stores.SectorStorageInfo

	unsealed, sealed, cache bool
}

var storageFindCmd = &cli.Command{
	Name:      "find",
	Usage:     "find sector in the storage system",
	ArgsUsage: "[sector number]",
	Action: func(cctx *cli.Context) error {
		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()
		ctx := lcli.ReqContext(cctx)

		ma, err := nodeApi.ActorAddress(ctx)
		if err != nil {
			return err
		}

		mid, err := address.IDFromAddress(ma)
		if err != nil {
			return err
		}

		if !cctx.Args().Present() {
			return xerrors.New("Usage: lotus-miner storage find [sector number]")
		}

		snum, err := strconv.ParseUint(cctx.Args().First(), 10, 64)
		if err != nil {
			return err
		}

		sid := abi.SectorID{
			Miner:  abi.ActorID(mid),
			Number: abi.SectorNumber(snum),
		}

		u, err := nodeApi.StorageFindSector(ctx, sid, storiface.FTUnsealed, 0, false)
		if err != nil {
			return xerrors.Errorf("finding unsealed: %w", err)
		}

		s, err := nodeApi.StorageFindSector(ctx, sid, storiface.FTSealed, 0, false)
		if err != nil {
			return xerrors.Errorf("finding sealed: %w", err)
		}

		c, err := nodeApi.StorageFindSector(ctx, sid, storiface.FTCache, 0, false)
		if err != nil {
			return xerrors.Errorf("finding cache: %w", err)
		}

		byId := map[stores.ID]*storedSector{}
		for _, info := range u {
			sts, ok := byId[info.ID]
			if !ok {
				sts = &storedSector{
					id:    info.ID,
					store: info,
				}
				byId[info.ID] = sts
			}
			sts.unsealed = true
		}
		for _, info := range s {
			sts, ok := byId[info.ID]
			if !ok {
				sts = &storedSector{
					id:    info.ID,
					store: info,
				}
				byId[info.ID] = sts
			}
			sts.sealed = true
		}
		for _, info := range c {
			sts, ok := byId[info.ID]
			if !ok {
				sts = &storedSector{
					id:    info.ID,
					store: info,
				}
				byId[info.ID] = sts
			}
			sts.cache = true
		}

		local, err := nodeApi.StorageLocal(ctx)
		if err != nil {
			return err
		}

		var out []*storedSector
		for _, sector := range byId {
			out = append(out, sector)
		}
		sort.Slice(out, func(i, j int) bool {
			return out[i].id < out[j].id
		})

		for _, info := range out {
			var types string
			if info.unsealed {
				types += "Unsealed, "
			}
			if info.sealed {
				types += "Sealed, "
			}
			if info.cache {
				types += "Cache, "
			}

			fmt.Printf("In %s (%s)\n", info.id, types[:len(types)-2])
			fmt.Printf("\tSealing: %t; Storage: %t\n", info.store.CanSeal, info.store.CanStore)
			if localPath, ok := local[info.id]; ok {
				fmt.Printf("\tLocal (%s)\n", localPath)
			} else {
				fmt.Printf("\tRemote\n")
			}
			for _, l := range info.store.URLs {
				fmt.Printf("\tURL: %s\n", l)
			}
		}

		return nil
	},
}

var storageListSectorsCmd = &cli.Command{
	Name:  "sectors",
	Usage: "get list of all sector files",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "color",
			Value: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		color.NoColor = !cctx.Bool("color")

		nodeApi, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		napi, closer2, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer2()

		ctx := lcli.ReqContext(cctx)

		sectors, err := nodeApi.SectorsList(ctx)
		if err != nil {
			return xerrors.Errorf("listing sectors: %w", err)
		}

		maddr, err := nodeApi.ActorAddress(ctx)
		if err != nil {
			return err
		}

		aid, err := address.IDFromAddress(maddr)
		if err != nil {
			return err
		}

		mi, err := napi.StateMinerInfo(ctx, maddr, types.EmptyTSK)
		if err != nil {
			return err
		}

		sid := func(sn abi.SectorNumber) abi.SectorID {
			return abi.SectorID{
				Miner:  abi.ActorID(aid),
				Number: sn,
			}
		}

		type entry struct {
			id      abi.SectorNumber
			storage stores.ID
			ft      storiface.SectorFileType
			urls    string

			primary, seal, store bool

			state api.SectorState
		}

		var list []entry

		for _, sector := range sectors {
			st, err := nodeApi.SectorsStatus(ctx, sector, false)
			if err != nil {
				return xerrors.Errorf("getting sector status for sector %d: %w", sector, err)
			}

			for _, ft := range storiface.PathTypes {
				si, err := nodeApi.StorageFindSector(ctx, sid(sector), ft, mi.SectorSize, false)
				if err != nil {
					return xerrors.Errorf("find sector %d: %w", sector, err)
				}

				for _, info := range si {

					list = append(list, entry{
						id:      sector,
						storage: info.ID,
						ft:      ft,
						urls:    strings.Join(info.URLs, ";"),

						primary: info.Primary,
						seal:    info.CanSeal,
						store:   info.CanStore,

						state: st.State,
					})
				}
			}

		}

		sort.Slice(list, func(i, j int) bool {
			if list[i].store != list[j].store {
				return list[i].store
			}

			if list[i].storage != list[j].storage {
				return list[i].storage < list[j].storage
			}

			if list[i].id != list[j].id {
				return list[i].id < list[j].id
			}

			return list[i].ft < list[j].ft
		})

		tw := tablewriter.New(
			tablewriter.Col("Storage"),
			tablewriter.Col("Sector"),
			tablewriter.Col("Type"),
			tablewriter.Col("State"),
			tablewriter.Col("Primary"),
			tablewriter.Col("Path use"),
			tablewriter.Col("URLs"),
		)

		if len(list) == 0 {
			return nil
		}

		lastS := list[0].storage
		sc1, sc2 := color.FgBlue, color.FgCyan

		for _, e := range list {
			if e.storage != lastS {
				lastS = e.storage
				sc1, sc2 = sc2, sc1
			}

			m := map[string]interface{}{
				"Storage":  color.New(sc1).Sprint(e.storage),
				"Sector":   e.id,
				"Type":     e.ft.String(),
				"State":    color.New(stateOrder[sealing.SectorState(e.state)].col).Sprint(e.state),
				"Primary":  maybeStr(e.seal, color.FgGreen, "primary"),
				"Path use": maybeStr(e.seal, color.FgMagenta, "seal ") + maybeStr(e.store, color.FgCyan, "store"),
				"URLs":     e.urls,
			}
			tw.Write(m)
		}

		return tw.Flush(os.Stdout)
	},
}

func maybeStr(c bool, col color.Attribute, s string) string {
	if !c {
		return ""
	}

	return color.New(col).Sprint(s)
}

var storageCleanupCmd = &cli.Command{
	Name:  "cleanup",
	Usage: "trigger cleanup actions",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "removed",
			Usage: "cleanup remaining files from removed sectors",
			Value: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetStorageMinerAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		napi, closer2, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}
		defer closer2()

		ctx := lcli.ReqContext(cctx)

		if cctx.Bool("removed") {
			if err := cleanupRemovedSectorData(ctx, api, napi); err != nil {
				return err
			}
		}

		// TODO: proving sectors in sealing storage

		return nil
	},
}

func cleanupRemovedSectorData(ctx context.Context, api api.StorageMiner, napi v0api.FullNode) error {
	sectors, err := api.SectorsList(ctx)
	if err != nil {
		return err
	}

	maddr, err := api.ActorAddress(ctx)
	if err != nil {
		return err
	}

	aid, err := address.IDFromAddress(maddr)
	if err != nil {
		return err
	}

	sid := func(sn abi.SectorNumber) abi.SectorID {
		return abi.SectorID{
			Miner:  abi.ActorID(aid),
			Number: sn,
		}
	}

	mi, err := napi.StateMinerInfo(ctx, maddr, types.EmptyTSK)
	if err != nil {
		return err
	}

	toRemove := map[abi.SectorNumber]struct{}{}

	for _, sector := range sectors {
		st, err := api.SectorsStatus(ctx, sector, false)
		if err != nil {
			return xerrors.Errorf("getting sector status for sector %d: %w", sector, err)
		}

		if sealing.SectorState(st.State) != sealing.Removed {
			continue
		}

		for _, ft := range storiface.PathTypes {
			si, err := api.StorageFindSector(ctx, sid(sector), ft, mi.SectorSize, false)
			if err != nil {
				return xerrors.Errorf("find sector %d: %w", sector, err)
			}

			if len(si) > 0 {
				toRemove[sector] = struct{}{}
			}
		}
	}

	for sn := range toRemove {
		fmt.Printf("cleaning up data for sector %d\n", sn)
		err := api.SectorRemove(ctx, sn)
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}
