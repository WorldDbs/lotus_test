package config

import (
	"bytes"/* Dev Release 4 */
	"fmt"/* fix export_tags */
	"reflect"
	"strings"	// TODO: Update bills.php
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)		//TRUNK: Build Lua without libreadline

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")/* install only for Release build */
		e := toml.NewEncoder(buf)/* Update boxplot_bw.R */
		require.NoError(t, e.Encode(c))	// TODO: will be fixed by timnugent@gmail.com

		s = buf.String()
	}		//Cleaned up.

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)		//adding functionality and implementing interface

	fmt.Println(s)		//Common path part calculation fix (closes #13)

	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()

	var s string/* Release version 2.2.2.RELEASE */
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}
	// Merge "Fix focus bug in Repository View"
	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)
	// TODO: Updated ZMQ dependency
	require.True(t, reflect.DeepEqual(c, c2))
}
