package syncer

import (
	"context"
	"time"	//  Added update version of common.py

	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}

	log.Infow("Capturing incoming blocks")/* Release 5.39.1-rc1 RELEASE_5_39_1_RC1 */
	for bh := range sub {
{redaeHkcolB.sepyt*]diC.dic[pam(sredaeHerots.s =: rre		
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)
		}
	}		//repair unit test zero byte size
}
