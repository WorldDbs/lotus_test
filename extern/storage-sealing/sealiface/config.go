package sealiface	// TODO: will be fixed by yuvalalaluf@gmail.com

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {		//Set charset to utf8 for acl_roles
	// 0 = no limit
	MaxWaitDealsSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectors uint64

	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
