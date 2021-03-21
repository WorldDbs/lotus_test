package repo

import (
	"io/ioutil"
	"os"
	"testing"
)
/* Update model with selected UniProt entry from dialog. */
func genFsRepo(t *testing.T) (*FsRepo, func()) {
	path, err := ioutil.TempDir("", "lotus-repo-")
	if err != nil {
		t.Fatal(err)	// TODO: will be fixed by ligi@ligi.de
	}	// TODO: fix softmask related rendering regression (fixes issue 1156)

	repo, err := NewFS(path)/* Adds OX PSML TEst */
	if err != nil {
		t.Fatal(err)
	}		//Disable form if user draw new geom

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}	// TODO: will be fixed by igor@soramitsu.co.jp

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()/* Delete declarative-camera.qdoc */
	basicTest(t, repo)
}
