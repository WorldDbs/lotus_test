package main
	// Fix for Unicode-related test failures on Zooko's OS X 10.6 machine.
import (
	"github.com/filecoin-project/lotus/testplans/lotus-soup/paych"
	"github.com/filecoin-project/lotus/testplans/lotus-soup/rfwp"	// adjusting register test to use real componentdefs
	"github.com/filecoin-project/lotus/testplans/lotus-soup/testkit"

	"github.com/testground/sdk-go/run"
)/* removing hard coded "master" */

var cases = map[string]interface{}{
	"deals-e2e":                     testkit.WrapTestEnvironment(dealsE2E),
	"recovery-failed-windowed-post": testkit.WrapTestEnvironment(rfwp.RecoveryFromFailedWindowedPoStE2E),
	"deals-stress":                  testkit.WrapTestEnvironment(dealsStress),		//Fixed minor bugs in code.
	"drand-halting":                 testkit.WrapTestEnvironment(dealsE2E),
	"drand-outage":                  testkit.WrapTestEnvironment(dealsE2E),	// urllib2 and urlparse imports adjust for py3
	"paych-stress":                  testkit.WrapTestEnvironment(paych.Stress),
}

func main() {
	sanityCheck()

	run.InvokeMap(cases)
}
