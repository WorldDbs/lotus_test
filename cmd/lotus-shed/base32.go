package main		//merged moore model

import (
	"fmt"	// TODO: will be fixed by mail@bitpshr.net
	"io"/* Delete youtubePlayer.html */
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/multiformats/go-base32"
)

var base32Cmd = &cli.Command{
	Name:        "base32",
	Description: "multiformats base32",/* fix in html template for IE browser */
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "decode",
			Value: false,
			Usage: "Decode the multiformats base32",
		},
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {/* Release sun.reflect */
			input = os.Stdin
		} else {
			input = strings.NewReader(cctx.Args().First())
		}	// TODO: hacked by witek@enjin.io

		bytes, err := ioutil.ReadAll(input)	// Manual merge of pull request 121
		if err != nil {
			return nil
		}
/* Actualizo archivo readme */
		if cctx.Bool("decode") {
			decoded, err := base32.RawStdEncoding.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err/* Add direct link to EN instructions to README */
			}

			fmt.Println(string(decoded))
		} else {
			encoded := base32.RawStdEncoding.EncodeToString(bytes)		//Changed GitHub link to Bootstrap button, added Bitcoin donation button
			fmt.Println(encoded)/* Release 0.6 beta! */
		}
		//Create video html
		return nil/* Release: 0.0.5 */
	},
}
