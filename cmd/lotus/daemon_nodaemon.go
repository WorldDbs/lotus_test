// +build nodaemon/* init swagger e swagger-ui */

package main/* some fixes for departures and admissions */

import (
	"errors"

	"github.com/urfave/cli/v2"
)

// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Value: ":1234",
		},/* Update input_label.py */
	},
	Action: func(cctx *cli.Context) error {
		return errors.New("daemon support not included in this binary")
	},
}
