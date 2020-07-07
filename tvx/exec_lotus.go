package main

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/lib/blockstore"
	"github.com/ipld/go-car"
	"github.com/urfave/cli/v2"

	"github.com/filecoin-project/oni/tvx/lotus"
	"github.com/filecoin-project/oni/tvx/schema"
)

var execLotusFlags struct {
	file string
}

var execLotusCmd = &cli.Command{
	Name:        "exec-lotus",
	Description: "execute a test vector against Lotus",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:        "file",
			Usage:       "input file",
			Destination: &execLotusFlags.file,
		},
	},
	Action: runExecLotus,
}

func runExecLotus(_ *cli.Context) error {
	switch {
	case execLotusFlags.file != "":
		file, err := os.Open(execLotusFlags.file)
		if err != nil {
			return fmt.Errorf("failed to open test vector: %w", err)
		}

		var (
			dec = json.NewDecoder(file)
			tv  schema.TestVector
		)

		if err = dec.Decode(&tv); err != nil {
			return fmt.Errorf("failed to decode test vector: %w", err)
		}

		return executeTestVector(tv)
	default:
		dec := json.NewDecoder(os.Stdin)
		for {
			var tv schema.TestVector

			err := dec.Decode(&tv)
			if err == io.EOF {
				return nil
			}
			if err != nil {
				return err
			}

			err = executeTestVector(tv)
			if err != nil {
				return err
			}
		}
	}
}

func executeTestVector(tv schema.TestVector) error {
	fmt.Println("executing test vector:", tv.Meta.Desc)
	switch tv.Class {
	case "message":
		var (
			ctx   = context.Background()
			epoch = tv.Pre.Epoch
			root  = tv.Pre.StateTree.RootCID
		)

		bs := blockstore.NewTemporary()

		buf := bytes.NewReader(tv.CAR)
		gr, err := gzip.NewReader(buf)
		if err != nil {
			return err
		}
		defer gr.Close()

		header, err := car.LoadCar(bs, gr)
		if err != nil {
			return fmt.Errorf("failed to load state tree car from test vector: %w", err)
		}

		fmt.Println("roots: ", header.Roots)

		driver := lotus.NewDriver(ctx)

		for i, m := range tv.ApplyMessages {
			fmt.Printf("decoding message %v\n", i)
			msg, err := types.DecodeMessage(m.Bytes)
			if err != nil {
				return err
			}

			// add an epoch if we have set one
			if m.Epoch != nil {
				epoch = *m.Epoch
			}

			fmt.Printf("executing message %v\n", i)
			var ret *vm.ApplyRet
			ret, root, err = driver.ExecuteMessage(msg, root, bs, epoch)
			if err != nil {
				return err
			}

			if expected, actual := tv.Post.Receipts[i].ExitCode, ret.ExitCode; expected != actual {
				return fmt.Errorf("exit code of msg %d did not match; expected: %s, got: %s", i, expected, actual)
			}
			if expected, actual := tv.Post.Receipts[i].GasUsed, ret.GasUsed; expected != actual {
				return fmt.Errorf("gas used of msg %d did not match; expected: %d, got: %d", i, expected, actual)
			}

			// TODO assert return value
			fmt.Printf("✅  message %d passed expectations\n", i)
		}

		if root != tv.Post.StateTree.RootCID {
			return fmt.Errorf("wrong post root cid; expected %v , but got %v", tv.Post.StateTree.RootCID, root)
		}

		return nil

	default:
		return fmt.Errorf("test vector class not supported")
	}
}
