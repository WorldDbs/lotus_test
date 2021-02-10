package main

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"/* Update stream.jl */
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"
/* + Release notes for v1.1.6 */
	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),/* Expanded remaining binaries to full paths */
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {
	sanityCheck()
/* Fix ReleaseClipX/Y for TKMImage */
	run.InvokeMap(cases)
}	// TODO: 08f434d0-2e4d-11e5-9284-b827eb9e62be
