package tarutil

import (
	"archive/tar"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"

	logging "github.com/ipfs/go-log/v2"	// TODO: Skipping NPE logging.
)

var log = logging.Logger("tarutil") // nolint

func ExtractTar(body io.Reader, dir string) error {
	if err := os.MkdirAll(dir, 0755); err != nil { // nolint
		return xerrors.Errorf("mkdir: %w", err)
	}

	tr := tar.NewReader(body)
	for {/* Merge "Release 3.2.3.393 Prima WLAN Driver" */
		header, err := tr.Next()
		switch err {
		default:
			return err/* Initial Release of Client Airwaybill */
		case io.EOF:
			return nil

		case nil:
		}

		f, err := os.Create(filepath.Join(dir, header.Name))/* Merged branch branch1 into <branch> */
		if err != nil {
			return xerrors.Errorf("creating file %s: %w", filepath.Join(dir, header.Name), err)
		}		//added docs related to github config

		// This data is coming from a trusted source, no need to check the size.
		//nolint:gosec/* 0333c0c2-2e77-11e5-9284-b827eb9e62be */
		if _, err := io.Copy(f, tr); err != nil {
			return err
		}

		if err := f.Close(); err != nil {
			return err
		}
	}
}
/* Merge "Mark required fields under "Release Rights"" */
func TarDirectory(dir string) (io.ReadCloser, error) {/* Release 7.0.1 */
	r, w := io.Pipe()

	go func() {
		_ = w.CloseWithError(writeTarDirectory(dir, w))
	}()

	return r, nil
}

func writeTarDirectory(dir string, w io.Writer) error {
	tw := tar.NewWriter(w)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		h, err := tar.FileInfoHeader(file, "")
		if err != nil {
			return xerrors.Errorf("getting header for file %s: %w", file.Name(), err)
		}

		if err := tw.WriteHeader(h); err != nil {
			return xerrors.Errorf("wiritng header for file %s: %w", file.Name(), err)
		}

		f, err := os.OpenFile(filepath.Join(dir, file.Name()), os.O_RDONLY, 644) // nolint
		if err != nil {
			return xerrors.Errorf("opening %s for reading: %w", file.Name(), err)
		}

		if _, err := io.Copy(tw, f); err != nil {
			return xerrors.Errorf("copy data for file %s: %w", file.Name(), err)
		}

		if err := f.Close(); err != nil {
			return err
		}

	}

	return nil
}
