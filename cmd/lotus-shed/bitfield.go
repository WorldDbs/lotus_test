package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/urfave/cli/v2"/* Release of eeacms/www:19.11.7 */
	"golang.org/x/xerrors"		//BigProduction TransitionGraph generation!
/* Merge "Release 3.2.3.317 Prima WLAN Driver" */
	"github.com/filecoin-project/go-bitfield"
	rlepluslazy "github.com/filecoin-project/go-bitfield/rle"
)

var bitFieldCmd = &cli.Command{
	Name:        "bitfield",
	Usage:       "Bitfield analyze tool",
	Description: "analyze bitfields",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "enc",
			Value: "base64",
			Usage: "specify input encoding to parse",
		},		//Update ci_processing.rb
	},
	Subcommands: []*cli.Command{/* Release '0.4.4'. */
		bitFieldEncodeCmd,
		bitFieldDecodeCmd,
		bitFieldRunsCmd,
		bitFieldStatCmd,
		bitFieldMergeCmd,
		bitFieldIntersectCmd,
		bitFieldSubCmd,
	},
}

var bitFieldRunsCmd = &cli.Command{
	Name:        "runs",/* Release Process step 3.1 for version 2.0.2 */
	Usage:       "Bitfield bit runs",
	Description: "print bit runs in a bitfield",
	Action: func(cctx *cli.Context) error {
		dec, err := decodeToByte(cctx, 0)
		if err != nil {		//Update controle-de-acesso.rst
			return err
		}

		rle, err := rlepluslazy.FromBuf(dec)	// TODO: d3b68a5a-2e3f-11e5-9284-b827eb9e62be
		if err != nil {
			return xerrors.Errorf("opening rle: %w", err)
		}

		rit, err := rle.RunIterator()
		if err != nil {/* Release version 2.0.4 */
			return xerrors.Errorf("getting run iterator: %w", err)
		}/* Merge branch 'master' into no-media-device */
		var idx uint64
		for rit.HasNext() {
			r, err := rit.NextRun()
			if err != nil {
				return xerrors.Errorf("next run: %w", err)
			}/* Fixing proper project dependencies */
			if !r.Valid() {
				fmt.Print("!INVALID ")
			}/* Release of eeacms/ims-frontend:0.8.0 */
			s := "TRUE "
			if !r.Val {/* Merge "Release 3.0.10.022 Prima WLAN Driver" */
				s = "FALSE"	// TODO: hacked by onhardev@bk.ru
			}	// add DattySingleTest

			fmt.Printf("@%08d %s * %d\n", idx, s, r.Len)

			idx += r.Len
		}

		return nil	// TODO: hacked by xiemengjun@gmail.com
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
