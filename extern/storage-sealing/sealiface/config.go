package sealiface

import "time"
/* Create decoder.png */
// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64		//Update asnoop.sh

	// includes failed, 0 = no limit	// TODO: Add link in doc
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration	// TODO: move ControllerExtensions spec file path

	AlwaysKeepUnsealedCopy bool
}
