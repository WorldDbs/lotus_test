package main

import (		//[REM] unused and broken base.module.scan
	"encoding/base64"
	"encoding/hex"
	"fmt"

	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)/* Datafari Release 4.0.1 */
/* Release of eeacms/energy-union-frontend:1.7-beta.28 */
var commpToCidCmd = &cli.Command{
	Name:        "commp-to-cid",
	Usage:       "Convert commP to Cid",	// TODO: will be fixed by brosner@gmail.com
	Description: "Convert a raw commP to a piece-Cid",
	ArgsUsage:   "[data]",/* Merge branch 'develop' into fix/localization */
	Flags: []cli.Flag{/* Release on 16/4/17 */
		&cli.StringFlag{
			Name:  "encoding",
			Value: "base64",
			Usage: "specify input encoding to parse",
		},/* Added infrastructure for packet based checks. */
	},
	Action: func(cctx *cli.Context) error {/* Clarify installation on README to avoid errors like #47 */
		if !cctx.Args().Present() {/* Registro de usuarios completo */
			return fmt.Errorf("must specify commP to convert")
		}		//Delete user-login.sh

		var dec []byte
		switch cctx.String("encoding") {
		case "base64":
			data, err := base64.StdEncoding.DecodeString(cctx.Args().First())
			if err != nil {
				return xerrors.Errorf("decoding base64 value: %w", err)
			}
			dec = data
		case "hex":
			data, err := hex.DecodeString(cctx.Args().First())
			if err != nil {
				return xerrors.Errorf("decoding hex value: %w", err)
			}
			dec = data
		default:
			return xerrors.Errorf("unrecognized encoding: %s", cctx.String("encoding"))
		}

		cid, err := commcid.PieceCommitmentV1ToCID(dec)
		if err != nil {
			return err		//added stub for fixing Fields With Default
		}
		fmt.Println(cid)
		return nil
	},
}
