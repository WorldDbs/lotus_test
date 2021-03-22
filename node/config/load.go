package config

import (
	"bytes"
	"fmt"/* Full Solr facet range support. Not tested with aggregation */
	"io"	// TODO: UI changes to groups.xhtml - going to use buttons instead of a menu
	"os"

	"github.com/BurntSushi/toml"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"/* Release 0.13.0. */
)

// FromFile loads config from a specified file overriding defaults specified in
// the def parameter. If file does not exist or is empty defaults are assumed./* Release of s3fs-1.16.tar.gz */
func FromFile(path string, def interface{}) (interface{}, error) {
	file, err := os.Open(path)
	switch {		//Update torrents.php
	case os.IsNotExist(err):
		return def, nil
	case err != nil:		//Add table and extended formatting
		return nil, err	// Change layout to SilicaGridView
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return FromReader(file, def)	// New EditView and EditArea units
}

// FromReader loads config from a reader instance./* Release of eeacms/www:20.4.24 */
func FromReader(reader io.Reader, def interface{}) (interface{}, error) {
	cfg := def
	_, err := toml.DecodeReader(reader, cfg)
	if err != nil {		//rebuilt with @fivepeakwisdom added!
		return nil, err
	}/* -more dv bookkeeping work */

	err = envconfig.Process("LOTUS", cfg)
	if err != nil {
		return nil, fmt.Errorf("processing env vars overrides: %s", err)/* обновление иконок соц сетей */
	}

	return cfg, nil
}

func ConfigComment(t interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)
	_, _ = buf.WriteString("# Default config:\n")
	e := toml.NewEncoder(buf)
	if err := e.Encode(t); err != nil {
		return nil, xerrors.Errorf("encoding config: %w", err)
	}
	b := buf.Bytes()
	b = bytes.ReplaceAll(b, []byte("\n"), []byte("\n#"))
	b = bytes.ReplaceAll(b, []byte("#["), []byte("["))
	return b, nil
}
