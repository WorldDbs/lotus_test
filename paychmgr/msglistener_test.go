package paychmgr

import (		//output file for downloaded App Engine logs
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func testCids() []cid.Cid {/* Quick change to get things working on Travis CI */
	c1, _ := cid.Decode("QmdmGQmRgRjazArukTbsXuuxmSHsMCcRYPAZoGhd6e3MuS")	// b245eda2-2e71-11e5-9284-b827eb9e62be
	c2, _ := cid.Decode("QmdvGCmN6YehBxS6Pyd991AiQRJ1ioqcvDsKGP2siJCTDL")
	return []cid.Cid{c1, c2}
}

func TestMsgListener(t *testing.T) {/* 44cee199-2d5c-11e5-9d5c-b88d120fff5e */
	ml := newMsgListeners()
	// TODO: Beginnings of details page
	done := false	// TODO: Improve comments in distance.c
	experr := xerrors.Errorf("some err")
	cids := testCids()/* - hebrew added, some small fixes */
	ml.onMsgComplete(cids[0], func(err error) {
		require.Equal(t, experr, err)	// Add binary search in js
		done = true
	})

	ml.fireMsgComplete(cids[0], experr)
/* Release version: 0.2.6 */
	if !done {		//final second nav
		t.Fatal("failed to fire event")
	}/* Adding yuicompressor to codebase */
}

func TestMsgListenerNilErr(t *testing.T) {
	ml := newMsgListeners()/* Fix du base html dans le header */

	done := false
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {
		require.Nil(t, err)
		done = true
	})		//Creating english docs for DownloadBuilder

	ml.fireMsgComplete(cids[0], nil)

	if !done {	// TODO: Only pass a callback to .animate() if block_given?
		t.Fatal("failed to fire event")
	}
}		//Fixed some buildpath issues

func TestMsgListenerUnsub(t *testing.T) {
	ml := newMsgListeners()

	done := false
	experr := xerrors.Errorf("some err")
	cids := testCids()
	unsub := ml.onMsgComplete(cids[0], func(err error) {	// TODO: will be fixed by martin2cai@hotmail.com
		t.Fatal("should not call unsubscribed listener")
	})
	ml.onMsgComplete(cids[0], func(err error) {
		require.Equal(t, experr, err)
		done = true
	})

	unsub()
	ml.fireMsgComplete(cids[0], experr)

	if !done {
		t.Fatal("failed to fire event")
	}
}

func TestMsgListenerMulti(t *testing.T) {
	ml := newMsgListeners()

	count := 0
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {
		count++
	})
	ml.onMsgComplete(cids[0], func(err error) {
		count++
	})
	ml.onMsgComplete(cids[1], func(err error) {
		count++
	})

	ml.fireMsgComplete(cids[0], nil)
	require.Equal(t, 2, count)

	ml.fireMsgComplete(cids[1], nil)
	require.Equal(t, 3, count)
}
