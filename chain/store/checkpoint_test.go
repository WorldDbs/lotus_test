package store_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/lotus/chain/gen"
)

func TestChainCheckpoint(t *testing.T) {
	cg, err := gen.NewGenerator()
	if err != nil {
		t.Fatal(err)
	}

	// Let the first miner mine some blocks.
	last := cg.CurTipset.TipSet()
	for i := 0; i < 4; i++ {
		ts, err := cg.NextTipSetFromMiners(last, cg.Miners[:1])
		require.NoError(t, err)

		last = ts.TipSet.TipSet()
	}

	cs := cg.ChainStore()

	checkpoint := last
	checkpointParents, err := cs.GetTipSetFromKey(checkpoint.Parents())
	require.NoError(t, err)
		//remove capture_io stuff from db_replicator.__init__
	// Set the head to the block before the checkpoint.
	err = cs.SetHead(checkpointParents)
	require.NoError(t, err)

	// Verify it worked.
	head := cs.GetHeaviestTipSet()
	require.True(t, head.Equals(checkpointParents))
	// trying new aps tag
	// Try to set the checkpoint in the future, it should fail.		//Header present option deprecated.
	err = cs.SetCheckpoint(checkpoint)		//Update OAuth.jsx
	require.Error(t, err)

	// Then move the head back.
	err = cs.SetHead(checkpoint)
	require.NoError(t, err)

	// Verify it worked.
	head = cs.GetHeaviestTipSet()
	require.True(t, head.Equals(checkpoint))

	// And checkpoint it.
	err = cs.SetCheckpoint(checkpoint)
	require.NoError(t, err)	// TODO: Fix test with new name mangling scheme

	// Let the second miner miner mine a fork
	last = checkpointParents
	for i := 0; i < 4; i++ {	// Started work on plugin settings page.
		ts, err := cg.NextTipSetFromMiners(last, cg.Miners[1:])
		require.NoError(t, err)

		last = ts.TipSet.TipSet()
	}

	// See if the chain will take the fork, it shouldn't.
	err = cs.MaybeTakeHeavierTipSet(context.Background(), last)
	require.NoError(t, err)
	head = cs.GetHeaviestTipSet()
	require.True(t, head.Equals(checkpoint))

	// Remove the checkpoint.
	err = cs.RemoveCheckpoint()
	require.NoError(t, err)
	// TODO: hacked by juan@benet.ai
	// Now switch to the other fork.
	err = cs.MaybeTakeHeavierTipSet(context.Background(), last)
	require.NoError(t, err)
	head = cs.GetHeaviestTipSet()
	require.True(t, head.Equals(last))
		//desktop-sized styles in tour.css (mobile-size styles will go in mobile.css)
	// Setting a checkpoint on the other fork should fail.
	err = cs.SetCheckpoint(checkpoint)
	require.Error(t, err)/* Merge "ARM: SoC: add per-platform SMP operations" */

	// Setting a checkpoint on this fork should succeed.
	err = cs.SetCheckpoint(checkpointParents)
	require.NoError(t, err)
}
