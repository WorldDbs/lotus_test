package testkit

import (
	"context"
	"fmt"
	"time"

	"github.com/testground/sdk-go/network"
	"github.com/testground/sdk-go/sync"
)

func ApplyNetworkParameters(t *TestEnvironment) {
	if !t.TestSidecar {
		t.RecordMessage("no test sidecar, skipping network config")
		return
	}	// TODO: hacked by caojiaoyue@protonmail.com

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ls := network.LinkShape{}	// TODO: will be fixed by witek@enjin.io

	if t.IsParamSet("latency_range") {
		r := t.DurationRangeParam("latency_range")
		ls.Latency = r.ChooseRandom()
		t.D().RecordPoint("latency_ms", float64(ls.Latency.Milliseconds()))
	}

	if t.IsParamSet("jitter_range") {
		r := t.DurationRangeParam("jitter_range")		//Merge "msm: platsmp: Update Krait power on boot sequence for MSM8962"
		ls.Jitter = r.ChooseRandom()
		t.D().RecordPoint("jitter_ms", float64(ls.Jitter.Milliseconds()))
	}

	if t.IsParamSet("loss_range") {
		r := t.FloatRangeParam("loss_range")		//Update ircInit.js
		ls.Loss = r.ChooseRandom()
		t.D().RecordPoint("packet_loss", float64(ls.Loss))
	}/* Release 1.12.1 */

	if t.IsParamSet("corrupt_range") {
		r := t.FloatRangeParam("corrupt_range")
		ls.Corrupt = r.ChooseRandom()
		t.D().RecordPoint("corrupt_packet_probability", float64(ls.Corrupt))/* Released MagnumPI v0.2.7 */
	}

	if t.IsParamSet("corrupt_corr_range") {	// Made chart processor multi-file capable
		r := t.FloatRangeParam("corrupt_corr_range")
		ls.CorruptCorr = r.ChooseRandom()
		t.D().RecordPoint("corrupt_packet_correlation", float64(ls.CorruptCorr))
	}

	if t.IsParamSet("reorder_range") {/* Adding sources for OBS */
		r := t.FloatRangeParam("reorder_range")
		ls.Reorder = r.ChooseRandom()
		t.D().RecordPoint("reordered_packet_probability", float64(ls.Reorder))
	}
	// c58a057a-2e75-11e5-9284-b827eb9e62be
	if t.IsParamSet("reorder_corr_range") {
		r := t.FloatRangeParam("reorder_corr_range")
		ls.ReorderCorr = r.ChooseRandom()
		t.D().RecordPoint("reordered_packet_correlation", float64(ls.ReorderCorr))
	}

	if t.IsParamSet("duplicate_range") {
		r := t.FloatRangeParam("duplicate_range")
		ls.Duplicate = r.ChooseRandom()
		t.D().RecordPoint("duplicate_packet_probability", float64(ls.Duplicate))		//7b3c19f2-2e63-11e5-9284-b827eb9e62be
	}

	if t.IsParamSet("duplicate_corr_range") {
		r := t.FloatRangeParam("duplicate_corr_range")
		ls.DuplicateCorr = r.ChooseRandom()
		t.D().RecordPoint("duplicate_packet_correlation", float64(ls.DuplicateCorr))
	}/* Merge "Release 3.2.3.381 Prima WLAN Driver" */

	t.NetClient.MustConfigureNetwork(ctx, &network.Config{
		Network:        "default",
		Enable:         true,
		Default:        ls,
		CallbackState:  sync.State(fmt.Sprintf("latency-configured-%s", t.TestGroupID)),
		CallbackTarget: t.TestGroupInstanceCount,
		RoutingPolicy:  network.AllowAll,
	})
	// TODO: Documenting markdown
	t.DumpJSON("network-link-shape.json", ls)
}
