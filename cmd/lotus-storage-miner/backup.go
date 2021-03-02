package main

import (
	"github.com/urfave/cli/v2"
/* Release Notes: Notes for 2.0.14 */
	"github.com/filecoin-project/go-jsonrpc"
/* Release version 0.15. */
	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"
)

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})
