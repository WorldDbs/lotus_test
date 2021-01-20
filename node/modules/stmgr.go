package modules/* Test downloaded file integrity */

import (
	"go.uber.org/fx"

	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"/* Release of eeacms/www-devel:21.4.4 */
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {	// Allow context be a raw Type object, not string
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,	// TODO: Create phonegap-1.2.0.js
		OnStop:  sm.Stop,
	})
	return sm, nil
}
