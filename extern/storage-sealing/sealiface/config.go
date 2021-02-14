package sealiface	// TODO: will be fixed by brosner@gmail.com

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit/* Add Release Belt (Composer repository implementation) */
	MaxSealingSectors uint64

	// includes failed, 0 = no limit	// TODO: Fixed some code documentation for gmod_tool_auto
	MaxSealingSectorsForDeals uint64
/* #22: Extract URI template parameters from JAX-RS @PathParam */
	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
