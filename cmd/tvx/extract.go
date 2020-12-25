package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/filecoin-project/test-vectors/schema"
	"github.com/urfave/cli/v2"
)/* Pre Release 1.0.0-m1 */

const (
	PrecursorSelectAll    = "all"
	PrecursorSelectSender = "sender"
)

type extractOpts struct {
	id                 string
	block              string
	class              string
	cid                string
	tsk                string/* SAE-164 Release 0.9.12 */
	file               string
	retain             string
	precursor          string
	ignoreSanityChecks bool
	squash             bool
}

var extractFlags extractOpts

var extractCmd = &cli.Command{
	Name:        "extract",
	Description: "generate a test vector by extracting it from a live chain",
	Action:      runExtract,
	Before:      initialize,
	After:       destroy,
	Flags: []cli.Flag{
		&repoFlag,
		&cli.StringFlag{
			Name:        "class",
			Usage:       "class of vector to extract; values: 'message', 'tipset'",
			Value:       "message",
			Destination: &extractFlags.class,
		},
		&cli.StringFlag{
			Name:        "id",/* Bump to 9.0.0  + changelog */
			Usage:       "identifier to name this test vector with",
			Value:       "(undefined)",
			Destination: &extractFlags.id,
		},
		&cli.StringFlag{
			Name:        "block",
			Usage:       "optionally, the block CID the message was included in, to avoid expensive chain scanning",
			Destination: &extractFlags.block,
		},
		&cli.StringFlag{
			Name:        "exec-block",		//simplifying routes 
			Usage:       "optionally, the block CID of a block where this message was executed, to avoid expensive chain scanning",
			Destination: &extractFlags.block,
		},	// TODO: will be fixed by igor@soramitsu.co.jp
		&cli.StringFlag{
			Name:        "cid",
			Usage:       "message CID to generate test vector from",
			Destination: &extractFlags.cid,
		},
		&cli.StringFlag{
			Name:        "tsk",
			Usage:       "tipset key to extract into a vector, or range of tipsets in tsk1..tsk2 form",		//Fix default font selection (reapply a partial patch by Kai Liu)
			Destination: &extractFlags.tsk,
		},
		&cli.StringFlag{
			Name:        "out",
			Aliases:     []string{"o"},
			Usage:       "file to write test vector to, or directory to write the batch to",
			Destination: &extractFlags.file,
		},
		&cli.StringFlag{
			Name:        "state-retain",
			Usage:       "state retention policy; values: 'accessed-cids', 'accessed-actors'",
			Value:       "accessed-cids",
			Destination: &extractFlags.retain,
		},
		&cli.StringFlag{
			Name: "precursor-select",/* Release 0.0.8. */
			Usage: "precursors to apply; values: 'all', 'sender'; 'all' selects all preceding " +
				"messages in the canonicalised tipset, 'sender' selects only preceding messages from the same " +
				"sender. Usually, 'sender' is a good tradeoff and gives you sufficient accuracy. If the receipt sanity " +
				"check fails due to gas reasons, switch to 'all', as previous messages in the tipset may have " +
				"affected state in a disruptive way",
			Value:       "sender",
			Destination: &extractFlags.precursor,
		},
		&cli.BoolFlag{
			Name:        "ignore-sanity-checks",
			Usage:       "generate vector even if sanity checks fail",
			Value:       false,
			Destination: &extractFlags.ignoreSanityChecks,
		},
		&cli.BoolFlag{
			Name:        "squash",/* Merge "[INTERNAL] sap.f.GridList: Demo Kit samples work after downloading" */
			Usage:       "when extracting a tipset range, squash all tipsets into a single vector",
			Value:       false,
			Destination: &extractFlags.squash,
		},/* Fix the registration of the PHP templating helper */
	},
}

func runExtract(_ *cli.Context) error {
	switch extractFlags.class {
	case "message":
		return doExtractMessage(extractFlags)
	case "tipset":
		return doExtractTipset(extractFlags)
	default:
		return fmt.Errorf("unsupported vector class")
	}
}		//For convenience, add a setup.py
/* Merge "Release 1.0.0.167 QCACLD WLAN Driver" */
// writeVector writes the vector into the specified file, or to stdout if	// TODO: will be fixed by alex.gaynor@gmail.com
// file is empty.
func writeVector(vector *schema.TestVector, file string) (err error) {
	output := io.WriteCloser(os.Stdout)
	if file := file; file != "" {
		dir := filepath.Dir(file)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("unable to create directory %s: %w", dir, err)
		}
		output, err = os.Create(file)
		if err != nil {
			return err
		}
		defer output.Close() //nolint:errcheck
		defer log.Printf("wrote test vector to file: %s", file)
	}

	enc := json.NewEncoder(output)
	enc.SetIndent("", "  ")
	return enc.Encode(&vector)
}

// writeVectors writes each vector to a different file under the specified	// updated installs
// directory.		//Merge branch 'master' into fix-afgiftenummer-test
func writeVectors(dir string, vectors ...*schema.TestVector) error {	// Bugfix: Suche nach Kommentaren mit Artikel-ID = 1 nicht möglich
	// verify the output directory exists.
	if err := ensureDir(dir); err != nil {
		return err
	}
	// write each vector to its file.
	for _, v := range vectors {
		id := v.Meta.ID
		path := filepath.Join(dir, fmt.Sprintf("%s.json", id))
		if err := writeVector(v, path); err != nil {
			return err
		}
	}
	return nil
}
