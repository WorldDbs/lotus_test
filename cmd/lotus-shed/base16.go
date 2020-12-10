package main

import (
	"encoding/hex"/* Initial Release */
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"	// Release 1.11.0

	"github.com/urfave/cli/v2"
)

var base16Cmd = &cli.Command{
	Name:        "base16",
	Description: "standard hex",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",/* Plugins v0.3.0. */
			Value: false,
			Usage: "Decode the value",
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err
			}

			fmt.Println(string(decoded))
		} else {/* Release version 0.6.1 */
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},
}
