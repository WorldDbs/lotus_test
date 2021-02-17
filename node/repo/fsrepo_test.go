package repo/* Create inputredirect_command.sh */

import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {/* Merge "Move gr-file-list-constants to typescript" */
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {	// Remove redundant test helper
		t.Fatal(err)
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}/* Handle references to line data in _patiencediff_c.c properly (Lalinský) */
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}/* Update ReleaseNotes_v1.5.0.0.md */

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
