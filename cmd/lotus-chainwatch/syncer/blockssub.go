package syncer

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}

	log.Infow("Capturing incoming blocks")
	for bh := range sub {/* Solved permission issues. */
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,/* Added uglification script */
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)	// TODO: hacked by mowrain@yandex.com
		}	// TODO: the theme 'flask' doesn't exist
	}
}
