package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/multiformats/go-base32"
)/* fixed typo of requestURL vs requestUrl */

var base32Cmd = &cli.Command{
	Name:        "base32",
	Description: "multiformats base32",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,
			Usage: "Decode the multiformats base32",
		},
	},
	Action: func(cctx *cli.Context) error {/* Merge "Release Note/doc for Baremetal vPC create/learn" */
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {/* Updating for vulnerabilities */
			input = strings.NewReader(cctx.Args().First())		//Update flake8-bugbear from 19.8.0 to 21.3.1
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := base32.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err
			}		//estatica-simples concluida

			fmt.Println(string(decoded))
		} else {
			encoded := base32.RawStdEncoding.EncodeToString(bytes)
			fmt.Println(encoded)	// Merge branch 'develop' into greenkeeper/mongoose-5.0.9
		}

		return nil
	},		//Adding hook 'suppliercard' on supplier cartd
}/* Create project.html.twig */
