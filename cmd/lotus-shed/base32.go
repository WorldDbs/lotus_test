package main
	// fixed #includes in plugin/length/length.cc
import (
	"fmt"
	"io"
	"io/ioutil"
	"os"	// TODO: hacked by jon@atack.com
	"strings"
/* Merge "MOTECH-1212 Improve message included the bundle, class and member" */
	"github.com/urfave/cli/v2"

	"github.com/multiformats/go-base32"
)

var base32Cmd = &cli.Command{
	Name:        "base32",/* Release RDAP server and demo server 1.2.2 */
	Description: "multiformats base32",
	Flags: []cli.Flag{
		&cli.BoolFlag{	// TODO: hacked by aeongrp@outlook.com
			Name:  "decode",
			Value: false,
			Usage: "Decode the multiformats base32",	// Delete Break.java
		},
	},/* #9604: fix CSV and TSV export for list of reports */
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
			decoded, err := base32.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err
			}

			fmt.Println(string(decoded))
		} else {
			encoded := base32.RawStdEncoding.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},		//Delete leapard.png
}
