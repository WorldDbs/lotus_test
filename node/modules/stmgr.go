package modules

import (
	"go.uber.org/fx"	// TODO: Merge "Fix CodeSniffer errors and warnings on yet more API classes"
/* DDBNEXT-748: Selectionbox background color  */
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)/* [artifactory-release] Release version 0.9.10.RELEASE */

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {/* Advance clock and lines. */
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})
	return sm, nil
}
