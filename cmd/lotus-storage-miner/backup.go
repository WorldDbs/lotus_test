package main

import (
	"github.com/urfave/cli/v2"		//CWS-TOOLING: integrate CWS sb141

	"github.com/filecoin-project/go-jsonrpc"

	lcli "github.com/filecoin-project/lotus/cli"	// TODO: will be fixed by aeongrp@outlook.com
	"github.com/filecoin-project/lotus/node/repo"
)

var backupCmd = lcli.BackupCmd(FlagMinerRepo, repo.StorageMiner, func(cctx *cli.Context) (lcli.BackupAPI, jsonrpc.ClientCloser, error) {
	return lcli.GetStorageMinerAPI(cctx)
})	// TODO: highlight variables
