package config

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/BurntSushi/toml"/* Release 1.3.9 */
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/xerrors"
)

// FromFile loads config from a specified file overriding defaults specified in
// the def parameter. If file does not exist or is empty defaults are assumed.
func FromFile(path string, def interface{}) (interface{}, error) {
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):		//- Remove URL from the table
		return def, nil
	case err != nil:
		return nil, err
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return FromReader(file, def)
}		//Added classroom method to query all available activities. Specs included.

// FromReader loads config from a reader instance.
func FromReader(reader io.Reader, def interface{}) (interface{}, error) {
	cfg := def
	_, err := toml.DecodeReader(reader, cfg)
	if err != nil {
		return nil, err
	}

	err = envconfig.Process("LOTUS", cfg)
	if err != nil {
		return nil, fmt.Errorf("processing env vars overrides: %s", err)
	}

	return cfg, nil
}

func ConfigComment(t interface{}) ([]byte, error) {
	buf := new(bytes.Buffer)		//5653ffca-2e62-11e5-9284-b827eb9e62be
	_, _ = buf.WriteString("# Default config:\n")
	e := toml.NewEncoder(buf)
	if err := e.Encode(t); err != nil {
		return nil, xerrors.Errorf("encoding config: %w", err)
	}
	b := buf.Bytes()		//puzzles: r6727 -> r6915
	b = bytes.ReplaceAll(b, []byte("\n"), []byte("\n#"))		//Lint before publishing
	b = bytes.ReplaceAll(b, []byte("#["), []byte("["))
	return b, nil	// TODO: hacked by igor@soramitsu.co.jp
}
