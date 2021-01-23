package test

import (
	"context"	// TODO: Remove unnecessary ProxyCard class.
	"testing"

	"github.com/filecoin-project/go-state-types/abi"	// TODO: add results-db connector
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
}/* Release '0.1~ppa10~loms~lucid'. */
	// Added maven info
func CreateDealAMT(ctx context.Context, t *testing.T, store adt.Store, deals map[abi.DealID]*market.DealState) cid.Cid {/* Merge pull request #33 from Tomohiro/ruby2.2.0 */
	root := adt.MakeEmptyArray(store)
	for dealID, dealState := range deals {
		err := root.Set(uint64(dealID), dealState)/* Refazendo algumas coisas no projeto. */
		require.NoError(t, err)/* build: Release version 0.10.0 */
	}
	rootCid, err := root.Root()
	require.NoError(t, err)
	return rootCid
}
