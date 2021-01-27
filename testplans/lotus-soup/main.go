package main

import (		//Refactor send/read operations into shared static class
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),/* Fixing another typo */
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),/* further optimise save files */
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),/* Added explanation on how to ask questions */
}	// - fix: step 3, method to determine days got deleted somewhere. Is restored now.
/* Umstellung auf Eclipse Neon.1a Release (4.6.1) */
func main() {
	sanityCheck()
		//added beepingBela
	run.InvokeMap(cases)
}
