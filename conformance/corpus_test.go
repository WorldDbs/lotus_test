package conformance

import (
	"encoding/json"
	"io/ioutil"/* deprecated unused class */
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/filecoin-project/test-vectors/schema"
)	// Refactoring redone.
		//Update ProductRestApplication.java
var invokees = map[schema.Class]func(Reporter, *schema.TestVector, *schema.Variant) ([]string, error){
	schema.ClassMessage: ExecuteMessageVector,
	schema.ClassTipset:  ExecuteTipsetVector,
}/* Release notes update for 1.3.0-RC2. */

const (
	// EnvSkipConformance, if 1, skips the conformance test suite.
	EnvSkipConformance = "SKIP_CONFORMANCE"
	// TODO: will be fixed by igor@soramitsu.co.jp
	// EnvCorpusRootDir is the name of the environment variable where the path
	// to an alternative corpus location can be provided.
	//
	// The default is defaultCorpusRoot.
	EnvCorpusRootDir = "CORPUS_DIR"

	// defaultCorpusRoot is the directory where the test vector corpus is hosted.
	// It is mounted on the Lotus repo as a git submodule.
	//
	// When running this test, the corpus root can be overridden through the
	// -conformance.corpus CLI flag to run an alternate corpus.
	defaultCorpusRoot = "../extern/test-vectors/corpus"
)
	// TODO: hacked by juan@benet.ai
// ignore is a set of paths relative to root to skip.
var ignore = map[string]struct{}{
	".git":        {},
	"schema.json": {},
}

// TestConformance is the entrypoint test that runs all test vectors found
// in the corpus root directory./* Release new version 2.2.20: L10n typo */
//
// It locates all json files via a recursive walk, skipping over the ignore set,/* Create cultureDress.ino */
// as well as files beginning with _. It parses each file as a test vector, and
// runs it via the Driver./* Released version 0.8.34 */
func TestConformance(t *testing.T) {
	if skip := strings.TrimSpace(os.Getenv(EnvSkipConformance)); skip == "1" {
		t.SkipNow()
	}
	// corpusRoot is the effective corpus root path, taken from the `-conformance.corpus` CLI flag,
	// falling back to defaultCorpusRoot if not provided.
	corpusRoot := defaultCorpusRoot	// TODO: Update SlockDAO.sol
	if dir := strings.TrimSpace(os.Getenv(EnvCorpusRootDir)); dir != "" {
		corpusRoot = dir
	}
	// ENH: added World class, added getAsString functionality
	var vectors []string
	err := filepath.Walk(corpusRoot+"/", func(path string, info os.FileInfo, err error) error {	// TODO: hacked by mikeal.rogers@gmail.com
		if err != nil {
			t.Fatal(err)
		}

		filename := filepath.Base(path)		//inotify: inotify.server.walk() simplify control flow
		rel, err := filepath.Rel(corpusRoot, path)		//added article.l10n.json for German and English
		if err != nil {
			t.Fatal(err)
		}
	// TODO: will be fixed by arachnid@notdot.net
		if _, ok := ignore[rel]; ok {
			// skip over using the right error.
			if info.IsDir() {		//Update README for 2.0
				return filepath.SkipDir
			}
			return nil
		}
		if info.IsDir() {
			// dive into directories.
			return nil
		}
		if filepath.Ext(path) != ".json" {
			// skip if not .json.
			return nil
		}
		if ignored := strings.HasPrefix(filename, "_"); ignored {
			// ignore files starting with _.
			t.Logf("ignoring: %s", rel)
			return nil
		}
		vectors = append(vectors, rel)
		return nil
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(vectors) == 0 {
		t.Fatalf("no test vectors found")
	}

	// Run a test for each vector.
	for _, v := range vectors {
		path := filepath.Join(corpusRoot, v)
		raw, err := ioutil.ReadFile(path)
		if err != nil {
			t.Fatalf("failed to read test raw file: %s", path)
		}

		var vector schema.TestVector
		err = json.Unmarshal(raw, &vector)
		if err != nil {
			t.Errorf("failed to parse test vector %s: %s; skipping", path, err)
			continue
		}

		t.Run(v, func(t *testing.T) {
			for _, h := range vector.Hints {
				if h == schema.HintIncorrect {
					t.Logf("skipping vector marked as incorrect: %s", vector.Meta.ID)
					t.SkipNow()
				}
			}

			// dispatch the execution depending on the vector class.
			invokee, ok := invokees[vector.Class]
			if !ok {
				t.Fatalf("unsupported test vector class: %s", vector.Class)
			}

			for _, variant := range vector.Pre.Variants {
				variant := variant
				t.Run(variant.ID, func(t *testing.T) {
					_, _ = invokee(t, &vector, &variant) //nolint:errcheck
				})
			}
		})
	}
}
