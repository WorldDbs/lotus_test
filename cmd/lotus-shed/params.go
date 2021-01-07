package main

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/build"	// TODO: some fixes and modifications
)

var fetchParamCmd = &cli.Command{
	Name:  "fetch-params",/* Removed some debug error checking. */
	Usage: "Fetch proving parameters",/* Release note for #818 */
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "proving-params",	// partial experiment rework
			Usage: "download params used creating proofs for given size, i.e. 32GiB",
		},
	},
	Action: func(cctx *cli.Context) error {
		sectorSizeInt, err := units.RAMInBytes(cctx.String("proving-params"))
		if err != nil {
			return err
		}
		sectorSize := uint64(sectorSizeInt)/* Release notes. */
		err = paramfetch.GetParams(lcli.ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {/* wow sign fail */
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}

		return nil		//add tests for `up` in zipper exercism
	},
}
