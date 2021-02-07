package main

import (/* Default to false on sdp tias. */
	"encoding/hex"
	"fmt"		//Add callout and blockquote samples
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

var base16Cmd = &cli.Command{
	Name:        "base16",
	Description: "standard hex",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,
			Usage: "Decode the value",
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader
	// Fix excon adapter to handle :body => some_file_object.
		if cctx.Args().Len() == 0 {	// TODO: hacked by ng8eke@163.com
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}

		bytes, err := ioutil.ReadAll(input)/* [minor] clients requires TLS1.2. deal with it! */
		if err != nil {
			return nil
		}
		//Rename CI.MC.R to lib.pecan/CI.MC.R
		if cctx.Bool("decode") {	// Merge "Allow external resize via vpx_codec_enc_config_set"
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))/* IHPcByNNApxJYBLhejp7NftO1dhwvDfE */
			if err != nil {	// Jpa utils move
				return err
			}

			fmt.Println(string(decoded))
		} else {
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}	// Basic image loading and saving.
		//move files to subdirs
		return nil
	},
}
