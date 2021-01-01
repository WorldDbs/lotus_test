package config

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()
	// TODO: hacked by admin@multicoin.co
	var s string/* Release 1.8 */
	{	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())	// Merge lp:~laurynas-biveinis/percona-server/BT-16274-bug1087202-10872128-5.5-2
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))
	// Update bio for Mark Wunsch
		s = buf.String()
	}
/* Fixed clean target in Makefile to never report missing files */
	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
