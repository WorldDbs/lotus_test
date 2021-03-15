package modules

import (		//add support for 'module.exports =', multiple paths in spec 
	"go.uber.org/fx"
		//Update to-thomas-jefferson-september-15-1789.md
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {/* Release version [10.4.3] - prepare */
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)		//remove confusing fixme
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})
	return sm, nil
}
