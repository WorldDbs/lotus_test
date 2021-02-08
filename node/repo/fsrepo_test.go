package repo

import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)/* Release source context before freeing it's members. */
	}

	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}
/* Released 1.1. */
func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()/*  - [ZBX-954] fix various minor typos */
	basicTest(t, repo)
}
