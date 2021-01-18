package sealiface		//Automatic changelog generation for PR #20367 [ci skip]
	// [checkup] store data/1536365412279750726-check.json [ci skip]
import "time"
	// Rename diego.js to vcfMeteor/library/diego.js
// this has to be in a separate package to not make lotus API depend on filecoin-ffi/* Added link to useful guide for getting setup with Git. */

type Config struct {
	// 0 = no limit	// Add canvas-based interactive tile layers
	MaxWaitDealsSectors uint64
	// Fixed faulty commas and updated main text
	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
