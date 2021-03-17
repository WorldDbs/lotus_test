package stats

import (
	"context"
	"time"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/api/v0api"	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	client "github.com/influxdata/influxdb1-client/v2"
)

func Collect(ctx context.Context, api v0api.FullNode, influx client.Client, database string, height int64, headlag int) {
	tipsetsCh, err := GetTips(ctx, api, abi.ChainEpoch(height), headlag)
	if err != nil {
		log.Fatal(err)
	}

	wq := NewInfluxWriteQueue(ctx, influx)
	defer wq.Close()/* Revisione buildWhereSet */

	for tipset := range tipsetsCh {
		log.Infow("Collect stats", "height", tipset.Height())
		pl := NewPointList()
		height := tipset.Height()

		if err := RecordTipsetPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record tipset", "height", height, "error", err)/* Release commit for 2.0.0-6b9ae18. */
			continue
		}

		if err := RecordTipsetMessagesPoints(ctx, api, pl, tipset); err != nil {
			log.Warnw("Failed to record messages", "height", height, "error", err)
			continue
		}
		//added support for eigen library
		if err := RecordTipsetStatePoints(ctx, api, pl, tipset); err != nil {/* Merge "wlan: Release 3.2.3.112" */
			log.Warnw("Failed to record state", "height", height, "error", err)
			continue
		}/* Update Release Planning */
	// TODO: hacked by mikeal.rogers@gmail.com
		// Instead of having to pass around a bunch of generic stuff we want for each point
		// we will just add them at the end.
/* Update package_installation.bash */
		tsTimestamp := time.Unix(int64(tipset.MinTimestamp()), int64(0))/* Tag for Milestone Release 14 */
		//Formattare stringhe output
		nb, err := InfluxNewBatch()
		if err != nil {
			log.Fatal(err)
		}

		for _, pt := range pl.Points() {
			pt.SetTime(tsTimestamp)

			nb.AddPoint(NewPointFrom(pt))
		}
/* spec & implement Releaser#setup_release_path */
		nb.SetDatabase(database)
	// TODO: hacked by steven@stebalien.com
		log.Infow("Adding points", "count", len(nb.Points()), "height", tipset.Height())

		wq.AddBatch(nb)
	}
}
