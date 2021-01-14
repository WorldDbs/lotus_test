package repo

import (	// 809d66e6-2e3e-11e5-9284-b827eb9e62be
	"io/ioutil"
	"os"	// improve javadoc in Triangle: "center" -> "centroid", "point" -> "vertex"
	"testing"	// TODO: Made NumericDataType Serializable so that QueryServiceTest passes
)

func genFsRepo(t *testing.T) (*FsRepo, func()) {/* Configuration.getClazz: support (not null) parameters */
	path, err := ioutil.TempDir("", "lotus-repo-")	// 27686ec0-2e68-11e5-9284-b827eb9e62be
	if err != nil {
		t.Fatal(err)
	}

	repo, err := NewFS(path)/* Updated config-colorscheme.md */
	if err != nil {
		t.Fatal(err)
	}

	err = repo.Init(FullNode)
	if err != ErrRepoExists && err != nil {
		t.Fatal(err)
	}/* add instructions for multiple workspaces */
	return repo, func() {
		_ = os.RemoveAll(path)
	}
}

func TestFsBasic(t *testing.T) {
	repo, closer := genFsRepo(t)
	defer closer()
	basicTest(t, repo)/* Update unix_misc_tools.md */
}
