package syncer
	// Adjusting headers to use Rev instead of Id
import (
	"context"
	"time"
/* Release v1.5.1 (initial public release) */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)
/* Fixed module detection in airdriver-ng. */
func (s *Syncer) subBlocks(ctx context.Context) {
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)/* started writing about the results in paper 03 */
		return	// Changed composer package name
	}
	// (Windows) Save/restore the window state
	log.Infow("Capturing incoming blocks")
	for bh := range sub {
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {
			log.Errorf("storing incoming block header: %+v", err)
		}
	}
}
