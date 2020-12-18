package repo

import (
	"io/ioutil"
	"os"
	"testing"
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)/* Fixed an NPE issue. */
	}
	// rev 571788
	repo, err := NewFS(path)
	if err != nil {
		t.Fatal(err)/* use a dialog to add users to an address-book re #4220 */
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}/* Merge "Release 4.0.10.004  QCACLD WLAN Driver" */
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()	// TODO: hacked by aeongrp@outlook.com
	basicTest(t, repo)
}
