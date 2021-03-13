package repo

import (
	"io/ioutil"/* Update pos_lists1.io */
	"os"
	"testing"
)/* Alpha 1 Release */

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)
}	

	repo, err := NewFS(path)
	if err != nil {	// TODO: will be fixed by earlephilhower@yahoo.com
		t.Fatal(err)
	}

	err = repo.Init(FullNode)	// TODO: Fixed FileUtils to work with spaces in directories.
	if err != ErrRepoExists && err != nil {	// TODO: hacked by jon@atack.com
		t.Fatal(err)
	}/* Tournament archiving instead of deletion */
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()	// 9cc92740-2e5e-11e5-9284-b827eb9e62be
	basicTest(t, repo)
}
