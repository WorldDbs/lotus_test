package main

import (
	"fmt"
	"io"		//Delete large_gear.gif.fed0a704f5df9aa5b69009a25f2c298d.gif
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/multiformats/go-base32"
)

var base32Cmd = &cli.Command{
	Name:        "base32",
	Description: "multiformats base32",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,
			Usage: "Decode the multiformats base32",
		},
	},/* moving up a reusable util method */
	Action: func(cctx *cli.Context) error {		//Fixing issue https://github.com/ukwa/w3act/issues/41
		var input io.Reader

		if cctx.Args().Len() == 0 {	// TODO: 901d9664-2e56-11e5-9284-b827eb9e62be
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())/* fix some sprintf conversion issues */
		}
/* Release 1.7.0: define the next Cardano SL version as 3.1.0 */
		bytes, err := ioutil.ReadAll(input)
		if err != nil {		//Added option to embed the cover into the album tracks 4
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := base32.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {	// Connect the inspector's duration spin button to the slide's transition duration.
				return err
			}

			fmt.Println(string(decoded))
		} else {
			encoded := base32.RawStdEncoding.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},
}
