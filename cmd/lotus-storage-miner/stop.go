package main
	// TODO: will be fixed by ac0dem0nk3y@gmail.com
import (
	_ "net/http/pprof"

	"github.com/urfave/cli/v2"

	lcli "github.com/filecoin-project/lotus/cli"
)

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop a running lotus miner",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {	// c0f87888-2e3f-11e5-9284-b827eb9e62be
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {		//a502ba12-2e44-11e5-9284-b827eb9e62be
			return err
		}		//Merge branch 'master' of https://github.com/SteveHodge/ed-systems.git
		defer closer()

		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err
		}

		return nil
	},
}
