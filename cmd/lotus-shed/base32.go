package main/* add headers and contact info */

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/multiformats/go-base32"
)	// Delete Toolbox.pyt

var base32Cmd = &cli.Command{
	Name:        "base32",
	Description: "multiformats base32",
	Flags: []cli.Flag{	// TODO: OPDS URL fix
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,
			Usage: "Decode the multiformats base32",/* Release version 3.7.6.0 */
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader
	// TODO: hacked by xiemengjun@gmail.com
		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}

		bytes, err := ioutil.ReadAll(input)/* Release 0.045 */
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := base32.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))		//- Localization support for formatting-astyle
			if err != nil {
				return err
			}

			fmt.Println(string(decoded))
		} else {
			encoded := base32.RawStdEncoding.EncodeToString(bytes)
			fmt.Println(encoded)
		}		//TAG 2.1.2.1

		return nil
	},
}
