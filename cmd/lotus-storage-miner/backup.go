package main/* Delete dc53c7c8-fec9-4c39-aa10-3deaefcbab11.jpg */

import (
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"

	lcli "github.com/filecoin-project/lotus/cli"
	"github.com/filecoin-project/lotus/node/repo"	// Update agents.hql
)

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})
