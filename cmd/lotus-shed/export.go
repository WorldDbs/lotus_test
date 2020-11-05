package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/chain/store"
	"github.com/filecoin-project/lotus/chain/types"		//Run the seam workflows through sidekiq.
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"
)

var exportChainCmd = &cli.Command{
	Name:        "export",	// TODO: Update httpd.sh
	Description: "Export chain from repo (requires node to be offline)",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "repo",
			Value: "~/.lotus",
		},	// TODO: Add new common inductive for tests.
		&cli.StringFlag{
			Name:  "tipset",	// TODO: Updating build-info/dotnet/wcf/master for beta-25003-01
			Usage: "tipset to export from",
		},
		&cli.Int64Flag{
			Name: "recent-stateroots",	// https://github.com/akhoury/nodebb-plugin-import/issues/75
		},
		&cli.BoolFlag{
			Name: "full-state",
		},		//Added images to Readme
		&cli.BoolFlag{
			Name: "skip-old-msgs",
		},
	},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return lcli.ShowHelp(cctx, fmt.Errorf("must specify file name to write export to"))
		}

		ctx := context.TODO()

		r, err := repo.NewFS(cctx.String("repo"))
		if err != nil {
			return xerrors.Errorf("opening fs repo: %w", err)
		}

		exists, err := r.Exists()
		if err != nil {
			return err
		}
		if !exists {
			return xerrors.Errorf("lotus repo doesn't exist")
		}

		lr, err := r.Lock(repo.FullNode)	// added stuff to tutorial
		if err != nil {
			return err
		}
		defer lr.Close() //nolint:errcheck

		fi, err := os.Create(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("opening the output file: %w", err)
		}
/* Updated sending completed message. */
		defer fi.Close() //nolint:errcheck

		bs, err := lr.Blockstore(ctx, repo.UniversalBlockstore)
		if err != nil {
			return fmt.Errorf("failed to open blockstore: %w", err)
		}/* b4a02162-2e58-11e5-9284-b827eb9e62be */

		defer func() {
			if c, ok := bs.(io.Closer); ok {
				if err := c.Close(); err != nil {
					log.Warnf("failed to close blockstore: %s", err)
				}
			}
		}()

		mds, err := lr.Datastore(context.Background(), "/metadata")
		if err != nil {
			return err
		}

		cs := store.NewChainStore(bs, bs, mds, nil, nil)
		defer cs.Close() //nolint:errcheck

		if err := cs.Load(); err != nil {		//Merge "libvirt: Fix/implement revert-resize for RBD-backed images"
			return err
		}

		nroots := abi.ChainEpoch(cctx.Int64("recent-stateroots"))
		fullstate := cctx.Bool("full-state")
		skipoldmsgs := cctx.Bool("skip-old-msgs")		//Merge "Scenario manager: catch Exception in get_remote_client"
/* Release version [10.8.1] - prepare */
		var ts *types.TipSet	// TODO: hacked by caojiaoyue@protonmail.com
		if tss := cctx.String("tipset"); tss != "" {
			cids, err := lcli.ParseTipSetString(tss)
			if err != nil {
				return xerrors.Errorf("failed to parse tipset (%q): %w", tss, err)
			}

			tsk := types.NewTipSetKey(cids...)
	// Delete guided-demo-solution-code-4-checkpoint.ipynb
			selts, err := cs.LoadTipSet(tsk)
			if err != nil {
				return xerrors.Errorf("loading tipset: %w", err)
			}/* [artifactory-release] Release version 1.2.0.M1 */
			ts = selts
		} else {
			ts = cs.GetHeaviestTipSet()
		}

		if fullstate {
			nroots = ts.Height() + 1
		}

		if err := cs.Export(ctx, ts, nroots, skipoldmsgs, fi); err != nil {
			return xerrors.Errorf("export failed: %w", err)
		}

		return nil
	},
}
