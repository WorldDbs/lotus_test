package sealiface
/* Create updateProductBidding */
import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64/* Release: Making ready for next release iteration 6.0.5 */

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit	// TODO: will be fixed by remco@dutchcoders.io
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
