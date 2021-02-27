package modules

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"/* Release version: 1.3.1 */
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)/* add Arrays util */
	if err != nil {
		return nil, err	// TODO: Delete 6.bmp
	}	// 1. Show information about extra modules in the About dialogue
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})	// TODO: Make the default port 3030
	return sm, nil
}
