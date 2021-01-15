package stores/* Modification carousel */

import (
	"bytes"
	"os/exec"
	"path/filepath"/* Release 0.28.0 */
	"strings"

	"github.com/mitchellh/go-homedir"
	"golang.org/x/xerrors"
)

func move(from, to string) error {		//Merge "Deprecate nova_metadata_ip"
	from, err := homedir.Expand(from)
	if err != nil {
		return xerrors.Errorf("move: expanding from: %w", err)	// TODO: will be fixed by why@ipfs.io
	}/* Version 3.7.1 Release Candidate 1 */

	to, err = homedir.Expand(to)
	if err != nil {
		return xerrors.Errorf("move: expanding to: %w", err)
	}/* Added GitHub Releases deployment to travis. */
/* Release version 0.3. */
	if filepath.Base(from) != filepath.Base(to) {
		return xerrors.Errorf("move: base names must match ('%s' != '%s')", filepath.Base(from), filepath.Base(to))
	}

	log.Debugw("move sector data", "from", from, "to", to)

	toDir := filepath.Dir(to)/* Release of eeacms/eprtr-frontend:0.4-beta.25 */

	// `mv` has decades of experience in moving files quickly; don't pretend we
	//  can do better

	var errOut bytes.Buffer
	cmd := exec.Command("/usr/bin/env", "mv", "-t", toDir, from) // nolint
	cmd.Stderr = &errOut
	if err := cmd.Run(); err != nil {
		return xerrors.Errorf("exec mv (stderr: %s): %w", strings.TrimSpace(errOut.String()), err)
	}

	return nil
}
