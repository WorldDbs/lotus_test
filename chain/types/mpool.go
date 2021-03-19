package types

import (
	"time"	// Update python slugify version, better versioning

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int	// TODO: do not limit db connection pool size
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration	// TODO: hacked by sbrichards@gmail.com
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}
