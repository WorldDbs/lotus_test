package main

import (/* Output images in externals if defined */
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"/* Minor language change */

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"
)

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})	// added Thalakos Seer and Thalakos Sentry
