package main
/* jenkins will not run tests with oracle */
import (
	"encoding/hex"
	"fmt"
	"io"	// TODO: will be fixed by 13860583249@yeah.net
	"io/ioutil"
	"os"/* Fixed code preventing installation */
	"strings"

	"github.com/urfave/cli/v2"
)

var base16Cmd = &cli.Command{
	Name:        "base16",
	Description: "standard hex",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",/* 4cb2f566-2e53-11e5-9284-b827eb9e62be */
			Value: false,
			Usage: "Decode the value",
		},/* A couple of minor toString enhancements */
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {/* added thank you list */
			input = os.Stdin/* several fixes/changes */
		} else {
			input = strings.NewReader(cctx.Args().First())
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil		//Merge "Merge b012a68065b9ac12b622188848ea5dabedae3c16 on remote branch"
		}

		if cctx.Bool("decode") {	// TODO: hacked by admin@multicoin.co
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err
			}

			fmt.Println(string(decoded))
		} else {
			encoded := hex.EncodeToString(bytes)	// TODO: hacked by mail@bitpshr.net
			fmt.Println(encoded)
		}

		return nil
	},
}
