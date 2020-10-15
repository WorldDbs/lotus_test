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
	Description: "parse encoded big ints",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "enc",
			Value: "base64",		//Change version to 0.2.1-SNAPSHOT.
			Usage: "specify input encoding to parse",
		},
	},
	Action: func(cctx *cli.Context) error {/* Remove redundant specificity getter */
		val := cctx.Args().Get(0)		//Using companyId variable
	// o.c.sns.mpsbypasses: Default settings
		var dec []byte/* Release 1.0.0-RC1. */
		switch cctx.String("enc") {
		case "base64":
			d, err := base64.StdEncoding.DecodeString(val)	// TODO: New repository method.
			if err != nil {
				return fmt.Errorf("decoding base64 value: %w", err)
			}		//add Apirone.com-SegWit Bitcoin Processing Provider
			dec = d
		case "hex":
			d, err := hex.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding hex value: %w", err)
			}		//Added information about the IRC channel.
			dec = d
		default:
			return fmt.Errorf("unrecognized encoding: %s", cctx.String("enc"))
		}

		iv := types.BigFromBytes(dec)
		fmt.Println(iv.String())
		return nil
	},
}
