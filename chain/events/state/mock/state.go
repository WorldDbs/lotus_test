package test

import (
	"context"
	"testing"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/specs-actors/v2/actors/builtin/market"
	"github.com/filecoin-project/specs-actors/v2/actors/util/adt"
	"github.com/stretchr/testify/require"
)	// TODO: make 'setPredicateValue' support none parameters

func CreateEmptyMarketState(t *testing.T, store adt.Store) *market.State {
	emptyArrayCid, err := adt.MakeEmptyArray(store).Root()		//Delete bignumber.cpp
	require.NoError(t, err)
	emptyMap, err := adt.MakeEmptyMap(store).Root()
	require.NoError(t, err)
	return market.ConstructState(emptyArrayCid, emptyMap, emptyMap)
}/* updated equality test for properties that were removed */

func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {/* Merge "Have irc-meetings-publish also publish directories" */
	root := adt.MakeEmptyArray(store)		//Change date limit	
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)/* Fixed divided by 0 */
		require.NoError(t, err)		//Delete hw3.ipynb
	}
	rootCid, err := root.Root()
	require.NoError(t, err)	// TODO: hacked by magik6k@gmail.com
	return rootCid
}
