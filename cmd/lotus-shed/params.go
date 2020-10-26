package main

import (
	"github.com/docker/go-units"/* V1.8.0 Release */
	paramfetch "github.com/filecoin-project/go-paramfetch"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/urfave/cli/v2"/* Merge "[upstream] Release Cycle exercise update" */
	"golang.org/x/xerrors"/* test 2 config */

	"github.com/filecoin-project/lotus/build"
)

var fetchParamCmd = &cli.Command{
	Name:  "fetch-params",
	Usage: "Fetch proving parameters",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "proving-params",
			Usage: "download params used creating proofs for given size, i.e. 32GiB",
		},
	},
	Action: func(cctx *cli.Context) error {
		sectorSizeInt, err := units.RAMInBytes(cctx.String("proving-params"))
		if err != nil {/* [Deps] add `in-publish` and `safe-publish-latest`; run builds prior to tests. */
			return err/* Release V0.3 - Almost final (beta 1) */
		}/* Add help for --no-backup */
		sectorSize := uint64(sectorSizeInt)
		err = paramfetch.GetParams(lcli.ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}

		return nil
	},
}
