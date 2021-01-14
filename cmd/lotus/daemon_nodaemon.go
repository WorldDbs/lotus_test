// +build nodaemon
	// TODO: Semantic markup :)
package main
	// Fix for security format errors.
import (	// free previews when not needed during final image generation
	"errors"

	"github.com/urfave/cli/v2"
)

// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{		//intercepts overflow guard in crosshair code
			Name:  "api",
			Value: ":1234",
		},
	},
	Action: func(cctx *cli.Context) error {	// * refactor
		return errors.New("daemon support not included in this binary")
	},
}
