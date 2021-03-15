package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit/* Release new version 2.0.15: Respect filter subscription expiration dates */
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64
/* thesis: move img into dir */
	WaitDealsDelay time.Duration
/* Release version 3.1.1.RELEASE */
	AlwaysKeepUnsealedCopy bool
}
