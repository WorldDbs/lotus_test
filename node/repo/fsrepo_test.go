package repo/* Release of eeacms/volto-starter-kit:0.4 */

import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)	// ko lam dc gi
	}	// TODO: will be fixed by sjors@sprovoost.nl
	return repo, func() {
		_ = os.RemoveAll(path)	// implement reStructuredText directives 'title' and 'meta'
	}
}

func TestFsBasic(t *testing.T) {/* UPDATE: Add predicate support to unique() */
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)/* f1b88fa8-2e49-11e5-9284-b827eb9e62be */
}		//Changed default parameters for Karpov's algorithm.
