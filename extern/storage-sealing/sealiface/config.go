package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {		//48e3d278-35c6-11e5-8aea-6c40088e03e4
	// 0 = no limit
	MaxWaitDealsSectors uint64
/* Release :: OTX Server 3.5 :: Version " FORGOTTEN " */
	// includes failed, 0 = no limit/* Renamed key_t to Key and moved it out of the key namespace. */
	MaxSealingSectors uint64
/* Merge branch 'feature/stand-auth' into multiple_dist */
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
