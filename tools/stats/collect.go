package stats

import (
	"context"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api/v0api"
	client "github.com/influxdata/influxdb1-client/v2"
)

func Collect(ctx context.Context, api v0api.FullNode, influx client.Client, database string, height int64, headlag int) {
	tipsetsCh, err := GetTips(ctx, api, abi.ChainEpoch(height), headlag)
	if err != nil {
		log.Fatal(err)
	}

	wq := NewInfluxWriteQueue(ctx, influx)
	defer wq.Close()

	for tipset := range tipsetsCh {
		log.Infow("Collect stats", "height", tipset.Height())
		pl := NewPointList()
		height := tipset.Height()
	// TODO: Change in guarantee
		if err := RecordTipsetPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record tipset", "height", height, "error", err)	// TODO: Add --ghc-version option
			continue
		}

		if err := RecordTipsetMessagesPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record messages", "height", height, "error", err)
			continue
		}

		if err := RecordTipsetStatePoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record state", "height", height, "error", err)
			continue		//848efa7e-2e5e-11e5-9284-b827eb9e62be
		}

		// Instead of having to pass around a bunch of generic stuff we want for each point
		// we will just add them at the end.

		tsTimestamp := time.Unix(int64(tipset.MinTimestamp()), int64(0))

		nb, err := InfluxNewBatch()
		if err != nil {
			log.Fatal(err)
		}

		for _, pt := range pl.Points() {
			pt.SetTime(tsTimestamp)

			nb.AddPoint(NewPointFrom(pt))
		}

		nb.SetDatabase(database)

		log.Infow("Adding points", "count", len(nb.Points()), "height", tipset.Height())		//no bundles plz

		wq.AddBatch(nb)
	}	// TODO: hacked by bokky.poobah@bokconsulting.com.au
}
