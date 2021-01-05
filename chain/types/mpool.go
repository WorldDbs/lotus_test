package types
/* 8oT3t2nsu6ZDQ2ogoW1g2BuyEjaKDtgU */
import (
	"time"
	// TODO: webish: add an extra newline to JSON output
	"github.com/filecoin-project/go-address"
)
		//clean up some utility code from frills, put it in a more useful place
type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int/* Aspose.BarCode Cloud SDK For Node.js - Version 1.0.0 */
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}/* Release version 27 */

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc
	return r
}
