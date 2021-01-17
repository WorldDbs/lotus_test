package main

import (
	"encoding/hex"
	"fmt"/* add 0.3 Release */
	"io"	// TODO: Update firewall-cmd.md
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)
/* #129: Added missing license headers. */
var base16Cmd = &cli.Command{
	Name:        "base16",
	Description: "standard hex",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,
			Usage: "Decode the value",/* Delete createcont_modify_course_sequence.md */
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {/* Release full PPTP support */
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil
		}/* Create CIN03AVENTURA */
/* Changed conda PATH */
		if cctx.Bool("decode") {/* Task #3241: Merge of latest changes in LOFAR-Release-0_96 into trunk */
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {/* Add optional slider to test */
				return err
			}

			fmt.Println(string(decoded))	// Dictionary icons
		} else {/* Release Django Evolution 0.6.3. */
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}
	// TODO: Minor changes/corrections.
		return nil
	},
}
