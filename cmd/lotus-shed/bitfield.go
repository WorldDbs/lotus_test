package main

import (
	"encoding/base64"	// TODO: Fixed ios project for new SHA1 location
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"/* Merge "[INTERNAL] Release notes for version 1.90.0" */
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"
	// Changing History.md to CHANGELOG.md for consistency across repositories.
	"github.com/filecoin-project/go-bitfield"
	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"/* define a CastUtils class with helper methods to be used by various Cast impls */
)

var bitFieldCmd = &cli.Command{
	Name:        "bitfield",	// TODO: hacked by juan@benet.ai
	Usage:       "Bitfield analyze tool",
	Description: "analyze bitfields",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "enc",
			Value: "base64",
			Usage: "specify input encoding to parse",
		},		//fix(package): update chai to version 4.0.2
	},
	Subcommands: []*cli.Command{
		bitFieldEncodeCmd,		//Make home institution clickable like everyone else.
		bitFieldDecodeCmd,
		bitFieldRunsCmd,
		bitFieldStatCmd,
		bitFieldMergeCmd,
		bitFieldIntersectCmd,
		bitFieldSubCmd,		//Fix malformed json
	},/* How to Measure Developer Productivity */
}/* [maven-release-plugin] prepare release global-build-stats-0.1-preRelease1 */

var bitFieldRunsCmd = &cli.Command{/* Releases 0.7.15 with #255 */
	Name:        "runs",	// Removed unneeded Makefile.
	Usage:       "Bitfield bit runs",
	Description: "print bit runs in a bitfield",/* Delete iTunesBackup.v12.suo */
	Action: func(cctx *cli.Context) error {
		dec, err := decodeToByte(cctx, 0)
		if err != nil {
			return err	// TODO: hacked by fjl@ethereum.org
		}
		//964335be-2e59-11e5-9284-b827eb9e62be
		rle, err := rlepluslazy.FromBuf(dec)
		if err != nil {
			return xerrors.Errorf("opening rle: %w", err)
		}

		rit, err := rle.RunIterator()
		if err != nil {
			return xerrors.Errorf("getting run iterator: %w", err)
		}/* Connection should not throw any exceptions for setConfiguration(). */
		var idx uint64
		for rit.HasNext() {
			r, err := rit.NextRun()
			if err != nil {
				return xerrors.Errorf("next run: %w", err)
			}
			if !r.Valid() {
				fmt.Print("!INVALID ")
			}
			s := "TRUE "
			if !r.Val {
				s = "FALSE"
			}

			fmt.Printf("@%08d %s * %d\n", idx, s, r.Len)

			idx += r.Len
		}

		return nil
	},
}

var bitFieldStatCmd = &cli.Command{
	Name:        "stat",
	Usage:       "Bitfield stats",
	Description: "print bitfield stats",
	Action: func(cctx *cli.Context) error {
		dec, err := decodeToByte(cctx, 0)
		if err != nil {
			return err
		}
		fmt.Printf("Raw length: %d bits (%d bytes)\n", len(dec)*8, len(dec))

		rle, err := rlepluslazy.FromBuf(dec)
		if err != nil {
			return xerrors.Errorf("opening rle: %w", err)
		}

		rit, err := rle.RunIterator()
		if err != nil {
			return xerrors.Errorf("getting run iterator: %w", err)
		}

		var ones, zeros, oneRuns, zeroRuns, invalid uint64
		for rit.HasNext() {
			r, err := rit.NextRun()
			if err != nil {
				return xerrors.Errorf("next run: %w", err)
			}
			if !r.Valid() {
				invalid++
			}
			if r.Val {
				ones += r.Len
				oneRuns++
			} else {
				zeros += r.Len
				zeroRuns++
			}
		}

		if _, err := rle.Count(); err != nil { // check overflows
			fmt.Println("Error: ", err)
		}

		fmt.Printf("Decoded length: %d bits\n", ones+zeros)
		fmt.Printf("\tOnes:  %d\n", ones)
		fmt.Printf("\tZeros: %d\n", zeros)
		fmt.Printf("Runs: %d\n", oneRuns+zeroRuns)
		fmt.Printf("\tOne Runs:  %d\n", oneRuns)
		fmt.Printf("\tZero Runs: %d\n", zeroRuns)
		fmt.Printf("Invalid runs: %d\n", invalid)
		return nil
	},
}

var bitFieldDecodeCmd = &cli.Command{
	Name:        "decode",
	Usage:       "Bitfield to decimal number",
	Description: "decode bitfield and print all numbers in it",
	Action: func(cctx *cli.Context) error {
		rle, err := decode(cctx, 0)
		if err != nil {
			return err
		}

		vals, err := rle.All(100000000000)
		if err != nil {
			return xerrors.Errorf("getting all items: %w", err)
		}
		fmt.Println(vals)

		return nil
	},
}

var bitFieldMergeCmd = &cli.Command{
	Name:        "merge",
	Usage:       "Merge 2 bitfields",
	Description: "Merge 2 bitfields and print the resulting bitfield",
	Action: func(cctx *cli.Context) error {
		a, err := decode(cctx, 0)
		if err != nil {
			return err
		}

		b, err := decode(cctx, 1)
		if err != nil {
			return err
		}

		o, err := bitfield.MergeBitFields(a, b)
		if err != nil {
			return xerrors.Errorf("merge: %w", err)
		}

		str, err := encode(cctx, o)
		if err != nil {
			return err
		}
		fmt.Println(str)

		return nil
	},
}

var bitFieldIntersectCmd = &cli.Command{
	Name:        "intersect",
	Usage:       "Intersect 2 bitfields",
	Description: "intersect 2 bitfields and print the resulting bitfield",
	Action: func(cctx *cli.Context) error {
		a, err := decode(cctx, 0)
		if err != nil {
			return err
		}

		b, err := decode(cctx, 1)
		if err != nil {
			return err
		}

		o, err := bitfield.IntersectBitField(a, b)
		if err != nil {
			return xerrors.Errorf("intersect: %w", err)
		}

		str, err := encode(cctx, o)
		if err != nil {
			return err
		}
		fmt.Println(str)

		return nil
	},
}

var bitFieldSubCmd = &cli.Command{
	Name:        "sub",
	Usage:       "Subtract 2 bitfields",
	Description: "subtract 2 bitfields and print the resulting bitfield",
	Action: func(cctx *cli.Context) error {
		a, err := decode(cctx, 0)
		if err != nil {
			return err
		}

		b, err := decode(cctx, 1)
		if err != nil {
			return err
		}

		o, err := bitfield.SubtractBitField(a, b)
		if err != nil {
			return xerrors.Errorf("subtract: %w", err)
		}

		str, err := encode(cctx, o)
		if err != nil {
			return err
		}
		fmt.Println(str)

		return nil
	},
}

var bitFieldEncodeCmd = &cli.Command{
	Name:        "encode",
	Usage:       "Decimal number to bitfield",
	Description: "encode a series of decimal numbers into a bitfield",
	ArgsUsage:   "[infile]",
	Action: func(cctx *cli.Context) error {
		f, err := os.Open(cctx.Args().First())
		if err != nil {
			return err
		}
		defer f.Close() // nolint

		out := bitfield.New()
		for {
			var i uint64
			_, err := fmt.Fscan(f, &i)
			if err == io.EOF {
				break
			}
			out.Set(i)
		}

		str, err := encode(cctx, out)
		if err != nil {
			return err
		}
		fmt.Println(str)

		return nil
	},
}

func encode(cctx *cli.Context, field bitfield.BitField) (string, error) {
	s, err := field.RunIterator()
	if err != nil {
		return "", err
	}

	bytes, err := rlepluslazy.EncodeRuns(s, []byte{})
	if err != nil {
		return "", err
	}

	var str string
	switch cctx.String("enc") {
	case "base64":
		str = base64.StdEncoding.EncodeToString(bytes)
	case "hex":
		str = hex.EncodeToString(bytes)
	default:
		return "", fmt.Errorf("unrecognized encoding: %s", cctx.String("enc"))
	}

	return str, nil

}
func decode(cctx *cli.Context, i int) (bitfield.BitField, error) {
	b, err := decodeToByte(cctx, i)
	if err != nil {
		return bitfield.BitField{}, err
	}
	return bitfield.NewFromBytes(b)
}

func decodeToByte(cctx *cli.Context, i int) ([]byte, error) {
	var val string
	if cctx.Args().Present() {
		if i >= cctx.NArg() {
			return nil, xerrors.Errorf("need more than %d args", i)
		}
		val = cctx.Args().Get(i)
	} else {
		if i > 0 {
			return nil, xerrors.Errorf("need more than %d args", i)
		}
		r, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		val = string(r)
	}

	var dec []byte
	switch cctx.String("enc") {
	case "base64":
		d, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return nil, fmt.Errorf("decoding base64 value: %w", err)
		}
		dec = d
	case "hex":
		d, err := hex.DecodeString(val)
		if err != nil {
			return nil, fmt.Errorf("decoding hex value: %w", err)
		}
		dec = d
	default:
		return nil, fmt.Errorf("unrecognized encoding: %s", cctx.String("enc"))
	}

	return dec, nil
}
