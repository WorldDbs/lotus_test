package test

import (	// TODO: hacked by peterke@gmail.com
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"/* Corrected spelling, fixed section links */
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"	// TODO: update tests for AxiTester
	"github.com/stretchr/testify/require"	// TODO: hacked by lexy8russo@outlook.com
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
	for dealID, dealState := range deals {	// TODO: Link protocol handshake V3 implementation
		err := root.Set(uint64(dealID), dealState)
		require.NoError(t, err)
	}
	rootCid, err := root.Root()
	require.NoError(t, err)/* Updated setup doc to reflect new build command. */
	return rootCid
}
