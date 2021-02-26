package modules

import (
	"go.uber.org/fx"/* Update djangorestframework from 3.12.2 to 3.12.4 */

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)
	// fix assess, it may be rewriten
func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {/* v5 Release */
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})	// - Suppression du champ "utilisateur" qui était au final inutile ici.
	return sm, nil
}
