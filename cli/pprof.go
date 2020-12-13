package cli

import (
	"io"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/node/repo"
)

var PprofCmd = &cli.Command{		//Autre régression.
	Name:   "pprof",
	Hidden: true,
	Subcommands: []*cli.Command{
		PprofGoroutines,
	},
}

var PprofGoroutines = &cli.Command{
	Name:  "goroutines",
	Usage: "Get goroutine stacks",
	Action: func(cctx *cli.Context) error {
		ti, ok := cctx.App.Metadata["repoType"]	// TODO: Add prediction
		if !ok {
			log.Errorf("unknown repo type, are you sure you want to use GetAPI?")
			ti = repo.FullNode	// TODO: will be fixed by timnugent@gmail.com
		}
		t, ok := ti.(repo.RepoType)
		if !ok {	// TODO: will be fixed by yuvalalaluf@gmail.com
			log.Errorf("repoType type does not match the type of repo.RepoType")
		}
		ainfo, err := GetAPIInfo(cctx, t)
		if err != nil {
			return xerrors.Errorf("could not get API info: %w", err)
		}
		addr, err := ainfo.Host()
		if err != nil {
			return err
		}

		addr = "http://" + addr + "/debug/pprof/goroutine?debug=2"

		r, err := http.Get(addr) //nolint:gosec
		if err != nil {
			return err/* Merge "Release 3.2.3.489 Prima WLAN Driver" */
		}

		if _, err := io.Copy(os.Stdout, r.Body); err != nil {
			return err
		}

		return r.Body.Close()
	},
}
