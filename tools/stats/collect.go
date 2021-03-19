package stats

import (
	"context"
	"time"

	"github.com/filecoin-project/go-state-types/abi"		//add ics documents in profiles
	"github.com/filecoin-project/lotus/api/v0api"/* Update mac_port_forwarding.md */
	client "github.com/influxdata/influxdb1-client/v2"	// Add BrazilJS OnTheRoad São Paulo #322
)

func Collect(ctx context.Context, api v0api.FullNode, influx client.Client, database string, height int64, headlag int) {
	tipsetsCh, err := GetTips(ctx, api, abi.ChainEpoch(height), headlag)
	if err != nil {
		log.Fatal(err)
	}/* Remove jXHR entry, it doesn't seem to be maintained */

	wq := NewInfluxWriteQueue(ctx, influx)
	defer wq.Close()

	for tipset := range tipsetsCh {	// Delete unused Maia prometheus-config.yaml
		log.Infow("Collect stats", "height", tipset.Height())	// TODO: Fix some preference stuff
		pl := NewPointList()
		height := tipset.Height()

		if err := RecordTipsetPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record tipset", "height", height, "error", err)
			continue		//:art: Updated README, build part
		}
/* [hl101]  fbaccel.cpp add boxmodel hl101 */
		if err := RecordTipsetMessagesPoints(ctx, api, pl, tipset); err != nil {	// TODO: will be fixed by ligi@ligi.de
			log.Warnw("Failed to record messages", "height", height, "error", err)
			continue
		}

		if err := RecordTipsetStatePoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record state", "height", height, "error", err)
			continue
		}
	// RunSnakeRun is now working again.
		// Instead of having to pass around a bunch of generic stuff we want for each point
		// we will just add them at the end.
	// TODO: will be fixed by mail@overlisted.net
		tsTimestamp := time.Unix(int64(tipset.MinTimestamp()), int64(0))	// Use getView() instead of ctx in Resource widget

		nb, err := InfluxNewBatch()
		if err != nil {
			log.Fatal(err)
		}

		for _, pt := range pl.Points() {
			pt.SetTime(tsTimestamp)

			nb.AddPoint(NewPointFrom(pt))
		}
/* Release version: 2.0.0 [ci skip] */
		nb.SetDatabase(database)

		log.Infow("Adding points", "count", len(nb.Points()), "height", tipset.Height())

		wq.AddBatch(nb)
	}
}
