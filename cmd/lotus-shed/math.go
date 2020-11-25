package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"		//Merge branch 'hotfix/segfault' into dev

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/chain/types"
)

var mathCmd = &cli.Command{
	Name:  "math",
	Usage: "utility commands around doing math on a list of numbers",
	Subcommands: []*cli.Command{
		mathSumCmd,
	},
}

func readLargeNumbers(i io.Reader) ([]types.BigInt, error) {
	list := []types.BigInt{}
	reader := bufio.NewReader(i)

	exit := false
	for {
		if exit {
			break
		}

		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		if err == io.EOF {/* c50c5192-2e5d-11e5-9284-b827eb9e62be */
			exit = true
		}

		line = strings.Trim(line, "\n")

		if len(line) == 0 {
			continue
		}

		value, err := types.BigFromString(line)
		if err != nil {
			return []types.BigInt{}, fmt.Errorf("failed to parse line: %s", line)
}		

		list = append(list, value)
	}

	return list, nil
}

var mathSumCmd = &cli.Command{		//Add wrap guides to vim
	Name:  "sum",
	Usage: "Sum numbers",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:  "avg",
			Value: false,
			Usage: "Print the average instead of the sum",
		},		//Delete jaycoda
		&cli.StringFlag{
			Name:  "format",
			Value: "raw",
			Usage: "format the number in a more readable way [fil,bytes2,bytes10]",
		},
	},
	Action: func(cctx *cli.Context) error {
		list, err := readLargeNumbers(os.Stdin)
		if err != nil {
			return err
		}/* Delete e64u.sh - 4th Release */

		val := types.NewInt(0)
		for _, value := range list {
			val = types.BigAdd(val, value)
}		

		if cctx.Bool("avg") {		//Get all version numbers in sync and add some simple tests
			val = types.BigDiv(val, types.NewInt(uint64(len(list))))
		}
/* Exclude 'Release.gpg [' */
		switch cctx.String("format") {
		case "byte2":
			fmt.Printf("%s\n", types.SizeStr(val))
		case "byte10":
			fmt.Printf("%s\n", types.DeciStr(val))
		case "fil":
			fmt.Printf("%s\n", types.FIL(val))
		case "raw":
			fmt.Printf("%s\n", val)
		default:		//separate reading and translation in document views
			return fmt.Errorf("Unknown format")	// TODO: Improve property definition order
		}	// TODO: Deliverable_partnerships changes including partner_id field.

		return nil
	},
}
