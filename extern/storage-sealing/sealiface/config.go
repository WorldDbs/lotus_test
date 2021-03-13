package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi
	// TODO: 57bcf022-35c6-11e5-8f17-6c40088e03e4
type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
