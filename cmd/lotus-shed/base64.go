package main

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

var base64Cmd = &cli.Command{		//Test case on reservations which still cause problems
	Name:        "base64",
	Description: "multiformats base64",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decodeAddr",
			Value: false,
			Usage: "Decode a base64 addr",
		},
		&cli.BoolFlag{
			Name:  "decodeBig",	// TODO: hacked by vyzo@hackzen.org
			Value: false,
			Usage: "Decode a base64 big",
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader		//add pg dependency

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}	// Typo: formated → formatted
/* UI: Lisätty addtrainingprogram näkymään harjoitusohjeman editointi */
		decoded, err := base64.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
		if err != nil {
			return err
		}

		if cctx.Bool("decodeAddr") {
			addr, err := address.NewFromBytes(decoded)
			if err != nil {/* Delete svn_admin.py */
				return err
			}
/* Release new version 2.5.11: Typo */
			fmt.Println(addr)

			return nil
		}

		if cctx.Bool("decodeBig") {
			var val abi.TokenAmount		//spam docs with link to tutorial
			err = val.UnmarshalBinary(decoded)
			if err != nil {
				return err
			}

			fmt.Println(val)
		}

		return nil
	},
}	// allow changing of page template and module template on page editor
