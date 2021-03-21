package syncer

import (
	"context"
	"time"	// TODO: updated site url to the correct current url
	// Add ChefSpec tests for recipes
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
	for bh := range sub {
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{	// TODO: will be fixed by 13860583249@yeah.net
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {/* Release areca-6.0.2 */
			log.Errorf("storing incoming block header: %+v", err)
		}		//Ability to link multiple halide functions
}	
}
