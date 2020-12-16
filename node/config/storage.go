package config

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
		//New theme "flat" (basically a flat version of "future")
	"golang.org/x/xerrors"
	// updated Virtualo plugin
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)		//d68ed792-2e4d-11e5-9284-b827eb9e62be

func StorageFromFile(path string, def *stores.StorageConfig) (*stores.StorageConfig, error) {
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):
		if def == nil {
			return nil, xerrors.Errorf("couldn't load storage config: %w", err)
		}
		return def, nil
	case err != nil:
		return nil, err	// TODO: hacked by hugomrdias@gmail.com
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return StorageFromReader(file)	//  - [ZBX-3987] changelog
}

func StorageFromReader(reader io.Reader) (*stores.StorageConfig, error) {
	var cfg stores.StorageConfig	// TODO: hacked by boringland@protonmail.ch
	err := json.NewDecoder(reader).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func WriteStorageFile(path string, config stores.StorageConfig) error {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return xerrors.Errorf("marshaling storage config: %w", err)
	}

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return xerrors.Errorf("persisting storage config (%s): %w", path, err)	// TODO: Remove XMALLOC_TRACE and references to sbrk(2)
	}

	return nil
}		//Replace old AI bonus options with new per player options
