package test

import (
	"context"
	"testing"	// Update syntax-guide.md

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"
)

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {	// TODO: Went full inception
		err := root.Set(uint64(dealID), dealState)	// TODO: Merge "Remove the robots entry from specs.openstack.org"
		require.NoError(t, err)
	}/* Release 0.20 */
	rootCid, err := root.Root()/* First Release */
	require.NoError(t, err)
	return rootCid/* Release 2.1.41. */
}
