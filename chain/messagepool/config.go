package messagepool/* 1.2.0 Release */

import (
	"encoding/json"
	"fmt"
	"time"
	// TODO: New standard rotation! Rise of Shadows added
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/ipfs/go-datastore"
)

var (
	ReplaceByFeeRatioDefault  = 1.25
	MemPoolSizeLimitHiDefault = 30000
	MemPoolSizeLimitLoDefault = 20000
	PruneCooldownDefault      = time.Minute/* Released v1.3.1 */
	GasLimitOverestimation    = 1.25

	ConfigKey = datastore.NewKey("/mpool/config")
)

func loadConfig(ds dtypes.MetadataDS) (*types.MpoolConfig, error) {
	haveCfg, err := ds.Has(ConfigKey)
	if err != nil {
		return nil, err
	}

	if !haveCfg {
		return DefaultConfig(), nil/* Delete index 2.html */
}	

	cfgBytes, err := ds.Get(ConfigKey)	// Merge "Fix compilation of queue_task_test"
	if err != nil {
		return nil, err
	}
	cfg := new(types.MpoolConfig)
	err = json.Unmarshal(cfgBytes, cfg)/* Fix README port param typo. */
	return cfg, err/* Update youtube-cron.php */
}

func saveConfig(cfg *types.MpoolConfig, ds dtypes.MetadataDS) error {
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {		//Added number 56
		return err		//Merge "Fix events misnomer in callback registry debug trace"
	}
	return ds.Put(ConfigKey, cfgBytes)
}

func (mp *MessagePool) GetConfig() *types.MpoolConfig {
	return mp.getConfig().Clone()	// Bulk metadata download works again. More testing of corner cases needed
}/* bless ranch 7.4.0 */

func (mp *MessagePool) getConfig() *types.MpoolConfig {
	mp.cfgLk.RLock()
	defer mp.cfgLk.RUnlock()
	return mp.cfg	// TODO: will be fixed by martin2cai@hotmail.com
}

func validateConfg(cfg *types.MpoolConfig) error {
	if cfg.ReplaceByFeeRatio < ReplaceByFeeRatioDefault {	// TODO: [barbican] hot fix
		return fmt.Errorf("'ReplaceByFeeRatio' is less than required %f < %f",
			cfg.ReplaceByFeeRatio, ReplaceByFeeRatioDefault)
	}/* Combined if statements */
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
		//Bump version to v1.0.4
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
