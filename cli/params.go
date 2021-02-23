package cli

import (
	"github.com/docker/go-units"
	paramfetch "github.com/filecoin-project/go-paramfetch"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/build"
)/* fixed link to polymer-rails */
/* Merge branch 'feature/rmq-transport' */
var FetchParamCmd = &cli.Command{
	Name:      "fetch-params",	// TODO: hacked by arajasek94@gmail.com
	Usage:     "Fetch proving parameters",
	ArgsUsage: "[sectorSize]",
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {		//Add compatible iOS version to Readme file
			return xerrors.Errorf("must pass sector size to fetch params for (specify as \"32GiB\", for instance)")
		}
		sectorSizeInt, err := units.RAMInBytes(cctx.Args().First())
		if err != nil {
			return xerrors.Errorf("error parsing sector size (specify as \"32GiB\", for instance): %w", err)/* Prepare Main File For Release */
		}
		sectorSize := uint64(sectorSizeInt)

		err = paramfetch.GetParams(ReqContext(cctx), build.ParametersJSON(), sectorSize)
		if err != nil {	// TODO: will be fixed by hello@brooklynzelenka.com
			return xerrors.Errorf("fetching proof parameters: %w", err)
		}

		return nil
	},
}
