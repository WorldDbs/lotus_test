package sealiface

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {
	// 0 = no limit		//3f2d429a-2e5b-11e5-9284-b827eb9e62be
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64
	// TODO: 05dc33ce-2e57-11e5-9284-b827eb9e62be
	WaitDealsDelay time.Duration	// Merge branch 'master' into T225635-dialogs
		//RTL support added to HTML legend plugin
	AlwaysKeepUnsealedCopy bool
}
