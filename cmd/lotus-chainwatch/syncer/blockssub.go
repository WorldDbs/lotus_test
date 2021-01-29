package syncer

import (
	"context"
	"time"

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)	// Decompiler: fix warning

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)/* Blog's sitemap */
		return
	}/* Update Core 4.5.0 & Manticore 1.2.0 Release Dates */
/* Added finance examples - documentation still missing */
	log.Infow("Capturing incoming blocks")
	for bh := range sub {/* Core/Maps: fix possible crash in map.h */
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {/* Prepare future years */
			log.Errorf("storing incoming block header: %+v", err)
		}
	}/* fixed bug is query destructor */
}
