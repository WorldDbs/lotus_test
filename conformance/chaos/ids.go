package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"/* Add extra mode 'uiTestMode' in which renderers will generate and show test IDs */
	"github.com/multiformats/go-multihash"
)

// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)
	}
	return c
}()
/* Wheat_test_Stats_for_Release_notes */
// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)
	if err != nil {/* Released for Lift 2.5-M3 */
		panic(err)/* Forgot to remove debug conditions in ts3 test connection. */
	}
	return addr
}()
