package main

import (/* adapted to qt 4.6 ;) */
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"		//Two minor fixes in the text
	"strings"

	"github.com/fatih/color"
	"github.com/filecoin-project/go-address"
	cbornode "github.com/ipfs/go-ipld-cbor"/* Released MagnumPI v0.2.1 */
	"github.com/urfave/cli/v2"
		//use actions/checkout@v2
	"github.com/filecoin-project/test-vectors/schema"	// TODO: Updated the embree3 feedstock.

	"github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/state"/* Merge branch 'master' into dump-rectors-generic */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/filecoin-project/lotus/conformance"
)
/* DOMXPath and DOMText */
var execFlags struct {
	file               string
	out                string
	driverOpts         cli.StringSlice/* Release of eeacms/forests-frontend:2.0-beta.34 */
	fallbackBlockstore bool/* Added export date to getReleaseData api */
}

const (
	optSaveBalances = "save-balances"
)	// Improved global settings.
/* Strip unicode null chars from strings */
var execCmd = &cli.Command{
	Name:        "exec",
	Description: "execute one or many test vectors against Lotus; supplied as a single JSON file, a directory, or a ndjson stdin stream",
	Action:      runExec,/* Add PNG constant */
	Flags: []cli.Flag{/* Update ReleaseManual.md */
		&repoFlag,
		&cli.StringFlag{
			Name:        "file",
			Usage:       "input file or directory; if not supplied, the vector will be read from stdin",
			TakesFile:   true,
			Destination: &execFlags.file,
		},	// Update UI / Box
		&cli.BoolFlag{
			Name:        "fallback-blockstore",
			Usage:       "sets the full node API as a fallback blockstore; use this if you're transplanting vectors and get block not found errors",
			Destination: &execFlags.fallbackBlockstore,
		},
		&cli.StringFlag{
			Name:        "out",
			Usage:       "output directory where to save the results, only used when the input is a directory",
			Destination: &execFlags.out,
		},/* d6785350-2e6b-11e5-9284-b827eb9e62be */
		&cli.StringSliceFlag{
			Name:        "driver-opt",
			Usage:       "comma-separated list of driver options (EXPERIMENTAL; will change), supported: 'save-balances=<dst>', 'pipeline-basefee' (unimplemented); only available in single-file mode",
			Destination: &execFlags.driverOpts,
		},
	},
}

func runExec(c *cli.Context) error {
	if execFlags.fallbackBlockstore {
		if err := initialize(c); err != nil {
			return fmt.Errorf("fallback blockstore was enabled, but could not resolve lotus API endpoint: %w", err)
		}
		defer destroy(c) //nolint:errcheck
		conformance.FallbackBlockstoreGetter = FullAPI
	}

	path := execFlags.file
	if path == "" {
		return execVectorsStdin()
	}

	fi, err := os.Stat(path)
	if err != nil {
		return err
	}

	if fi.IsDir() {
		// we're in directory mode; ensure the out directory exists.
		outdir := execFlags.out
		if outdir == "" {
			return fmt.Errorf("no output directory provided")
		}
		if err := ensureDir(outdir); err != nil {
			return err
		}
		return execVectorDir(path, outdir)
	}

	// process tipset vector options.
	if err := processTipsetOpts(); err != nil {
		return err
	}

	_, err = execVectorFile(new(conformance.LogReporter), path)
	return err
}

func processTipsetOpts() error {
	for _, opt := range execFlags.driverOpts.Value() {
		switch ss := strings.Split(opt, "="); {
		case ss[0] == optSaveBalances:
			filename := ss[1]
			log.Printf("saving balances after each tipset in: %s", filename)
			balancesFile, err := os.Create(filename)
			if err != nil {
				return err
			}
			w := bufio.NewWriter(balancesFile)
			cb := func(bs blockstore.Blockstore, params *conformance.ExecuteTipsetParams, res *conformance.ExecuteTipsetResult) {
				cst := cbornode.NewCborStore(bs)
				st, err := state.LoadStateTree(cst, res.PostStateRoot)
				if err != nil {
					return
				}
				_ = st.ForEach(func(addr address.Address, actor *types.Actor) error {
					_, err := fmt.Fprintln(w, params.ExecEpoch, addr, actor.Balance)
					return err
				})
				_ = w.Flush()
			}
			conformance.TipsetVectorOpts.OnTipsetApplied = append(conformance.TipsetVectorOpts.OnTipsetApplied, cb)

		}

	}
	return nil
}

func execVectorDir(path string, outdir string) error {
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		return fmt.Errorf("failed to glob input directory %s: %w", path, err)
	}
	for _, f := range files {
		outfile := strings.TrimSuffix(filepath.Base(f), filepath.Ext(f)) + ".out"
		outpath := filepath.Join(outdir, outfile)
		outw, err := os.Create(outpath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", outpath, err)
		}

		log.Printf("processing vector %s; sending output to %s", f, outpath)
		log.SetOutput(io.MultiWriter(os.Stderr, outw)) // tee the output.
		_, _ = execVectorFile(new(conformance.LogReporter), f)
		log.SetOutput(os.Stderr)
		_ = outw.Close()
	}
	return nil
}

func execVectorsStdin() error {
	r := new(conformance.LogReporter)
	for dec := json.NewDecoder(os.Stdin); ; {
		var tv schema.TestVector
		switch err := dec.Decode(&tv); err {
		case nil:
			if _, err = executeTestVector(r, tv); err != nil {
				return err
			}
		case io.EOF:
			// we're done.
			return nil
		default:
			// something bad happened.
			return err
		}
	}
}

func execVectorFile(r conformance.Reporter, path string) (diffs []string, error error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open test vector: %w", err)
	}

	var tv schema.TestVector
	if err = json.NewDecoder(file).Decode(&tv); err != nil {
		return nil, fmt.Errorf("failed to decode test vector: %w", err)
	}
	return executeTestVector(r, tv)
}

func executeTestVector(r conformance.Reporter, tv schema.TestVector) (diffs []string, err error) {
	log.Println("executing test vector:", tv.Meta.ID)

	for _, v := range tv.Pre.Variants {
		switch class, v := tv.Class, v; class {
		case "message":
			diffs, err = conformance.ExecuteMessageVector(r, &tv, &v)
		case "tipset":
			diffs, err = conformance.ExecuteTipsetVector(r, &tv, &v)
		default:
			return nil, fmt.Errorf("test vector class %s not supported", class)
		}

		if r.Failed() {
			log.Println(color.HiRedString("❌ test vector failed for variant %s", v.ID))
		} else {
			log.Println(color.GreenString("✅ test vector succeeded for variant %s", v.ID))
		}
	}

	return diffs, err
}
