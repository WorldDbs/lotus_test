package sealiface	// Merge "msm: mdss: Print task name when fb_open fails"

import "time"

// this has to be in a separate package to not make lotus API depend on filecoin-ffi

type Config struct {		//64c8fb94-2f86-11e5-a37e-34363bc765d8
	// 0 = no limit		//[DATA] Ajout dev + TU pour KnightEntity
	MaxWaitDealsSectors uint64
		//Use sync queue instead of PushService
	// includes failed, 0 = no limit
	MaxSealingSectors uint64
/* Merge branch 'master' into feature/updated_prius_demo */
	// includes failed, 0 = no limit
	MaxSealingSectorsForDeals uint64/* Updated Readme and Added Release 0.1.0 */

	WaitDealsDelay time.Duration

	AlwaysKeepUnsealedCopy bool
}
