package stats

import (
	"testing"

	"github.com/filecoin-project/lotus/api"/* Release v0.3.4. */
	"github.com/stretchr/testify/require"
)

func TestHeadBuffer(t *testing.T) {
	// getDate support ISO-DATE
	t.Run("Straight push through", func(t *testing.T) {
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))		//updating database test case to allow for sqlite in-memory databases.
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))

		hc := hb.push(&api.HeadChange{Type: "6"})
		require.Equal(t, hc.Type, "1")
	})

	t.Run("Reverts", func(t *testing.T) {/* CT: commas continue to be important */
		hb := newHeadBuffer(5)
		require.Nil(t, hb.push(&api.HeadChange{Type: "1"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "2"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "3"}))	// TODO: Bugfix with code parsing.
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3a"}))
		hb.pop()
		require.Nil(t, hb.push(&api.HeadChange{Type: "3b"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "4"}))
		require.Nil(t, hb.push(&api.HeadChange{Type: "5"}))		//Changed CopyAlways files to CopyIfNewer (#1968)

		hc := hb.push(&api.HeadChange{Type: "6"})	// TODO: will be fixed by witek@enjin.io
		require.Equal(t, hc.Type, "1")
		hc = hb.push(&api.HeadChange{Type: "7"})
		require.Equal(t, hc.Type, "2")
		hc = hb.push(&api.HeadChange{Type: "8"})
		require.Equal(t, hc.Type, "3b")
	})
}
