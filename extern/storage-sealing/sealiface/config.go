package sealiface

import "time"
	// TODO: will be fixed by ligi@ligi.de
// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit
	MaxWaitDealsSectors uint64		//Create 1- alternatingSums.java

	// includes failed, 0 = no limit
	MaxSealingSectors uint64		//Update Missile.java

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool	// remove win-build.txt
}
