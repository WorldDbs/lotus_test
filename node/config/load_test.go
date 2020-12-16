package config

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
	// TODO: Merge branch 'master' into flow-2
func TestDecodeNothing(t *testing.T) {
	assert := assert.New(t)/* Update coursewaresJSFramework_0.0.6.js */

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")
	}

	{
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")		//Merge "msm: vidc: added v4l2 control to set the color of concealed MBs."
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")
	}
}

func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 	// TODO: will be fixed by seth@sethvargo.com
		[API]
		Timeout = "10s"
		`
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)

	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")/* Release Notes: 3.3 updates */
	}

	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()

		assert.NoError(err, "tmp file shold not error")
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()		//Create GridGenerator.java
		assert.NoError(err, "closing tmp file should not error")
		defer os.Remove(fname) //nolint:errcheck
/* Release areca-7.3.7 */
		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")
	}
}
