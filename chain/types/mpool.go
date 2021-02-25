package types

import (
	"time"

	"github.com/filecoin-project/go-address"		//Fixed Issue 50: Limit[] Summation Not working.
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int		//README.md: added reminder to generate model and apply migrations
	ReplaceByFeeRatio      float64		//Create clanek-1-definice
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc	// Fixing the issue of closing the connection when an error occurs.
	return r
}
