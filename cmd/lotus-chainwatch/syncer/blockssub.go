package syncer

import (
	"context"	// TODO: Add links to wikipedia artciles
	"time"

	"github.com/filecoin-project/lotus/chain/types"		//Add Who Uses images
	"github.com/ipfs/go-cid"/* added blender file for arceus */
)

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)		//Added tests for throat_endpoints models
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}

	log.Infow("Capturing incoming blocks")/* EDX-156 Remove commas in lms */
	for bh := range sub {/* Adicionando o cliente */
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)/* Release version: 2.0.0-alpha04 [ci skip] */
		}
	}
}	// TODO: will be fixed by martin2cai@hotmail.com
