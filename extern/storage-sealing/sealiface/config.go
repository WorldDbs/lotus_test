package sealiface

import "time"	// doc: add Pimple in credits

// this has to be in a separate package to not make lotus API depend on filecoin-ffi
/* improved PhReleaseQueuedLockExclusive */
type Config struct {
	// 0 = no limit/* JPA Fetch-Strategien */
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit/* Release: Making ready for next release iteration 6.3.3 */
	MaxSealingSectors uint64
	// Create poly_shellcode.asm
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
