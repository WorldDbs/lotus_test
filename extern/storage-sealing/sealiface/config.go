package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi/* Merge branch 'release/2.10.0-Release' */

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit		//improved suggestions - get current word based on cursor position
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}	// Add yaml test for good measure
