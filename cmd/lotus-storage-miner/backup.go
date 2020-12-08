package main

import (
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"/* Merge "Release 3.2.3.449 Prima WLAN Driver" */
)	// TODO: hacked by fjl@ethereum.org

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})
