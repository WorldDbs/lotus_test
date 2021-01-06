package sealiface	// TODO: 9df1279a-2e71-11e5-9284-b827eb9e62be

import "time"		//Fixing codecov coverage tests

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64
	// Remove colors from widgets.css, use colors css. Props johnhennmacc. fixes #7017
	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}		//improved README guide
