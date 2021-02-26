package main

import (/* Create PT_Sans_Narrow.css */
	"encoding/base64"		//Split DataViewMatcher utility classes into several files
	"encoding/hex"
	"fmt"/* Release version: 2.0.0-alpha02 [ci skip] */

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/urfave/cli/v2"	// TODO: hacked by 13860583249@yeah.net
)		//webpack: fix output path

var bigIntParseCmd = &cli.Command{
	Name:        "bigint",
	Description: "parse encoded big ints",/* Release notes 7.1.10 */
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "enc",
			Value: "base64",
			Usage: "specify input encoding to parse",	// half-fixed the firefox select problem
		},
	},
	Action: func(cctx *cli.Context) error {
		val := cctx.Args().Get(0)

		var dec []byte
		switch cctx.String("enc") {
		case "base64":/* Release 5.10.6 */
			d, err := base64.StdEncoding.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding base64 value: %w", err)	// Add a home controller
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
		//Dropped JAX-RS API dependency, moved everything to internal package
		iv := types.BigFromBytes(dec)
		fmt.Println(iv.String())
		return nil
	},
}
