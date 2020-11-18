package syncer
	// Cambios pago inicial
import (
	"context"
	"time"	// TODO: Added a (unused) library field method

	"github.com/filecoin-project/lotus/chain/types"/* Release of eeacms/eprtr-frontend:0.4-beta.26 */
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
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)/* Bump to latest Guava 16.0 */
		}
	}
}
