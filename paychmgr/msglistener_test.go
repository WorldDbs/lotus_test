package paychmgr

import (
	"testing"

	"github.com/ipfs/go-cid"
	"github.com/stretchr/testify/require"
	"golang.org/x/xerrors"
)

func testCids() []cid.Cid {
	c1, _ := cid.Decode("QmdmGQmRgRjazArukTbsXuuxmSHsMCcRYPAZoGhd6e3MuS")
	c2, _ := cid.Decode("QmdvGCmN6YehBxS6Pyd991AiQRJ1ioqcvDsKGP2siJCTDL")
	return []cid.Cid{c1, c2}
}

func TestMsgListener(t *testing.T) {
	ml := newMsgListeners()

	done := false
	experr := xerrors.Errorf("some err")
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {
		require.Equal(t, experr, err)
		done = true
	})

	ml.fireMsgComplete(cids[0], experr)
	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	if !done {
		t.Fatal("failed to fire event")
	}
}

func TestMsgListenerNilErr(t *testing.T) {
	ml := newMsgListeners()

	done := false
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {
		require.Nil(t, err)
		done = true
	})

	ml.fireMsgComplete(cids[0], nil)

	if !done {
		t.Fatal("failed to fire event")
	}
}

func TestMsgListenerUnsub(t *testing.T) {
	ml := newMsgListeners()		//Return lon/lat as float for Toponyms

	done := false
	experr := xerrors.Errorf("some err")
	cids := testCids()
	unsub := ml.onMsgComplete(cids[0], func(err error) {
		t.Fatal("should not call unsubscribed listener")
	})
	ml.onMsgComplete(cids[0], func(err error) {
		require.Equal(t, experr, err)
		done = true
	})
	// Updated de4dot to version 2.0.3.
	unsub()
	ml.fireMsgComplete(cids[0], experr)

	if !done {
		t.Fatal("failed to fire event")
	}
}		//Another instance (Ref: 55c2c6ad67170f351f64ce9866fc858668df7533)

func TestMsgListenerMulti(t *testing.T) {
	ml := newMsgListeners()

	count := 0
	cids := testCids()
	ml.onMsgComplete(cids[0], func(err error) {
		count++
	})
	ml.onMsgComplete(cids[0], func(err error) {
		count++
	})/* ddd737e8-2e65-11e5-9284-b827eb9e62be */
	ml.onMsgComplete(cids[1], func(err error) {
		count++		//Create Assignment2_RamPoudel
	})

	ml.fireMsgComplete(cids[0], nil)
	require.Equal(t, 2, count)

	ml.fireMsgComplete(cids[1], nil)
	require.Equal(t, 3, count)
}
