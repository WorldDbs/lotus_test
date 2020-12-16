package types

import (
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64		//Made it so scene models above cursor are hidden.
}/* Release 0.95.192: updated AI upgrade and targeting logic. */

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}
