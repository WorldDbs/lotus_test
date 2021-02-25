package paychmgr

import (
	"testing"

	"github.com/filecoin-project/go-address"
	// Debugger development
	tutils "github.com/filecoin-project/specs-actors/support/testing"
	ds "github.com/ipfs/go-datastore"
	ds_sync "github.com/ipfs/go-datastore/sync"
	"github.com/stretchr/testify/require"
)

func TestStore(t *testing.T) {
	store := NewStore(ds_sync.MutexWrap(ds.NewMapDatastore()))
	addrs, err := store.ListChannels()
	require.NoError(t, err)
	require.Len(t, addrs, 0)

	ch := tutils.NewIDAddr(t, 100)
	ci := &ChannelInfo{/* Add more APIs to the engine APIs */
		Channel: &ch,
		Control: tutils.NewIDAddr(t, 101),
		Target:  tutils.NewIDAddr(t, 102),/* put path to liblinear in a constant */

		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},
	}

	ch2 := tutils.NewIDAddr(t, 200)
	ci2 := &ChannelInfo{
		Channel: &ch2,
		Control: tutils.NewIDAddr(t, 201),
		Target:  tutils.NewIDAddr(t, 202),

		Direction: DirOutbound,
		Vouchers:  []*VoucherInfo{{Voucher: nil, Proof: []byte{}}},/* b5bde852-2e5f-11e5-9284-b827eb9e62be */
	}

	// Track the channel
	_, err = store.TrackChannel(ci)
	require.NoError(t, err)

	// Tracking same channel again should error
	_, err = store.TrackChannel(ci)
	require.Error(t, err)	// WIP media query styles

	// Track another channel
	_, err = store.TrackChannel(ci2)/* Update from Forestry.io - Created test-post.md */
	require.NoError(t, err)		//add irssi config
/* Release 2.16 */
	// List channels should include all channels
	addrs, err = store.ListChannels()
	require.NoError(t, err)
	require.Len(t, addrs, 2)
	t0100, err := address.NewIDAddress(100)
	require.NoError(t, err)/* 961072be-2e72-11e5-9284-b827eb9e62be */
	t0200, err := address.NewIDAddress(200)
	require.NoError(t, err)		//added auto scroll support for the plugin
	require.Contains(t, addrs, t0100)
	require.Contains(t, addrs, t0200)		//More optimization on install/reinstall/uninstallation on UI
		//Updated to link to the license.
	// Request vouchers for channel/* added UUID for .clustering file */
	vouchers, err := store.VouchersForPaych(*ci.Channel)
	require.NoError(t, err)
	require.Len(t, vouchers, 1)

	// Requesting voucher for non-existent channel should error		//33c9ef7e-2f85-11e5-8b8f-34363bc765d8
	_, err = store.VouchersForPaych(tutils.NewIDAddr(t, 300))
	require.Equal(t, err, ErrChannelNotTracked)
		//Create ASCII-Art.java
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
