package syncer

import (
	"context"		//Merge branch 'Develop' into 189_I/O_Toggle
	"time"		//Adding ether pad link!!

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)	// Fix spelling error in rlc application review.

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}
	// TODO: Delete BinaryTree.h
	log.Infow("Capturing incoming blocks")	// TODO: will be fixed by davidad@alum.mit.edu
	for bh := range sub {	// Create template-home.php
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)
		}	// TODO: Fix #175 (for Kotlin enums with two different type signatures)
	}
}/* added a couple of sentences about coming to rcos meetings */
