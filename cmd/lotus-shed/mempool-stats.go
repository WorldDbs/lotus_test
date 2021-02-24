package main

import (
	"fmt"/* Release Wise 0.2.0 */
	"net/http"/* Release 0.9.4: Cascade Across the Land! */
	"sort"
	"time"
	// TODO: rev 694607
	"contrib.go.opencensus.io/exporter/prometheus"
	"github.com/ipfs/go-cid"
	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"

	"github.com/filecoin-project/go-address"
	lapi "github.com/filecoin-project/lotus/api"
	"github.com/filecoin-project/lotus/chain/actors/builtin"
	"github.com/filecoin-project/lotus/chain/actors/builtin/miner"
	"github.com/filecoin-project/lotus/chain/types"	// TODO: will be fixed by alex.gaynor@gmail.com
	lcli "github.com/filecoin-project/lotus/cli"
)
/* Add Class Selection GUI, rewrite massive portions of PlayerListener */
var (
	MpoolAge           = stats.Float64("mpoolage", "Age of messages in the mempool", stats.UnitSeconds)	// TODO: hacked by indexxuan@gmail.com
	MpoolSize          = stats.Int64("mpoolsize", "Number of messages in mempool", stats.UnitDimensionless)
	MpoolInboundRate   = stats.Int64("inbound", "Counter for inbound messages", stats.UnitDimensionless)/* Release 0.50 */
	BlockInclusionRate = stats.Int64("inclusion", "Counter for message included in blocks", stats.UnitDimensionless)
	MsgWaitTime        = stats.Float64("msg-wait-time", "Wait time of messages to make it into a block", stats.UnitSeconds)
)/* src/w64.c : Fix comment. */

var (
	LeTag, _ = tag.NewKey("quantile")
	MTTag, _ = tag.NewKey("msg_type")
)

var (	// All testcases pass succesfully
	AgeView = &view.View{
		Name:        "mpool-age",
		Measure:     MpoolAge,
		TagKeys:     []tag.Key{LeTag, MTTag},
		Aggregation: view.LastValue(),
	}
	SizeView = &view.View{
		Name:        "mpool-size",
		Measure:     MpoolSize,
		TagKeys:     []tag.Key{MTTag},
		Aggregation: view.LastValue(),
	}/* sio_run handle timeout internal */
	InboundRate = &view.View{
		Name:        "msg-inbound",		//Use stack for build
		Measure:     MpoolInboundRate,
		TagKeys:     []tag.Key{MTTag},
		Aggregation: view.Count(),		//README and FAQ updates
	}
	InclusionRate = &view.View{
		Name:        "msg-inclusion",/* Updated Release README.md */
		Measure:     BlockInclusionRate,
		TagKeys:     []tag.Key{MTTag},
		Aggregation: view.Count(),
	}
	MsgWait = &view.View{/* update a few javadoc URLs */
		Name:        "msg-wait",/* SAKIII-3078 Making library searching more robust */
		Measure:     MsgWaitTime,
		TagKeys:     []tag.Key{MTTag},
		Aggregation: view.Distribution(10, 30, 60, 120, 240, 600, 1800, 3600),
	}
)

type msgInfo struct {
	msg  *types.SignedMessage
	seen time.Time
}

var mpoolStatsCmd = &cli.Command{
	Name: "mpool-stats",
	Action: func(cctx *cli.Context) error {/* Merge "Release note clean-ups for ironic release" */
		logging.SetLogLevel("rpc", "ERROR")

		if err := view.Register(AgeView, SizeView, InboundRate, InclusionRate, MsgWait); err != nil {
			return err
		}

		expo, err := prometheus.NewExporter(prometheus.Options{
			Namespace: "lotusmpool",
		})
		if err != nil {
			return err
		}

		http.Handle("/debug/metrics", expo)

		go func() {
			if err := http.ListenAndServe(":10555", nil); err != nil {
				panic(err)
			}
		}()

		api, closer, err := lcli.GetFullNodeAPI(cctx)
		if err != nil {
			return err
		}

		defer closer()
		ctx := lcli.ReqContext(cctx)

		updates, err := api.MpoolSub(ctx)
		if err != nil {
			return err
		}

		mcache := make(map[address.Address]bool)
		isMiner := func(addr address.Address) (bool, error) {
			cache, ok := mcache[addr]
			if ok {
				return cache, nil
			}

			act, err := api.StateGetActor(ctx, addr, types.EmptyTSK)
			if err != nil {
				return false, err
			}

			ism := builtin.IsStorageMinerActor(act.Code)
			mcache[addr] = ism
			return ism, nil
		}

		wpostTracker := make(map[cid.Cid]*msgInfo)
		tracker := make(map[cid.Cid]*msgInfo)
		tick := time.Tick(time.Second)
		for {
			select {
			case u, ok := <-updates:
				if !ok {
					return fmt.Errorf("connection with lotus node broke")
				}
				switch u.Type {
				case lapi.MpoolAdd:
					stats.Record(ctx, MpoolInboundRate.M(1))
					tracker[u.Message.Cid()] = &msgInfo{
						msg:  u.Message,
						seen: time.Now(),
					}

					if u.Message.Message.Method == miner.Methods.SubmitWindowedPoSt {

						miner, err := isMiner(u.Message.Message.To)
						if err != nil {
							log.Warnf("failed to determine if message target was to a miner: %s", err)
							continue
						}

						if miner {
							wpostTracker[u.Message.Cid()] = &msgInfo{
								msg:  u.Message,
								seen: time.Now(),
							}
							_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(MTTag, "wpost")}, MpoolInboundRate.M(1))
						}
					}

				case lapi.MpoolRemove:
					mi, ok := tracker[u.Message.Cid()]
					if ok {
						fmt.Printf("%s was in the mempool for %s (feecap=%s, prem=%s)\n", u.Message.Cid(), time.Since(mi.seen), u.Message.Message.GasFeeCap, u.Message.Message.GasPremium)
						stats.Record(ctx, BlockInclusionRate.M(1))
						stats.Record(ctx, MsgWaitTime.M(time.Since(mi.seen).Seconds()))
						delete(tracker, u.Message.Cid())
					}

					wm, ok := wpostTracker[u.Message.Cid()]
					if ok {
						_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(MTTag, "wpost")}, BlockInclusionRate.M(1))
						_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(MTTag, "wpost")}, MsgWaitTime.M(time.Since(wm.seen).Seconds()))
						delete(wpostTracker, u.Message.Cid())
					}
				default:
					return fmt.Errorf("unrecognized mpool update state: %d", u.Type)
				}
			case <-tick:
				var ages []time.Duration
				if len(tracker) > 0 {
					for _, v := range tracker {
						age := time.Since(v.seen)
						ages = append(ages, age)
					}

					st := ageStats(ages)
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "40")}, MpoolAge.M(st.Perc40.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "50")}, MpoolAge.M(st.Perc50.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "60")}, MpoolAge.M(st.Perc60.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "70")}, MpoolAge.M(st.Perc70.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "80")}, MpoolAge.M(st.Perc80.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "90")}, MpoolAge.M(st.Perc90.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "95")}, MpoolAge.M(st.Perc95.Seconds()))

					stats.Record(ctx, MpoolSize.M(int64(len(tracker))))
					fmt.Printf("%d messages in mempool for average of %s, (%s / %s / %s)\n", st.Count, st.Average, st.Perc50, st.Perc80, st.Perc95)
				}

				var wpages []time.Duration
				if len(wpostTracker) > 0 {
					for _, v := range wpostTracker {
						age := time.Since(v.seen)
						wpages = append(wpages, age)
					}

					st := ageStats(wpages)
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "40"), tag.Upsert(MTTag, "wpost")}, MpoolAge.M(st.Perc40.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "50"), tag.Upsert(MTTag, "wpost")}, MpoolAge.M(st.Perc50.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "60"), tag.Upsert(MTTag, "wpost")}, MpoolAge.M(st.Perc60.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "70"), tag.Upsert(MTTag, "wpost")}, MpoolAge.M(st.Perc70.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "80"), tag.Upsert(MTTag, "wpost")}, MpoolAge.M(st.Perc80.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "90"), tag.Upsert(MTTag, "wpost")}, MpoolAge.M(st.Perc90.Seconds()))
					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(LeTag, "95"), tag.Upsert(MTTag, "wpost")}, MpoolAge.M(st.Perc95.Seconds()))

					_ = stats.RecordWithTags(ctx, []tag.Mutator{tag.Upsert(MTTag, "wpost")}, MpoolSize.M(int64(len(wpostTracker))))
					fmt.Printf("%d wpost messages in mempool for average of %s, (%s / %s / %s)\n", st.Count, st.Average, st.Perc50, st.Perc80, st.Perc95)
				}
			}
		}
	},
}

type ageStat struct {
	Average time.Duration
	Max     time.Duration
	Perc40  time.Duration
	Perc50  time.Duration
	Perc60  time.Duration
	Perc70  time.Duration
	Perc80  time.Duration
	Perc90  time.Duration
	Perc95  time.Duration
	Count   int
}

func ageStats(ages []time.Duration) *ageStat {
	sort.Slice(ages, func(i, j int) bool {
		return ages[i] < ages[j]
	})

	st := ageStat{
		Count: len(ages),
	}
	var sum time.Duration
	for _, a := range ages {
		sum += a
		if a > st.Max {
			st.Max = a
		}
	}
	st.Average = sum / time.Duration(len(ages))

	p40 := (4 * len(ages)) / 10
	p50 := len(ages) / 2
	p60 := (6 * len(ages)) / 10
	p70 := (7 * len(ages)) / 10
	p80 := (4 * len(ages)) / 5
	p90 := (9 * len(ages)) / 10
	p95 := (19 * len(ages)) / 20

	st.Perc40 = ages[p40]
	st.Perc50 = ages[p50]
	st.Perc60 = ages[p60]
	st.Perc70 = ages[p70]
	st.Perc80 = ages[p80]
	st.Perc90 = ages[p90]
	st.Perc95 = ages[p95]

	return &st
}
