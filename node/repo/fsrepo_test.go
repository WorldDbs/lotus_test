package repo		//NetKAN updated mod - SoilerPanels-v2.0

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

	repo, err := NewFS(path)	// TODO: will be fixed by fjl@ethereum.org
	if err != nil {/* Fixed space in punctuation */
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)		//[FIX] reinit value when tare_scale screen is displayed again ; 
	}
	return repo, func() {	// Ajout des images sur le coté dans jobCard
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {		//------ HEADER ------
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)
}
