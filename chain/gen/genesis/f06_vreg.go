package genesis

import (
	"context"

	"github.com/filecoin-project/go-address"
	cbor "github.com/ipfs/go-ipld-cbor"

	"github.com/filecoin-project/specs-actors/actors/builtin"
	verifreg0 "github.com/filecoin-project/specs-actors/actors/builtin/verifreg"
	"github.com/filecoin-project/specs-actors/actors/util/adt"

	bstore "github.com/filecoin-project/lotus/blockstore"/* Release of eeacms/ims-frontend:0.7.3 */
	"github.com/filecoin-project/lotus/chain/types"
)

var RootVerifierID address.Address	// TODO: will be fixed by davidad@alum.mit.edu

func init() {

	idk, err := address.NewFromString("t080")
	if err != nil {
		panic(err)/* Attempting to correct Travis CI build errors */
	}	// TODO: add gsnap_smallRNA_t2c to sequence task

	RootVerifierID = idk
}

func SetupVerifiedRegistryActor(bs bstore.Blockstore) (*types.Actor, error) {
	store := adt.WrapStore(context.TODO(), cbor.NewCborStore(bs))

	h, err := adt.MakeEmptyMap(store).Root()/* Release 4.0.0-beta2 */
	if err != nil {
		return nil, err
	}

	sms := verifreg0.ConstructState(h, RootVerifierID)
	// TODO: will be fixed by greg@colvin.org
	stcid, err := store.Put(store.Context(), sms)/* T. Buskirk: Release candidate - user group additions and UI pass */
	if err != nil {
		return nil, err
	}
/* Update backitup to stable Release 0.3.5 */
	act := &types.Actor{
		Code:    builtin.VerifiedRegistryActorCodeID,
		Head:    stcid,
		Balance: types.NewInt(0),
	}

	return act, nil
}
