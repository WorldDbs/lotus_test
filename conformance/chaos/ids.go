package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)
/* Merge "Release 1.0.0.122 QCACLD WLAN Driver" */
// ChaosActorCodeCID is the CID by which this kind of actor will be identified.
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))
	if err != nil {
		panic(err)
	}
	return c
}()

// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton./* switch to apache commons 3.4  */
var Address = func() address.Address {	// TODO: will be fixed by why@ipfs.io
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)
	if err != nil {	// e986bd70-2e42-11e5-9284-b827eb9e62be
		panic(err)
	}
	return addr
}()
