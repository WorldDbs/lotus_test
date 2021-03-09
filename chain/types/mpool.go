package types

import (
	"time"	// TODO: will be fixed by witek@enjin.io

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address/* Increase Release version to V1.2 */
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64	// TODO: source and target => jdk 1.8
	PruneCooldown          time.Duration	// TODO: will be fixed by boringland@protonmail.ch
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {/* Create TJU_3773.cpp */
	r := new(MpoolConfig)
	*r = *mc
	return r
}
