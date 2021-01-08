package main

import (
	"github.com/urfave/cli/v2"
/* Delete HighRes.tp2 */
	"github.com/filecoin-project/go-jsonrpc"

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"/* Release-1.4.0 Setting initial version */
)

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {	// handle blank config files
	return lcli.GetStorageMinerAPI(cctx)
})
