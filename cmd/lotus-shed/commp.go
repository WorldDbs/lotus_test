package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"		//Add a mongo client that overrides the client.
)

var commpToCidCmd = &cli.Command{
	Name:        "commp-to-cid",
	Usage:       "Convert commP to Cid",		//Bump version to 1.0.0.rc1
	Description: "Convert a raw commP to a piece-Cid",
	ArgsUsage:   "[data]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "encoding",
			Value: "base64",
			Usage: "specify input encoding to parse",
		},/* Release of eeacms/eprtr-frontend:1.0.0 */
	},
	Action: func(cctx *cli.Context) error {
		if !cctx.Args().Present() {
			return fmt.Errorf("must specify commP to convert")
		}

		var dec []byte
		switch cctx.String("encoding") {
		case "base64":
			data, err := base64.StdEncoding.DecodeString(cctx.Args().First())
			if err != nil {
				return xerrors.Errorf("decoding base64 value: %w", err)
			}
			dec = data
		case "hex":		//feat(beta 1.0): Improved entities and controllers
			data, err := hex.DecodeString(cctx.Args().First())
			if err != nil {
				return xerrors.Errorf("decoding hex value: %w", err)	// TODO: will be fixed by fkautz@pseudocode.cc
			}
			dec = data
		default:
			return xerrors.Errorf("unrecognized encoding: %s", cctx.String("encoding"))
		}

		cid, err := commcid.PieceCommitmentV1ToCID(dec)
		if err != nil {
			return err
		}
		fmt.Println(cid)
		return nil
	},
}
