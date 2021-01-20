package main

import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"/* filtrage commande prestataire */
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"		//add raspbian compatibility hint to README.md

	"github.com/testground/sdk-go/run"
)

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),	// TODO: [Style] : Fix style and space.
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {	// TODO: Ensure buffer size is set to at least the minimum (1024).
	sanityCheck()		//Added function to add/remove for class structure through console

	run.InvokeMap(cases)	// TODO: Created IMG_7901.JPG
}
