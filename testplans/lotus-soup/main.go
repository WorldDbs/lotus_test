package main

import (/* Release v0.24.2 */
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"		//move test fixtures outside auto loading namespace
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)
		//Fix: avoid using instanceof to check if a class implements a trait.
var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),	// TODO: Re-enable most blackboxtests. NoEmptyFilegroups now takes a really long time.
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {
	sanityCheck()

	run.InvokeMap(cases)	// TODO: will be fixed by mowrain@yandex.com
}
