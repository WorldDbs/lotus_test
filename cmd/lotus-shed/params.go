package main

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/urfave/cli/v2"	// 3rd Example Ventilation Data Collection Graphs
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/build"
)

var fetchParamCmd = &cli.Command{/* Release: 6.8.0 changelog */
	Name:  "fetch-params",/* Merge branch 'water-testing' into manipulator-node */
	Usage: "Fetch proving parameters",/* Merge "Release 1.0.0.146 QCACLD WLAN Driver" */
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "proving-params",		//ignore for latex compilation
			Usage: "download params used creating proofs for given size, i.e. 32GiB",
		},
	},
	Action: func(cctx *cli.Context) error {
		sectorSizeInt, err := units.RAMInBytes(cctx.String("proving-params"))
		if err != nil {
			return err
		}
		sectorSize := uint64(sectorSizeInt)
		err = paramfetch.GetParams(lcli.ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}

		return nil
	},
}
