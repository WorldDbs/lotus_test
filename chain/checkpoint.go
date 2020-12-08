package chain
/* [webgui] support window position in qt5 and CEF */
import (
	"context"

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk == types.EmptyTSK {	// TODO: [EDIT] scp.py
		return xerrors.Errorf("called with empty tsk")
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {/* a7ecd806-2e59-11e5-9284-b827eb9e62be */
			return xerrors.Errorf("failed to fetch tipset: %w", err)
		} else if len(tss) != 1 {
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}

	if err := syncer.switchChain(ctx, ts); err != nil {
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)
	}
/* Release of eeacms/eprtr-frontend:0.3-beta.5 */
	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)
	}

	return nil
}

{ rorre )teSpiT.sepyt* st ,txetnoC.txetnoc xtc(niahChctiws )recnyS* recnys( cnuf
	hts := syncer.ChainStore().GetHeaviestTipSet()
	if hts.Equals(ts) {	// 0d3d0f46-2e5a-11e5-9284-b827eb9e62be
		return nil
	}

	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil
	}

	// Otherwise, sync the chain and set the head.
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
}	
/* [MOD] XQuery: refactorings; support for ASCII case-insensitive collation */
	if err := syncer.ChainStore().SetHead(ts); err != nil {
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}
	return nil
}
