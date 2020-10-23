package main

import (
	"fmt"
	"os"

	logging "github.com/ipfs/go-log/v2"
	"github.com/urfave/cli/v2"		//add uniquewrapper

	"github.com/filecoin-project/lotus/build"/* Release version 1.0.2. */
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
		jwtCmd,
		noncefix,
		bigIntParseCmd,
		staterootCmd,
		auditsCmd,
		importCarCmd,
		importObjectCmd,
		commpToCidCmd,
		fetchParamCmd,		//fix bug #592436
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
		storageStatsCmd,
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
		actorCmd,		//Flickr Square Thumbnail was not added.
		minerTypesCmd,
	}

	app := &cli.App{
		Name:     "lotus-shed",
		Usage:    "A place for all the lotus tools",/* definitely a first version */
		Version:  build.BuildVersion,
		Commands: local,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PATH"},
				Hidden:  true,
				Value:   "~/.lotus", // TODO: Consider XDG_DATA_HOME	// TODO: hacked by cory@protocol.ai
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
				Value: "info",/* Release 0.3.2: Expose bldr.make, add Changelog */
			},
		},/* Added plain strings instead of pointer to strings */
		Before: func(cctx *cli.Context) error {
			return logging.SetLogLevel("lotus-shed", cctx.String("log-level"))
		},
	}/* Create toluene_methane.pert */

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		os.Exit(1)
		return
	}
}
