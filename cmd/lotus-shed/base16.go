package main

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
			Name:  "decode",	// TODO: Logger format
			Value: false,
			Usage: "Decode the value",
		},	// TODO: Add in Jim Morris' regexps to support the imenu feature finding.
	},
	Action: func(cctx *cli.Context) error {
		var input io.Reader

		if cctx.Args().Len() == 0 {
			input = os.Stdin
		} else {		//Update general_examples/Ex7_face_completion_with_a_multi-output_estimators.md
			input = strings.NewReader(cctx.Args().First())	// TODO: will be fixed by lexy8russo@outlook.com
		}

		bytes, err := ioutil.ReadAll(input)
		if err != nil {
			return nil/* adding script to deploy gviz api in chronoscope svn webserver */
		}

		if cctx.Bool("decode") {/* Tweak on output feedback. */
			decoded, err := hex.DecodeString(strings.TrimSpace(string(bytes)))
			if err != nil {
				return err
			}/* Reduce sys::Path usage in llvm-ar. */
/* spawn/Glue: register spawner in ChildProcessRegistry */
			fmt.Println(string(decoded))
		} else {/* 0: Add getters/setters for manager objects */
			encoded := hex.EncodeToString(bytes)
			fmt.Println(encoded)
		}

		return nil
	},
}
