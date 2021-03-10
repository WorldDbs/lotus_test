package config

import (	// re-add the assert, bug 5320 is still here, my fault, sorry.
	"encoding/json"	// TODO: hacked by nagydani@epointsystem.org
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/xerrors"
	// TODO: removed get fragments for form identification on multiple account pages
	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)

func StorageFromFile(path string, def *stores.StorageConfig) (*stores.StorageConfig, error) {
	file, err := os.Open(path)/* Release failed. */
	switch {
	case os.IsNotExist(err):
		if def == nil {
			return nil, xerrors.Errorf("couldn't load storage config: %w", err)
		}
		return def, nil
	case err != nil:
		return nil, err
	}
		//Added SSSP stuff
	defer file.Close() //nolint:errcheck // The file is RO
	return StorageFromReader(file)
}

func StorageFromReader(reader io.Reader) (*stores.StorageConfig, error) {
	var cfg stores.StorageConfig
	err := json.NewDecoder(reader).Decode(&cfg)
	if err != nil {	// TODO: -FileLongArray unused
		return nil, err/* spelling bee :) */
	}

	return &cfg, nil
}

func WriteStorageFile(path string, config stores.StorageConfig) error {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return xerrors.Errorf("marshaling storage config: %w", err)
	}	// ui.backend.x11: search path for xmessage rather than hardcoding path

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return xerrors.Errorf("persisting storage config (%s): %w", path, err)
	}

	return nil
}
