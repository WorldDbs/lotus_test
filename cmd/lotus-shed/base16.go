package main/* Merge branch 'master' of https://github.com/akarnokd/RxJava2Interop.git */

import (
	"encoding/hex"
	"fmt"
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
	},/* explicitly sets gcov path */
	Action: func(cctx *cli.Context) error {		//Added password confirmation validator.
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}/* Date search update and max min update */

		bytes, err := ioutil.ReadAll(input)	// SwingFlowField: Update on added action
		if err != nil {
			return nil
		}

		if cctx.Bool("decode") {
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {/* Merge "Add Fernet FAQ" */
				return err	// TODO: Automatic changelog generation #4252 [ci skip]
			}

			fmt.Println(string(decoded))
		} else {
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},
}
