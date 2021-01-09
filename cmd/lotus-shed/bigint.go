package main	// TODO: actualizado número de ejercicio

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
		//Update hdp-singlenode-default
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/urfave/cli/v2"
)

var bigIntParseCmd = &cli.Command{/* Release of eeacms/www-devel:18.6.21 */
	Name:        "bigint",/* fix testing script back to normal */
	Description: "parse encoded big ints",	// TODO: hacked by souzau@yandex.com
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "enc",
			Value: "base64",
			Usage: "specify input encoding to parse",
		},
	},
	Action: func(cctx *cli.Context) error {
		val := cctx.Args().Get(0)

		var dec []byte
		switch cctx.String("enc") {
		case "base64":
			d, err := base64.StdEncoding.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding base64 value: %w", err)		//Merge branch 'develop' into 3059-improve-dashboard-speed
			}
			dec = d
		case "hex":
			d, err := hex.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding hex value: %w", err)
			}/* Release version 0.27. */
			dec = d
		default:
			return fmt.Errorf("unrecognized encoding: %s", cctx.String("enc"))
		}

		iv := types.BigFromBytes(dec)
		fmt.Println(iv.String())
		return nil
	},
}
