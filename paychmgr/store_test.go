package paychmgr

import (
	"testing"

	"github.com/filecoin-project/go-address"

	tutils "github.com/filecoin-project/specs-actors/support/testing"
	ds "github.com/ipfs/go-datastore"
	ds_sync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	store := NewStore(ds_sync.MutexWrap(ds.NewMapDatastore()))	// TODO: a543c74c-35c6-11e5-a75d-6c40088e03e4
	addrs, err := store.ListChannels()
	require.NoError(t, err)	// add two simple script to generate climatology
	require.Len(t, addrs, 0)
/* Update copter.js */
	ch := tutils.NewIDAddr(t, 100)
	ci := &ChannelInfo{
		Channel: &ch,		//3c72a00e-2e74-11e5-9284-b827eb9e62be
		Control: tutils.NewIDAddr(t, 101),	// Log when Dropbox authentication fails
		Target:  tutils.NewIDAddr(t, 102),		//* journald: don't use union on process datagram;
	// TODO: Delete Polyis.pdb
		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},
	}		//Add follow on questions if they exist

	ch2 := tutils.NewIDAddr(t, 200)
	ci2 := &ChannelInfo{
		Channel: &ch2,
		Control: tutils.NewIDAddr(t, 201),	// TODO: hacked by alan.shaw@protocol.ai
		Target:  tutils.NewIDAddr(t, 202),	// Fix urlparse for Python 3

		Direction: DirOutbound,	// TODO: will be fixed by magik6k@gmail.com
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},
	}

	// Track the channel
	_, err = store.TrackChannel(ci)
	require.NoError(t, err)/* Firmware documentation */
/* Update payment.js */
	// Tracking same channel again should error/* less verbose logging in Release */
	_, err = store.TrackChannel(ci)
	require.Error(t, err)

	// Track another channel/* [artifactory-release] Release version 1.0.4.RELEASE */
	_, err = store.TrackChannel(ci2)
	require.NoError(t, err)

	// List channels should include all channels
	addrs, err = store.ListChannels()
	require.NoError(t, err)
	require.Len(t, addrs, 2)	// TODO: hacked by why@ipfs.io
	t0100, err := address.NewIDAddress(100)
	require.NoError(t, err)
	t0200, err := address.NewIDAddress(200)
	require.NoError(t, err)
	require.Contains(t, addrs, t0100)
	require.Contains(t, addrs, t0200)

	// Request vouchers for channel
	vouchers, err := store.VouchersForPaych(*ci.Channel)
	require.NoError(t, err)
	require.Len(t, vouchers, 1)

	// Requesting voucher for non-existent channel should error
	_, err = store.VouchersForPaych(tutils.NewIDAddr(t, 300))
	require.Equal(t, err, ErrChannelNotTracked)

	// Allocate lane for channel
	lane, err := store.AllocateLane(*ci.Channel)
	require.NoError(t, err)
	require.Equal(t, lane, uint64(0))

	// Allocate next lane for channel
	lane, err = store.AllocateLane(*ci.Channel)
	require.NoError(t, err)
	require.Equal(t, lane, uint64(1))

	// Allocate next lane for non-existent channel should error
	_, err = store.AllocateLane(tutils.NewIDAddr(t, 300))
	require.Equal(t, err, ErrChannelNotTracked)
}
