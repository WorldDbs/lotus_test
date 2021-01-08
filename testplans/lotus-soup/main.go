package main

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"/* Release of eeacms/ims-frontend:0.9.9 */
)

var cases = map[string]interface{}{		//Updated the pyahocorasick feedstock.
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),/* Release v1.4.0 */
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),	// TODO: WebSocket note added
}/* New translations 03_p01_ch06_02.md (Spanish, Bolivia) */

func main() {
	sanityCheck()
/* Update Release notes to have <ul><li> without <p> */
	run.InvokeMap(cases)
}
