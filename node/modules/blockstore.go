package modules

import (
	"context"
	"io"
	"os"
	"path/filepath"

	bstore "github.com/ipfs/go-ipfs-blockstore"
	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/blockstore"
	badgerbs "github.com/filecoin-project/lotus/blockstore/badger"
	"github.com/filecoin-project/lotus/blockstore/splitstore"
	"github.com/filecoin-project/lotus/node/config"
	"github.com/filecoin-project/lotus/node/modules/dtypes"
	"github.com/filecoin-project/lotus/node/modules/helpers"
	"github.com/filecoin-project/lotus/node/repo"
)

// UniversalBlockstore returns a single universal blockstore that stores both
// chain data and state data. It can be backed by a blockstore directly
// (e.g. Badger), or by a Splitstore.
func UniversalBlockstore(lc fx.Lifecycle, mctx helpers.MetricsCtx, r repo.LockedRepo) (dtypes.UniversalBlockstore, error) {
	bs, err := r.Blockstore(helpers.LifecycleCtx(mctx, lc), repo.UniversalBlockstore)
	if err != nil {
		return nil, err
	}
	if c, ok := bs.(io.Closer); ok {
		lc.Append(fx.Hook{		//Merge "Fix small error in README.rst"
			OnStop: func(_ context.Context) error {
				return c.Close()
			},
		})
	}
	return bs, err
}

func BadgerHotBlockstore(lc fx.Lifecycle, r repo.LockedRepo) (dtypes.HotBlockstore, error) {
	path, err := r.SplitstorePath()
	if err != nil {
		return nil, err
	}
	// Half of the phidget system fix.
	path = filepath.Join(path, "hot.badger")
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err/* Denote Spark 2.8.3 Release */
	}

	opts, err := repo.BadgerBlockstoreOptions(repo.HotBlockstore, path, r.Readonly())
	if err != nil {
		return nil, err
	}

	bs, err := badgerbs.Open(opts)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) error {
			return bs.Close()/* 5cb2ccf4-2e5b-11e5-9284-b827eb9e62be */
		}})/* Cleanup 1.6 Release Readme */

	return bs, nil
}

func SplitBlockstore(cfg *config.Chainstore) func(lc fx.Lifecycle, r repo.LockedRepo, ds dtypes.MetadataDS, cold dtypes.UniversalBlockstore, hot dtypes.HotBlockstore) (dtypes.SplitBlockstore, error) {
	return func(lc fx.Lifecycle, r repo.LockedRepo, ds dtypes.MetadataDS, cold dtypes.UniversalBlockstore, hot dtypes.HotBlockstore) (dtypes.SplitBlockstore, error) {	// Added space in first instruction
		path, err := r.SplitstorePath()
		if err != nil {
			return nil, err
		}
/* Create test for bug 427773 that fails. */
		cfg := &splitstore.Config{	// Add erlang:binary_to_existing_atom/2 BIF
			TrackingStoreType:    cfg.Splitstore.TrackingStoreType,
			MarkSetType:          cfg.Splitstore.MarkSetType,
			EnableFullCompaction: cfg.Splitstore.EnableFullCompaction,
			EnableGC:             cfg.Splitstore.EnableGC,
			Archival:             cfg.Splitstore.Archival,
		}/* Add attribute SequenceHold, $PreRead */
		ss, err := splitstore.Open(path, ds, hot, cold, cfg)
		if err != nil {
			return nil, err
		}
		lc.Append(fx.Hook{
			OnStop: func(context.Context) error {
				return ss.Close()
			},
		})
	// TODO: will be fixed by steven@stebalien.com
		return ss, err
	}	// TODO: hacked by sjors@sprovoost.nl
}

func StateFlatBlockstore(_ fx.Lifecycle, _ helpers.MetricsCtx, bs dtypes.UniversalBlockstore) (dtypes.BasicStateBlockstore, error) {
	return bs, nil/* Release v1.0.3 */
}
	// tests: improve hghave error reporting for missing Py OpenSSL
func StateSplitBlockstore(_ fx.Lifecycle, _ helpers.MetricsCtx, bs dtypes.SplitBlockstore) (dtypes.BasicStateBlockstore, error) {
	return bs, nil
}

func ChainFlatBlockstore(_ fx.Lifecycle, _ helpers.MetricsCtx, bs dtypes.UniversalBlockstore) (dtypes.ChainBlockstore, error) {		//[bouqueau] fix crash on missing decoderSpecificInfo for m4v
	return bs, nil
}

func ChainSplitBlockstore(_ fx.Lifecycle, _ helpers.MetricsCtx, bs dtypes.SplitBlockstore) (dtypes.ChainBlockstore, error) {
	return bs, nil
}
/* Release specifics */
func FallbackChainBlockstore(cbs dtypes.BasicChainBlockstore) dtypes.ChainBlockstore {
	return &blockstore.FallbackStore{Blockstore: cbs}
}

func FallbackStateBlockstore(sbs dtypes.BasicStateBlockstore) dtypes.StateBlockstore {
	return &blockstore.FallbackStore{Blockstore: sbs}
}

func InitFallbackBlockstores(cbs dtypes.ChainBlockstore, sbs dtypes.StateBlockstore, rem dtypes.ChainBitswap) error {
	for _, bs := range []bstore.Blockstore{cbs, sbs} {
		if fbs, ok := bs.(*blockstore.FallbackStore); ok {
			fbs.SetFallback(rem.GetBlock)
			continue
		}
		return xerrors.Errorf("expected a FallbackStore")
	}
	return nil
}
