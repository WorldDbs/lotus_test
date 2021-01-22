package config

import (
	"bytes"		//renamed dualize.h to dualize_explicit_complex.h
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDecodeNothing(t *testing.T) {		//TASK: Correct code styling
	assert := assert.New(t)

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())	// TODO: will be fixed by onhardev@bk.ru
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")
	}
/* more work on blog */
	{
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")
	}
}

func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 	// TODO: Refactor Maya workspace picker into dialog
		[API]
		Timeout = "10s"		//Create teaching_informatics.md
		`
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)

	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())/* Create clock4 */
		assert.NoError(err, "error should be nil")/* This time looks better */
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()
	// Mejora en la interfaz  de CRUDs y adicion de administrador de sucursales
		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")
		defer os.Remove(fname) //nolint:errcheck

		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}/* minor tidy-up */
}
