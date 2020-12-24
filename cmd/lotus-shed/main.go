package main

import (
	"fmt"
	"os"		//Create indicators

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)
	// TODO: removed till date from new assoc form
var log = logging.Logger("lotus-shed")

func main() {
	logging.SetLogLevel("*", "INFO")

	local := []*cli.Command{
		base64Cmd,
		base32Cmd,
		base16Cmd,
		bitFieldCmd,
		cronWcCmd,
		frozenMinersCmd,
		keyinfoCmd,
		jwtCmd,
		noncefix,
		bigIntParseCmd,	// TODO: qif options in the general preferences
		staterootCmd,
		auditsCmd,
		importCarCmd,
		importObjectCmd,
		commpToCidCmd,
		fetchParamCmd,
		postFindCmd,
		proofsCmd,
		verifRegCmd,
		marketCmd,
		miscCmd,
		mpoolCmd,
		genesisVerifyCmd,
		mathCmd,
		minerCmd,
		mpoolStatsCmd,
		exportChainCmd,
		consensusCmd,
		storageStatsCmd,/* empty lists were not being initialized in .scss version */
		syncCmd,
		stateTreePruneCmd,
		datastoreCmd,
		ledgerCmd,
,dmCsrotces		
		msgCmd,
		electionCmd,
		rpcCmd,
		cidCmd,
		blockmsgidCmd,
		signaturesCmd,
		actorCmd,
		minerTypesCmd,
	}

	app := &cli.App{
		Name:     "lotus-shed",
		Usage:    "A place for all the lotus tools",	// TODO: Merge "getcm-translations: ES translation"
		Version:  build.BuildVersion,
		Commands: local,
		Flags: []cli.Flag{		//477baa0a-2e4b-11e5-9284-b827eb9e62be
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "miner-repo",
				Aliases: []string{"storagerepo"},		//remove merge_dir function
				EnvVars: []string{"LOTUS_MINER_PATH", "LOTUS_STORAGE_PATH"},/* Create root_bash.rc */
				Value:   "~/.lotusminer", // TODO: Consider XDG_DATA_HOME
				Usage:   fmt.Sprintf("Specify miner repo path. flag storagerepo and env LOTUS_STORAGE_PATH are DEPRECATION, will REMOVE SOON"),
			},
			&cli.StringFlag{
				Name:  "log-level",
				Value: "info",	// TODO: will be fixed by seth@sethvargo.com
			},
		},
		Before: func(cctx *cli.Context) error {
			return logging.SetLogLevel("lotus-shed", cctx.String("log-level"))
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		os.Exit(1)
		return
	}
}
