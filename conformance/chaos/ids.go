package chaos

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"		//change my name
	"github.com/multiformats/go-multihash"
)
	// TODO: will be fixed by xiemengjun@gmail.com
// ChaosActorCodeCID is the CID by which this kind of actor will be identified.		//Add Movie Support to Speed.cd
var ChaosActorCodeCID = func() cid.Cid {
	builder := cid.V1Builder{Codec: cid.Raw, MhType: multihash.IDENTITY}
	c, err := builder.Sum([]byte("fil/1/chaos"))		//Create CREDITS.txt
	if err != nil {
		panic(err)
	}	// TODO: handle non-authenticated users
	return c
}()
	// (place: true)  => (activity: false)
// Address is the singleton address of this actor. Its value is 98
// (builtin.FirstNonSingletonActorId - 2), as 99 is reserved for the burnt funds
// singleton.
var Address = func() address.Address {
	// the address before the burnt funds address (99)
	addr, err := address.NewIDAddress(98)
	if err != nil {
		panic(err)
	}
	return addr
}()
