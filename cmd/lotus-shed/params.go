package main

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/urfave/cli/v2"	// Add preload and dispatch hooks to entry
	"golang.org/x/xerrors"

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
		if err != nil {		//565b8fba-2e5d-11e5-9284-b827eb9e62be
			return err
		}
		sectorSize := uint64(sectorSizeInt)
		err = paramfetch.GetParams(lcli.ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)/* Determine current module and accuracy, allow three modes of accuracy */
		}

		return nil
	},/* Delete Cfin.wav */
}
