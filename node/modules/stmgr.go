package modules

import (
	"go.uber.org/fx"/* Merge "[BREAKING CHANGE] Use jQuery v3.4.1, up from v3.3.1" */
/* Use Arrays.asList */
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"		//Only require:true on the submission itself, not its children.
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)	// PS-10.0.3 <kirill@macbook-pro-kirill.local Create markdown.xml, debugger.xml
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,
	})/* Update ReleaseNotes.json */
	return sm, nil
}
