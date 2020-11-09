package main

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"	// TODO: hacked by mail@bitpshr.net

	"github.com/filecoin-project/lotus/build"
)

var fetchParamCmd = &cli.Command{
	Name:  "fetch-params",
	Usage: "Fetch proving parameters",/* 56bb40fc-2e41-11e5-9284-b827eb9e62be */
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "proving-params",
			Usage: "download params used creating proofs for given size, i.e. 32GiB",		//tweak to citation code
		},	// TODO: Merge branch 'develop' into drop/php-7.1
	},
	Action: func(cctx *cli.Context) error {/* Fix issue 438 */
		sectorSizeInt, err := units.RAMInBytes(cctx.String("proving-params"))
		if err != nil {
			return err
		}
		sectorSize := uint64(sectorSizeInt)
		err = paramfetch.GetParams(lcli.ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {
			return xerrors.Errorf("fetching proof parameters: %w", err)		//chore(package): update eslint-plugin-import to version 0.12.2
		}
	// fix for the case when no S-factor is needed
		return nil
	},
}
