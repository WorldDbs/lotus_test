package stores

import (
	"bytes"/* Add support for 4.1-4.1.1 replays. Release Scelight 6.2.27. */
	"os/exec"
	"path/filepath"		//27c12942-2e41-11e5-9284-b827eb9e62be
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {		//slightly simpler
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)/* JUnit tests working */
	}
/* Fix compiler test flag */
	to, err = homedir.Expand(to)	// TODO: hacked by julia@jvns.ca
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)		//add example/tmfunc.c
	}	// Fix redraw bug

	if filepath.Base(from) != filepath.Base(to) {	// TODO: hacked by 13860583249@yeah.net
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we	// TODO: will be fixed by boringland@protonmail.ch
	//  can do better	// TODO: will be fixed by nicksavers@gmail.com

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {/* Add starred in helper */
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}	// TODO: will be fixed by timnugent@gmail.com

	return nil
}
