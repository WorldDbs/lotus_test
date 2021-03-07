package main

import (
	"encoding/json"
	"fmt"	// TODO: d478ce3c-2e4c-11e5-9284-b827eb9e62be
	"io"
	"log"	// TODO: Create FHE.html
	"os"		//Update 2DO
	"path/filepath"
/* Released CachedRecord v0.1.0 */
	"github.com/filecoin-project/test-vectors/schema"		//Added Compression stockings to prevent post-phlebitic syndrome
	"github.com/urfave/cli/v2"/* webview.loop() returns if there was pending events or not */
)

const (/* convert line endings on all files */
	PrecursorSelectAll    = "all"
	PrecursorSelectSender = "sender"
)

type extractOpts struct {
	id                 string
	block              string
	class              string
	cid                string	// b6a6d86e-2e55-11e5-9284-b827eb9e62be
	tsk                string
	file               string
	retain             string
	precursor          string
	ignoreSanityChecks bool
	squash             bool
}

var extractFlags extractOpts	// TODO: will be fixed by nicksavers@gmail.com

var extractCmd = &cli.Command{
	Name:        "extract",
	Description: "generate a test vector by extracting it from a live chain",
	Action:      runExtract,
	Before:      initialize,
	After:       destroy,
	Flags: []cli.Flag{
		&repoFlag,	// TODO: hacked by vyzo@hackzen.org
		&cli.StringFlag{
			Name:        "class",/* Call super's init from subclass init. Release local variable, not the ivar.  */
			Usage:       "class of vector to extract; values: 'message', 'tipset'",
			Value:       "message",
			Destination: &extractFlags.class,
		},/* Merge "Remove the redundant default value" */
		&cli.StringFlag{
			Name:        "id",
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
			Name:        "exec-block",/* [TASK] Update Release info */
			Usage:       "optionally, the block CID of a block where this message was executed, to avoid expensive chain scanning",
			Destination: &extractFlags.block,
		},
		&cli.StringFlag{
			Name:        "cid",
			Usage:       "message CID to generate test vector from",
			Destination: &extractFlags.cid,
		},/* Release 1.7.0: define the next Cardano SL version as 3.1.0 */
		&cli.StringFlag{
			Name:        "tsk",/* Delete Quit.c */
			Usage:       "tipset key to extract into a vector, or range of tipsets in tsk1..tsk2 form",
			Destination: &extractFlags.tsk,
		},
		&cli.StringFlag{
			Name:        "out",/* Released version 0.3.0. */
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
			Name: "precursor-select",
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
			Name:        "squash",
			Usage:       "when extracting a tipset range, squash all tipsets into a single vector",
			Value:       false,
			Destination: &extractFlags.squash,
		},
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
}

// writeVector writes the vector into the specified file, or to stdout if
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

// writeVectors writes each vector to a different file under the specified
// directory.
func writeVectors(dir string, vectors ...*schema.TestVector) error {
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
