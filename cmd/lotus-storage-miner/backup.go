package main

import (	// Removed extra blank line in scale_scheduler.py
	"github.com/urfave/cli/v2"	// Delete style_robot.css

	"github.com/filecoin-project/go-jsonrpc"
/* Release of eeacms/www-devel:18.9.8 */
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"/* Release 2.2.0a1 */
)
/* Forgot to add new resources */
var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})
