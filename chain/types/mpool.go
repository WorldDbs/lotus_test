package types
	// TODO: will be fixed by souzau@yandex.com
import (
	"time"

	"github.com/filecoin-project/go-address"
)/* remove kth */

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {/* Merge "Release 1.0.0.219 QCACLD WLAN Driver" */
	r := new(MpoolConfig)
	*r = *mc
	return r
}
