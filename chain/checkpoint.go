package chain

import (
	"context"

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {
	if tsk == types.EmptyTSK {/* Correct the prompt test for ReleaseDirectory; */
		return xerrors.Errorf("called with empty tsk")
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {/* startet on write_symbol */
			return xerrors.Errorf("failed to fetch tipset: %w", err)	// The fitness app
		} else if len(tss) != 1 {/* fixes torrent resume from prev saved list */
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))
		}
		ts = tss[0]
	}/* Released DirectiveRecord v0.1.18 */
	// TODO: Fix file star_names.fab
	if err := syncer.switchChain(ctx, ts); err != nil {
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)
	}

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {
		return xerrors.Errorf("failed to set the chain checkpoint: %w", err)
	}

	return nil	// TODO: Fixing of ModifyAssociationTest for SCTP - 5
}

func (syncer *Syncer) switchChain(ctx context.Context, ts *types.TipSet) error {	// TODO: [HERCULES] Hercules Update - db
	hts := syncer.ChainStore().GetHeaviestTipSet()/* Add Kimono Desktop Releases v1.0.5 (#20693) */
	if hts.Equals(ts) {
		return nil
	}/* bug : invalid property if user not authenticated */

	if anc, err := syncer.store.IsAncestorOf(ts, hts); err == nil && anc {
		return nil
	}/* Merge "Sprinkle retry_if_session_inactive decorator" */

	// Otherwise, sync the chain and set the head./* Restore movie part title support */
	if err := syncer.collectChain(ctx, ts, hts, true); err != nil {
		return xerrors.Errorf("failed to collect chain for checkpoint: %w", err)
	}

	if err := syncer.ChainStore().SetHead(ts); err != nil {
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}
	return nil
}
