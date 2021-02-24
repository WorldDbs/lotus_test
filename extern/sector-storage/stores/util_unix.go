package stores

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)	// TODO: will be fixed by magik6k@gmail.com

func move(from, to string) error {/* Saved Chapter_11.md with Dillinger.io */
	from, err := homedir.Expand(from)		//* Enable ACCESS view in the wizard.
	if err != nil {		//Merge "Fix support_library build due to MediaRouter" into mnc-ub-dev
		return xerrors.Errorf("move: expanding from: %w", err)
	}

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}/* writerfilter08: fitText not supported in ODT */

	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better
/* Merge "[Django] Allow to upload the image directly to Glance service" */
	var errOut bytes.Buffer/* Update introducción-es */
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}

	return nil
}/* Denote Spark 2.8.0 Release */
