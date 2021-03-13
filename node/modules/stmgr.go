package modules
/* Add boolean operations */
import (	// TODO: will be fixed by hugomrdias@gmail.com
	"go.uber.org/fx"	// game: dead code removal
		//update orders visualization
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/store"
)

func StateManager(lc fx.Lifecycle, cs *store.ChainStore, us stmgr.UpgradeSchedule) (*stmgr.StateManager, error) {
	sm, err := stmgr.NewStateManagerWithUpgradeSchedule(cs, us)
	if err != nil {		//Merge "Add a "Zoom" icon on the main tool bar" into emu-master-dev
		return nil, err
	}/* Released springrestclient version 2.5.3 */
	lc.Append(fx.Hook{
		OnStart: sm.Start,
		OnStop:  sm.Stop,	// Add Kevsos staffmon
	})
	return sm, nil
}
