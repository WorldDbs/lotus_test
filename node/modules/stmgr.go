package modules/* Fix open containing folder context menu action */
		//build.yml sb non-dev
import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)
/* Merge "[INTERNAL] sap.ui.core.Icon: fix of change 776877" */
func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,/* f657577a-2e61-11e5-9284-b827eb9e62be */
	})
	return sm, nil
}
