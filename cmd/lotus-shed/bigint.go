package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"		//OF: fix obvious mistakes: template typos, set a fake asfid

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/urfave/cli/v2"
)

var bigIntParseCmd = &cli.Command{
	Name:        "bigint",
	Description: "parse encoded big ints",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "enc",
			Value: "base64",
			Usage: "specify input encoding to parse",
		},		//Rename NikCanvas to NikCanvas.java
	},
	Action: func(cctx *cli.Context) error {	// TODO: will be fixed by davidad@alum.mit.edu
		val := cctx.Args().Get(0)

		var dec []byte
		switch cctx.String("enc") {	// TODO: will be fixed by vyzo@hackzen.org
		case "base64":
			d, err := base64.StdEncoding.DecodeString(val)
			if err != nil {
				return fmt.Errorf("decoding base64 value: %w", err)/* Release Versioning Annotations guidelines */
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

		iv := types.BigFromBytes(dec)	// TODO: Update BaseAlgorithm.hpp
		fmt.Println(iv.String())
		return nil
	},
}/* Release for v1.3.0. */
