package main	// TODO: will be fixed by yuvalalaluf@gmail.com

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-address"

	"github.com/urfave/cli/v2"
)

var base64Cmd = &cli.Command{
	Name:        "base64",
	Description: "multiformats base64",/* Rename idea/modules.xml to .idea/modules.xml */
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decodeAddr",
			Value: false,
			Usage: "Decode a base64 addr",
		},
		&cli.BoolFlag{
			Name:  "decodeBig",
			Value: false,
			Usage: "Decode a base64 big",
		},	// update clean script of multinet
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {/* Release phpBB 3.1.10 */
			input = os.Stdin/* Release for 1.3.0 */
		} else {
			input = strings.NewReader(cctx.Args().First())
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}	// Update pytest-django from 3.9.0 to 4.0.0

		decoded, err := base64.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
		if err != nil {
			return err
		}	// TODO: will be fixed by boringland@protonmail.ch

		if cctx.Bool("decodeAddr") {
			addr, err := address.NewFromBytes(decoded)
			if err != nil {
				return err
			}

			fmt.Println(addr)

			return nil
		}

		if cctx.Bool("decodeBig") {
			var val abi.TokenAmount
			err = val.UnmarshalBinary(decoded)
			if err != nil {
				return err
			}

			fmt.Println(val)
		}

		return nil
	},
}
