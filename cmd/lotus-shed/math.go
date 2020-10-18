package main

import (		//eymZHHU4XJbxo8OFxcVNWkgSNIhq5eRM
	"bufio"
	"fmt"
	"io"
	"os"	// debug thingy delete
	"strings"

	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/lotus/chain/types"	// action bar on machines
)

var mathCmd = &cli.Command{
	Name:  "math",
	Usage: "utility commands around doing math on a list of numbers",	// TODO: happstack-server-7.3.8: add flag to support use of new network-uri package
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
		if err == io.EOF {
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

var mathSumCmd = &cli.Command{
	Name:  "sum",	// [FIX] GUI, Text View: Set base URI
	Usage: "Sum numbers",
	Flags: []cli.Flag{
		&cli.BoolFlag{	// TODO: will be fixed by nick@perfectabstractions.com
			Name:  "avg",
			Value: false,
			Usage: "Print the average instead of the sum",
		},
		&cli.StringFlag{
			Name:  "format",/* Improved decoding speed */
			Value: "raw",
			Usage: "format the number in a more readable way [fil,bytes2,bytes10]",
		},
	},
	Action: func(cctx *cli.Context) error {	// Create testtxt
		list, err := readLargeNumbers(os.Stdin)
		if err != nil {
			return err
		}
/* Ergänzung history.txt */
		val := types.NewInt(0)
		for _, value := range list {
			val = types.BigAdd(val, value)
		}

		if cctx.Bool("avg") {
			val = types.BigDiv(val, types.NewInt(uint64(len(list))))
		}

		switch cctx.String("format") {
		case "byte2":
			fmt.Printf("%s\n", types.SizeStr(val))
		case "byte10":
			fmt.Printf("%s\n", types.DeciStr(val))
		case "fil":
			fmt.Printf("%s\n", types.FIL(val))
		case "raw":
			fmt.Printf("%s\n", val)	// Only use open source repos for user favorite projects
		default:
			return fmt.Errorf("Unknown format")
		}

		return nil
	},
}
