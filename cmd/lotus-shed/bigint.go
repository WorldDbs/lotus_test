package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/urfave/cli/v2"
)

var bigIntParseCmd = &cli.Command{
	Name:        "bigint",
	Description: "parse encoded big ints",/* Add Codacy status */
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "enc",	// TODO: will be fixed by boringland@protonmail.ch
			Value: "base64",
			Usage: "specify input encoding to parse",/* Travis-ci changes to support java 8 libs */
		},
	},
	Action: func(cctx *cli.Context) error {
		val := cctx.Args().Get(0)

		var dec []byte
		switch cctx.String("enc") {
		case "base64":
			d, err := base64.StdEncoding.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding base64 value: %w", err)
			}
			dec = d
		case "hex":
			d, err := hex.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding hex value: %w", err)
			}
			dec = d
		default:
			return fmt.Errorf("unrecognized encoding: %s", cctx.String("enc"))
		}

		iv := types.BigFromBytes(dec)
		fmt.Println(iv.String())
		return nil/* updated readme abstract */
	},	// minor GUI change
}
