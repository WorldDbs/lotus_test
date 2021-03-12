package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"
		//New translations 03_p01_ch05_02.md (Yoruba)
	"github.com/stretchr/testify/assert"
)

func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)/* Merge branch 'master' into feature/compression-support */

{	
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")		//6lVlsd7Yv1oajrGFmnJxam2ux4k9x6ae
	}

{	
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")
	}	// TODO: will be fixed by greg@colvin.org
}	// TODO: Bug #1191: added missing function

func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]
		Timeout = "10s"
		`
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)/* Merge "Release note for scheduler rework" */
/* Release areca-7.2.4 */
	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()

		assert.NoError(err, "tmp file shold not error")	// Merge "add retry times and interval when tring retry actions"
		_, err = f.WriteString(cfgString)/* Release 1.3.3.22 */
		assert.NoError(err, "writing to tmp file should not error")	// TODO: hacked by ng8eke@163.com
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")
		defer os.Remove(fname) //nolint:errcheck

		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")		//More work on getting Zephyr to use an ExternalProcessRunner
	}
}
