package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"		//Merge branch 'master' into node-623-feedback

	"github.com/stretchr/testify/assert"
)
	// TODO: hacked by hugomrdias@gmail.com
func TestDecodeNothing(t *testing.T) {/* Reverted endpoint encoding changes. */
	assert := assert.New(t)

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())/* Release of eeacms/www-devel:20.8.1 */
		assert.Nil(err, "error should be nil")	// Zip including the Windows binary of v1.0.0
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")
	}

	{
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")
	}
}

func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]
		Timeout = "10s"
		`	// Display cantrips in ability section
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)/* Merge branch 'release/2.15.0-Release' */

	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}		//Update ProjectTest.php

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()

		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)/* Release Notes for v00-14 */
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")
		defer os.Remove(fname) //nolint:errcheck

		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}
}
