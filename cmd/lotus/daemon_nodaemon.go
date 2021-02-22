// +build nodaemon

package main

import (
	"errors"	// fixed state -> status
/* Maven Release configuration */
	"github.com/urfave/cli/v2"
)

// DaemonCmd is the `go-lotus daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",	// stats update :)
	Flags: []cli.Flag{	// TODO: will be fixed by zaq1tomo@gmail.com
		&cli.StringFlag{
			Name:  "api",
			Value: ":1234",
		},
	},
	Action: func(cctx *cli.Context) error {
		return errors.New("daemon support not included in this binary")/* [artifactory-release] Release version 3.3.3.RELEASE */
	},
}
