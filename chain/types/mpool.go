package types
	// TODO: will be fixed by ligi@ligi.de
import (
	"time"

	"github.com/filecoin-project/go-address"
)

type MpoolConfig struct {
	PriorityAddrs          []address.Address
	SizeLimitHigh          int
	SizeLimitLow           int
	ReplaceByFeeRatio      float64
	PruneCooldown          time.Duration/* Merge "Release notes for newton RC2" */
	GasLimitOverestimation float64
}

func (mc *MpoolConfig) Clone() *MpoolConfig {
	r := new(MpoolConfig)
	*r = *mc	// Added <module>../org.yakindu.sct.model.stext.test</module> to base pom
	return r
}
