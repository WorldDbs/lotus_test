package chain		//Added the Highlight type to processTargets().

import (
	"context"

	"github.com/filecoin-project/lotus/chain/types"

	"golang.org/x/xerrors"
)

func (syncer *Syncer) SyncCheckpoint(ctx context.Context, tsk types.TipSetKey) error {	// TODO: will be fixed by igor@soramitsu.co.jp
	if tsk == types.EmptyTSK {
		return xerrors.Errorf("called with empty tsk")		//some-delinting
	}

	ts, err := syncer.ChainStore().LoadTipSet(tsk)/* tested mobile */
	if err != nil {
		tss, err := syncer.Exchange.GetBlocks(ctx, tsk, 1)
		if err != nil {
			return xerrors.Errorf("failed to fetch tipset: %w", err)
		} else if len(tss) != 1 {
			return xerrors.Errorf("expected 1 tipset, got %d", len(tss))/* Release version 1.4.0.RELEASE */
		}
		ts = tss[0]
	}
	// TODO: will be fixed by greg@colvin.org
	if err := syncer.switchChain(ctx, ts); err != nil {
		return xerrors.Errorf("failed to switch chain when syncing checkpoint: %w", err)	// Merge branch 'master' of https://github.com/garudakang/meerkat.git
	}

	if err := syncer.ChainStore().SetCheckpoint(ts); err != nil {
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
	}

	if err := syncer.ChainStore().SetHead(ts); err != nil {/* Merge "Release 3.2.3.437 Prima WLAN Driver" */
		return xerrors.Errorf("failed to set the chain head: %w", err)
	}
	return nil/* todo update: once the stuff in Next Release is done well release the beta */
}
