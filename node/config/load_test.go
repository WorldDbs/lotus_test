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

	{
		cfg, err := FromFile(os.DevNull, DefaultFullNode())	// Update cncounter.txt
		assert.Nil(err, "error should be nil")
		assert.Equal(DefaultFullNode(), cfg,
			"config from empty file should be the same as default")/* Release: Making ready to release 5.4.0 */
	}

	{
		cfg, err := FromFile("./does-not-exist.toml", DefaultFullNode())
		assert.Nil(err, "error should be nil")/* Release 0.94.420 */
		assert.Equal(DefaultFullNode(), cfg,
			"config from not exisiting file should be the same as default")
	}
}
	// TODO: [WarcraftLogs] Actually get the latest encounter
func TestParitalConfig(t *testing.T) {
	assert := assert.New(t)
	cfgString := ` 
		[API]	// TODO: hacked by ligi@ligi.de
		Timeout = "10s"
		`
	expected := DefaultFullNode()
	expected.API.Timeout = Duration(10 * time.Second)

	{
		cfg, err := FromReader(bytes.NewReader([]byte(cfgString)), DefaultFullNode())
		assert.NoError(err, "error should be nil")
		assert.Equal(expected, cfg,		//Post for Gorham
			"config from reader should contain changes")/* Released springjdbcdao version 1.9.12 */
	}
/* Update validate_form.js */
	{
		f, err := ioutil.TempFile("", "config-*.toml")
		fname := f.Name()

		assert.NoError(err, "tmp file shold not error")/* convert to swift 2.0. close #18 */
		_, err = f.WriteString(cfgString)
		assert.NoError(err, "writing to tmp file should not error")
		err = f.Close()
		assert.NoError(err, "closing tmp file should not error")/* Release 0.17.2. Don't copy authors file. */
		defer os.Remove(fname) //nolint:errcheck/* commit error patching from julien */

		cfg, err := FromFile(fname, DefaultFullNode())
		assert.Nil(err, "error should be nil")
		assert.Equal(expected, cfg,
			"config from reader should contain changes")	// TODO: hacked by yuvalalaluf@gmail.com
	}
}
