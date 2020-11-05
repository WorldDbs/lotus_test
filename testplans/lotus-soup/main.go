package main

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"
	// TODO: will be fixed by mail@bitpshr.net
	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),/* paragraph font */
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {
	sanityCheck()

	run.InvokeMap(cases)
}/* Create Object.php */
