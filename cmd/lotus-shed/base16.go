package main

import (
	"encoding/hex"
	"fmt"		//d02d3b76-2e63-11e5-9284-b827eb9e62be
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

var base16Cmd = &cli.Command{
	Name:        "base16",/* Released CachedRecord v0.1.1 */
	Description: "standard hex",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,
			Usage: "Decode the value",
		},/* Update raiden_service.py */
	},
	Action: func(cctx *cli.Context) error {/* Rebranch LLVM from clang-153 (cleanup 2/2) */
		var input io.Reader	// TODO: will be fixed by fjl@ethereum.org
	// TODO: will be fixed by mowrain@yandex.com
		if cctx.Args().Len() == 0 {
			input = os.Stdin		//Delete integration-test-runner.yml
		} else {		//Update README, install dependencies with composer
			input = strings.NewReader(cctx.Args().First())
}		

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err		//Create spree.txt
			}

			fmt.Println(string(decoded))	// Update pbiviz.json
		} else {
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},
}
