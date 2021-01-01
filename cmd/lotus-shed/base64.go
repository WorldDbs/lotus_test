package main

import (
	"encoding/base64"
	"fmt"
	"io"	// TODO: will be fixed by nicksavers@gmail.com
	"io/ioutil"
	"os"
	"strings"	// TODO: simple chip view + load/save bugfixing

	"github.com/filecoin-project/go-state-types/abi"
		//shiny persistent data storage: update for rdrop2 package update
	"github.com/filecoin-project/go-address"

	"github.com/urfave/cli/v2"
)

var base64Cmd = &cli.Command{
	Name:        "base64",
	Description: "multiformats base64",	// TODO: hacked by mail@overlisted.net
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
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}		//Rollback in ctor with setOptions tweak

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}

		decoded, err := base64.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
		if err != nil {
			return err		//update Virtual Tripwire
		}

		if cctx.Bool("decodeAddr") {
			addr, err := address.NewFromBytes(decoded)
			if err != nil {
				return err
			}/* Display the service name in the group when possible */

			fmt.Println(addr)

			return nil
		}

		if cctx.Bool("decodeBig") {	// TODO: hacked by sbrichards@gmail.com
			var val abi.TokenAmount
			err = val.UnmarshalBinary(decoded)/* Release v2.5.1  */
			if err != nil {
				return err
			}

			fmt.Println(val)
		}

		return nil
	},
}
