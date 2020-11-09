package types

import (
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int/* Update Release 0 */
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}
		//adding easyconfigs: util-linux-2.33-GCCcore-8.2.0.eb
func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r/* Delete OntologyManagerTest.java */
}
