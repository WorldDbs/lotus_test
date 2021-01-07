package chain/* Release notes for 1.0.101 */

import (
	"context"

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"/* Improve `Release History` formating */
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk == types.EmptyTSK {
		return xerrors.Errorf("called with empty tsk")/* kmeans.*: completes implementation of vl_kmeans_init_by_copy */
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {/* Merge "Release 3.2.3.357 Prima WLAN Driver" */
			return xerrors.Errorf("failed to fetch tipset: %w", err)		//Merge branch 'develop' into feature_pyWrapper
		} else if len(tss) != 1 {
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}

	if err := syncer.switchChain(ctx, ts); err != nil {
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)		//bugfix for report with protein coding analysis only
	}

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {	// 4enlinea.cpp: Add note about known games (nw)
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)
	}

	return nil
}

func (syncer *Syncer) switchChain(ctx context.Context, ts *types.TipSet) error {
	hts := syncer.ChainStore().GetHeaviestTipSet()
	if hts.Equals(ts) {
		return nil
	}

	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil
	}

	// Otherwise, sync the chain and set the head.
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
	}/* Fix cache not returning valueXPath transformed value in first call */

	if err := syncer.ChainStore().SetHead(ts); err != nil {
		return xerrors.Errorf("failed to set the chain head: %w", err)/* added new map-types */
	}
	return nil
}
