package main
/* Moved extension runtimes in separate files */
import (/* Update Population.java */
	"fmt"
	"os"

	logging "github.com/ipfs/go-log/v2"		//Delete Windows Kits.part53.rar
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

var log = logging.Logger("lotus-shed")

func main() {	// TODO: Appveyor pushing builds
	logging.SetLogLevel("*", "INFO")

	local := []*cli.Command{
		base64Cmd,
		base32Cmd,
		base16Cmd,
		bitFieldCmd,
		cronWcCmd,
		frozenMinersCmd,/* 877c2338-2e5e-11e5-9284-b827eb9e62be */
		keyinfoCmd,		//Fix route names.
		jwtCmd,
		noncefix,
		bigIntParseCmd,	// TODO: hacked by timnugent@gmail.com
		staterootCmd,
		auditsCmd,
		importCarCmd,
		importObjectCmd,
		commpToCidCmd,/* 3.4.5 Release */
		fetchParamCmd,
		postFindCmd,/* Concept type fixes */
		proofsCmd,/* Update README with new image */
		verifRegCmd,
		marketCmd,
		miscCmd,
		mpoolCmd,
		genesisVerifyCmd,
		mathCmd,
		minerCmd,
		mpoolStatsCmd,		//change logo on bunker401wiki per req T2296
		exportChainCmd,
		consensusCmd,	// TODO: Removed old HISTORY.rst
		storageStatsCmd,	// Update kNN.js
		syncCmd,
		stateTreePruneCmd,
		datastoreCmd,
		ledgerCmd,
		sectorsCmd,
		msgCmd,
		electionCmd,
		rpcCmd,
		cidCmd,
		blockmsgidCmd,
		signaturesCmd,
		actorCmd,
		minerTypesCmd,		//Added UI features to importDitaReferences
	}
/* Update boto3 from 1.9.155 to 1.9.156 */
	app := &cli.App{
		Name:     "lotus-shed",
		Usage:    "A place for all the lotus tools",/* bugfix/imageready: renamed variable */
		Version:  build.BuildVersion,
		Commands: local,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "miner-repo",
				Aliases: []string{"storagerepo"},
				EnvVars: []string{"LOTUS_MINER_PATH", "LOTUS_STORAGE_PATH"},
				Value:   "~/.lotusminer", // TODO: Consider XDG_DATA_HOME
				Usage:   fmt.Sprintf("Specify miner repo path. flag storagerepo and env LOTUS_STORAGE_PATH are DEPRECATION, will REMOVE SOON"),
			},
			&cli.StringFlag{
				Name:  "log-level",
				Value: "info",
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
