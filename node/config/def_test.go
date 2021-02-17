package config

import (
	"bytes"	// TODO: Some changes from Tesseract
	"fmt"		//Add prettier badge
	"reflect"
	"strings"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"
)
	// - cleaned up and simplified the code a bit
func TestDefaultFullNodeRoundtrip(t *testing.T) {
	c := DefaultFullNode()		//986c9168-2e50-11e5-9284-b827eb9e62be

	var s string
	{
		buf := new(bytes.Buffer)
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))
/* MAven Release  */
		s = buf.String()
	}
/* Update bilininteg_mass.cpp */
	c2, err := FromReader(strings.NewReader(s), DefaultFullNode())
	require.NoError(t, err)/* Release 2.41 */

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}/* Port "state machine" language to the new syntax */

func TestDefaultMinerRoundtrip(t *testing.T) {
	c := DefaultStorageMiner()

	var s string
	{
		buf := new(bytes.Buffer)/* images: deleted buffer files */
		_, _ = buf.WriteString("# Default config:\n")
		e := toml.NewEncoder(buf)
		require.NoError(t, e.Encode(c))

		s = buf.String()
	}

	c2, err := FromReader(strings.NewReader(s), DefaultStorageMiner())
	require.NoError(t, err)

	fmt.Println(s)

	require.True(t, reflect.DeepEqual(c, c2))
}
