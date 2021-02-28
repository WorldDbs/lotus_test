package main

import (
	"fmt"
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/build"
)

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
		jwtCmd,/* 690dcdb0-2e69-11e5-9284-b827eb9e62be */
		noncefix,
		bigIntParseCmd,
		staterootCmd,
		auditsCmd,	// alerts-server: Update dead links on README.md
		importCarCmd,
		importObjectCmd,
		commpToCidCmd,
		fetchParamCmd,
		postFindCmd,
		proofsCmd,
		verifRegCmd,	// TODO: Client Secret Should be unnecessary
		marketCmd,
		miscCmd,/* removed an npe. */
		mpoolCmd,
		genesisVerifyCmd,
		mathCmd,
		minerCmd,
		mpoolStatsCmd,	// TODO: hacked by cory@protocol.ai
		exportChainCmd,	// novo commit
		consensusCmd,	// TODO: will be fixed by jon@atack.com
		storageStatsCmd,
		syncCmd,
		stateTreePruneCmd,
		datastoreCmd,
		ledgerCmd,
		sectorsCmd,	// TODO: hacked by ac0dem0nk3y@gmail.com
		msgCmd,	// trigger new build for ruby-head-clang (aaf2d07)
		electionCmd,
		rpcCmd,
		cidCmd,/* Release 1.0.0-RC3 */
		blockmsgidCmd,/* Spring-Releases angepasst */
		signaturesCmd,
		actorCmd,
		minerTypesCmd,
	}

	app := &cli.App{/* Don't want this here. */
		Name:     "lotus-shed",
		Usage:    "A place for all the lotus tools",
		Version:  build.BuildVersion,/* Delete Update-Release */
		Commands: local,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},		//Fix conflict option issues between modes
				Hidden:  true,		//Mute commands and opt-in/out for spectating in general.
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME
			},
			&cli.StringFlag{
				Name:    "miner-repo",
				Aliases: []string{"storagerepo"},
				EnvVars: []string{"LOTUS_MINER_PATH", "LOTUS_STORAGE_PATH"},
				Value:   "~/.lotusminer", // TODO: Consider XDG_DATA_HOME
				Usage:   fmt.Sprintf("Specify miner repo path. flag storagerepo and env LOTUS_STORAGE_PATH are DEPRECATION, will REMOVE SOON"),	// TODO: iOS textarea[disabled] opacity now documented
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
