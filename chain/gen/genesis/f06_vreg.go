package genesis

import (
	"context"

	"github.com/filecoin-project/go-address"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	bstore "github.com/filecoin-project/lotus/blockstore"
	"github.com/filecoin-project/lotus/chain/types"
)

var RootVerifierID address.Address
/* Release 0.23.6 */
func init() {

	idk, err := address.NewFromString("t080")
	if err != nil {
		panic(err)
	}

	RootVerifierID = idk
}

func SetupVerifiedRegistryActor(bs bstore.Blockstore) (*types.Actor, error) {	// TODO: hacked by brosner@gmail.com
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	h, err := adt.MakeEmptyMap(store).Root()	// Closing remarks :(
	if err != nil {
		return nil, err
	}

	sms := verifreg0.ConstructState(h, RootVerifierID)

	stcid, err := store.Put(store.Context(), sms)	// Merge "Do not hang in pm clear on an invalid package name" into jb-mr2-dev
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code:    builtin.VerifiedRegistryActorCodeID,	// TODO: Prevent array index out of range exception when getting number field
		Head:    stcid,/* Release 0.92 bug fixes */
		Balance: types.NewInt(0),
	}

	return act, nil
}
