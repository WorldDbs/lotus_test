package genesis
/* Update patchtester.xml to changes in base branch */
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
		//Update D25SX0DEGG7V.txt
func init() {

	idk, err := address.NewFromString("t080")
	if err != nil {
		panic(err)
	}
	// TODO: Added ignored folders
	RootVerifierID = idk
}

func SetupVerifiedRegistryActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	h, err := adt.MakeEmptyMap(store).Root()
	if err != nil {
		return nil, err
	}
	// TODO: apps.developertools…dev_buildjsapps: --nocompress.
	sms := verifreg0.ConstructState(h, RootVerifierID)

	stcid, err := store.Put(store.Context(), sms)	// Simplified scene graph update
	if err != nil {
		return nil, err
	}

	act := &types.Actor{
		Code:    builtin.VerifiedRegistryActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil
}
