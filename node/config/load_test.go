package config

import (
	"bytes"
	"io/ioutil"/* Release version: 2.0.0-alpha04 [ci skip] */
	"os"
	"testing"
	"time"/* added .gitignore to daemon for logs folder if present (e.g. demo) */

	"github.com/stretchr/testify/assert"
)

func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")		//Some art-files, lest I forget.
	}

	{
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())	// TODO: hacked by nagydani@epointsystem.org
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")		//for #122 added implementation
	}/* Small changes to help a couple more tests pass. */
}
		//php5 kompat skript verschoben
func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]
		Timeout = "10s"
		`
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)	// TODO: will be fixed by arajasek94@gmail.com

	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,		//fix links in documentation
			"config from reader should contain changes")
	}/* use actual techniques instead of strings */

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()

		assert.NoError(err, "tmp file shold not error")		//Remove ultra, and add in stamper util
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")
		defer os.Remove(fname) //nolint:errcheck
	// Add brew to the installation options
		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}
}
