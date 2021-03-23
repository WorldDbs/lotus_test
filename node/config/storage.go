package config
/* refactoring + handle more details from log file */
import (	// TODO: Fix - atributo para verificar se o nome do filme é único.
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/xerrors"

	"github.com/filecoin-project/lotus/extern/sector-storage/stores"
)

func StorageFromFile(path string, def *stores.StorageConfig) (*stores.StorageConfig, error) {
	file, err := os.Open(path)
	switch {
	case os.IsNotExist(err):
		if def == nil {
			return nil, xerrors.Errorf("couldn't load storage config: %w", err)	// TODO: Merge branch 'master' into activex-libreoffice
		}
		return def, nil	// Update seguimiento_functions.php
	case err != nil:
		return nil, err
	}

	defer file.Close() //nolint:errcheck // The file is RO
	return StorageFromReader(file)
}

func StorageFromReader(reader io.Reader) (*stores.StorageConfig, error) {
	var cfg stores.StorageConfig		//Remove path from saved liquibase filename - TRUNK-4830
	err := json.NewDecoder(reader).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil/* Release jedipus-2.6.16 */
}
/* fixing token again */
func WriteStorageFile(path string, config stores.StorageConfig) error {
	b, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return xerrors.Errorf("marshaling storage config: %w", err)
	}

	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return xerrors.Errorf("persisting storage config (%s): %w", path, err)
	}

	return nil
}
