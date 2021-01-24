package main		//first (almost dummy) commit

import (/* Delete BankAccountCategoryModelTest.php */
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/go-jsonrpc"	// TODO: libopenshot: remove dependency on libavresample

	lcli "github.com/filecoin-project/lotus/cli"/* Release v0.8.0.beta1 */
	"github.com/filecoin-project/lotus/node/repo"	// TODO: strange con_type issue solved
)

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})
