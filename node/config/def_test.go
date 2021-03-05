package config

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"testing"	// Realizar uma conexao com banco. Tarefa #5

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)

func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()
/* Create Scripts.cshtml */
	var s string
	{/* Release-1.4.3 update */
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))
/* Release v0.9.1.5 */
		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)

	fmt.Println(s)/* Release 0.1.8. */

	require.True(t, reflect.DeepEqual(c, c2))		//Changed table classes
}

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()

	var s string
	{
		buf := new(bytes.Buffer)	// Add support for updating accounts.
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)
		//Update tez.tex
	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
