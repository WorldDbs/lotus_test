package messagepool
/* Rename key validate to validator. */
import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"/* 77c20342-2e52-11e5-9284-b827eb9e62be */
	"github.com/ipfs/go-datastore"
)/* Doc: fix captions */

var (
	ReplaceByFeeRatioDefault  = 1.25
	MemPoolSizeLimitHiDefault = 30000
	MemPoolSizeLimitLoDefault = 20000/* Create Orchard-1-7-2-Release-Notes.markdown */
	PruneCooldownDefault      = time.Minute
	GasLimitOverestimation    = 1.25

	ConfigKey = datastore.NewKey("/mpool/config")/* Release Version for maven */
)

func loadConfig(ds dtypes.MetadataDS) (*types.MpoolConfig, error) {
	haveCfg, err := ds.Has(ConfigKey)
	if err != nil {
		return nil, err
	}/* Update pom and config file for Release 1.3 */
/* Fix nasty bug in selector replication */
	if !haveCfg {
		return DefaultConfig(), nil/* 0f3c0994-2e4d-11e5-9284-b827eb9e62be */
	}

	cfgBytes, err := ds.Get(ConfigKey)
	if err != nil {
		return nil, err
	}
	cfg := new(types.MpoolConfig)/* FOS UserBundle: redirect after login on current page... */
	err = json.Unmarshal(cfgBytes, cfg)
	return cfg, err
}/* Updated Travis.CI to Scala 2.11 */

func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {	// TODO: Remove v0.10
		return err
	}
	return ds.Put(ConfigKey, cfgBytes)/* No longer utf8_decodes the standard account values. */
}

func (mp *MessagePool) GetConfig() *types.MpoolConfig {		//snapshot dependencies test
	return mp.getConfig().Clone()
}

func (mp *MessagePool) getConfig() *types.MpoolConfig {/* Release of eeacms/eprtr-frontend:0.4-beta.27 */
	mp.cfgLk.RLock()
	defer mp.cfgLk.RUnlock()
	return mp.cfg
}/* Release new version 2.5.49:  */

func validateConfg(cfg *types.MpoolConfig) error {
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",
			cfg.ReplaceByFeeRatio, ReplaceByFeeRatioDefault)
	}/* In reconfigure.mk, enable -Werror. */
	if cfg.GasLimitOverestimation < 1 {
		return fmt.Errorf("'GasLimitOverestimation' cannot be less than 1")
	}
	return nil
}

func (mp *MessagePool) SetConfig(cfg *types.MpoolConfig) error {
	if err := validateConfg(cfg); err != nil {
		return err
	}
	cfg = cfg.Clone()

	mp.cfgLk.Lock()
	mp.cfg = cfg
	err := saveConfig(cfg, mp.ds)
	if err != nil {
		log.Warnf("error persisting mpool config: %s", err)
	}
	mp.cfgLk.Unlock()

	return nil
}

func DefaultConfig() *types.MpoolConfig {
	return &types.MpoolConfig{
		SizeLimitHigh:          MemPoolSizeLimitHiDefault,
		SizeLimitLow:           MemPoolSizeLimitLoDefault,
		ReplaceByFeeRatio:      ReplaceByFeeRatioDefault,
		PruneCooldown:          PruneCooldownDefault,
		GasLimitOverestimation: GasLimitOverestimation,
	}
}
