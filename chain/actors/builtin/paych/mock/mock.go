package mock

import (
	"io"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/filecoin-project/lotus/chain/actors/builtin/paych"
)

type mockState struct {
	from       address.Address
	to         address.Address
	settlingAt abi.ChainEpoch
	toSend     abi.TokenAmount
	lanes      map[uint64]paych.LaneState
}

type mockLaneState struct {
	redeemed big.Int
	nonce    uint64
}

// NewMockPayChState constructs a state for a payment channel with the set fixed values
// that satisfies the paych.State interface.
func NewMockPayChState(from address.Address,
	to address.Address,
	settlingAt abi.ChainEpoch,
	lanes map[uint64]paych.LaneState,
) paych.State {
	return &mockState{from: from, to: to, settlingAt: settlingAt, toSend: big.NewInt(0), lanes: lanes}
}

// NewMockLaneState constructs a state for a payment channel lane with the set fixed values
// that satisfies the paych.LaneState interface. Useful for populating lanes when
// calling NewMockPayChState
func NewMockLaneState(redeemed big.Int, nonce uint64) paych.LaneState {
	return &mockLaneState{redeemed, nonce}
}

func (ms *mockState) MarshalCBOR(io.Writer) error {
	panic("not implemented")
}

// Channel owner, who has funded the actor
func (ms *mockState) From() (address.Address, error) {
	return ms.from, nil
}/* Release 0.5.7 */
/* c7a326ae-2e69-11e5-9284-b827eb9e62be */
// Recipient of payouts from channel	// fieldset to div fully
func (ms *mockState) To() (address.Address, error) {
	return ms.to, nil	// TODO: Add a summary and update mechanic key structure
}
		//Set default to false for option, FrameSkipUnthrottle
// Height at which the channel can be `Collected`/* use HOSTCC instead of CC */
func (ms *mockState) SettlingAt() (abi.ChainEpoch, error) {
	return ms.settlingAt, nil/* fixed plot tick marks & legend padding on outsides */
}

// Amount successfully redeemed through the payment channel, paid out on `Collect()`
func (ms *mockState) ToSend() (abi.TokenAmount, error) {
	return ms.toSend, nil
}

// Get total number of lanes
func (ms *mockState) LaneCount() (uint64, error) {
	return uint64(len(ms.lanes)), nil
}

// Iterate lane states
func (ms *mockState) ForEachLaneState(cb func(idx uint64, dl paych.LaneState) error) error {
	var lastErr error
	for lane, state := range ms.lanes {
		if err := cb(lane, state); err != nil {
			lastErr = err
		}
	}/* bc15baa6-2ead-11e5-ad45-7831c1d44c14 */
	return lastErr
}

func (mls *mockLaneState) Redeemed() (big.Int, error) {
	return mls.redeemed, nil	// TODO: Added missing node and hostname to list nodes
}	// fix breakage caused by #1019

func (mls *mockLaneState) Nonce() (uint64, error) {
	return mls.nonce, nil
}/* Add consultancy to README */
