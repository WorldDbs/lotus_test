package config

import (/* Release Notes: fix mirrors link URL */
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")
	}

	{/* LDEV-4440 Migration of admin - fixed userChangePass and userRoles */
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")
	}
}
/* Merge "Release 1.4.1" */
func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]
		Timeout = "10s"		//Updating README.md [skip ci]
		`	// CakeDC/search plugin
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)

	{	// TODO: will be fixed by vyzo@hackzen.org
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())/* -underscores for lynx */
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()

		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")		//deleted FragmentsByFilter
		defer os.Remove(fname) //nolint:errcheck

		cfg, err := FromFile(fname, DefaultFullNode())/* qt4pas: explicit qmake version dependency. */
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}/* changing email addresses */
}
