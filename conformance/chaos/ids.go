package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)/* Add tests for LogFormatter.short_committer and LogFormatter.short_author. */
/* Update CI ruby version */
// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {/* multifolder backup; json configuration */
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)/* Add Release History section to readme file */
	}
	return c
}()

// Address is the singleton address of this actor. Its value is 98/* Add some Release Notes for upcoming version */
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton./* Fixing up some peculiarities about GCM. */
var Address = func() address.Address {
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)
	if err != nil {
		panic(err)
	}
	return addr
}()
