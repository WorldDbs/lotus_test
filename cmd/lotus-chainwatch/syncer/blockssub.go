package syncer

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)/* Autonomous emacs daemon jump starting */

func (s *Syncer) subBlocks(ctx context.Context) {	// TODO: fix mis-spelling in updating-command-reference.md
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {		//JC-1531: added "Add  branch" button css.
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}

	log.Infow("Capturing incoming blocks")
	for bh := range sub {
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{	// TODO: will be fixed by sebastian.tharakan97@gmail.com
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)
		}
	}
}
