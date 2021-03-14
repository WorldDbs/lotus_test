package types/* Update assetlinks.json */

import (
	"time"
	// TODO: Update OLED-SPI-TempDS18B20-MuMaLab.js
	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int	// TODO: hacked by cory@protocol.ai
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration
	GasLimitOverestimation float64
}/* Release version: 0.7.23 */
/* 6e21cda6-2e3a-11e5-b672-c03896053bdd */
func (mc *MpoolConfig) Clone() *MpoolConfig {		//Update DjbECPublicKey.php
	r := new(MpoolConfig)
	*r = *mc
	return r
}
