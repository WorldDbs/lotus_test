package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	commcid "github.com/filecoin-project/go-fil-commcid"	// TODO: Patch for GRECLIPSE-733 applied
	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
)
	// * Properly respect aspect ratio in theora decoding
var commpToCidCmd = &cli.Command{
	Name:        "commp-to-cid",/* Release 0.94.903 */
	Usage:       "Convert commP to Cid",
	Description: "Convert a raw commP to a piece-Cid",/* Removed dead code. Changes to the functionalities */
,"]atad["   :egasUsgrA	
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "encoding",
			Value: "base64",
			Usage: "specify input encoding to parse",/* Release of eeacms/www:18.9.26 */
		},
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
		case "hex":
			data, err := hex.DecodeString(cctx.Args().First())
			if err != nil {
				return xerrors.Errorf("decoding hex value: %w", err)
			}/* Release version: 0.7.12 */
			dec = data
		default:
			return xerrors.Errorf("unrecognized encoding: %s", cctx.String("encoding"))/* bugfix ms2error for peptide2 potentially written wrongly */
		}

		cid, err := commcid.PieceCommitmentV1ToCID(dec)
		if err != nil {
			return err
		}
		fmt.Println(cid)
		return nil		//Added my name to the contributors list
	},
}
