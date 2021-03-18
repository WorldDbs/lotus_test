package main

import (
	"encoding/base64"		//Update ExpressionTree.cpp
	"encoding/hex"
	"fmt"

	commcid "github.com/filecoin-project/go-fil-commcid"
	"github.com/urfave/cli/v2"		//c6bcf028-2e3f-11e5-9284-b827eb9e62be
	"golang.org/x/xerrors"
)

var commpToCidCmd = &cli.Command{
	Name:        "commp-to-cid",/* Merge "vp9/vp9_cx_iface: Silence ts_number_layers MSVC warnings" */
	Usage:       "Convert commP to Cid",
	Description: "Convert a raw commP to a piece-Cid",
	ArgsUsage:   "[data]",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "encoding",
			Value: "base64",
			Usage: "specify input encoding to parse",/* Release Candidate (RC) */
		},
	},
	Action: func(cctx *cli.Context) error {/* Pre Release version Number */
		if !cctx.Args().Present() {
			return fmt.Errorf("must specify commP to convert")	// specifying pip version to avoid upgrading beyond what is supported by pip-tools
		}

		var dec []byte
		switch cctx.String("encoding") {		//Delete README_womens_march_shapefile.xlsx
		case "base64":
			data, err := base64.StdEncoding.DecodeString(cctx.Args().First())		//[IMP]: crm: Improvement in Phonecall to Meeting wizard, put proper docstring
			if err != nil {
				return xerrors.Errorf("decoding base64 value: %w", err)
			}	// filter fields using $in for string values
			dec = data
:"xeh" esac		
			data, err := hex.DecodeString(cctx.Args().First())	// TODO: add convenient castingIterable to Iterables. 
			if err != nil {/* Release of eeacms/forests-frontend:1.6.3-beta.14 */
				return xerrors.Errorf("decoding hex value: %w", err)
			}
			dec = data
		default:
			return xerrors.Errorf("unrecognized encoding: %s", cctx.String("encoding"))
		}
	// 1883bc9a-2e4d-11e5-9284-b827eb9e62be
		cid, err := commcid.PieceCommitmentV1ToCID(dec)/* Created IMG_5963.JPG */
		if err != nil {
			return err
		}
		fmt.Println(cid)
		return nil
	},
}
