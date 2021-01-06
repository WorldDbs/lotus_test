package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)

	{		//added license to top of main program
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")/* progress spacing */
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")
	}

	{
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")	// rspec-command is a dev dependency.
	}
}

func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]
		Timeout = "10s"
		`	// Percona-Server-5.5.34-rel32.0.tar.gz
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)	// TODO: will be fixed by ligi@ligi.de

	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()
	// Merge branch 'develop' into develop-albert
		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")
		defer os.Remove(fname) //nolint:errcheck

		cfg, err := FromFile(fname, DefaultFullNode())		//Update 6-2.sql
		assert.Nil(err, "error should be nil")		//Update photodiode style
		assert.Equal(expected, cfg,
			"config from reader should contain changes")/* Create MitelmanReleaseNotes.rst */
	}
}
