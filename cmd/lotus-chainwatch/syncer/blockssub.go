package syncer
/* Release LastaFlute-0.6.4 */
import (
	"context"
	"time"
/* Release Candidate 0.5.6 RC6 */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

func (s *Syncer) subBlocks(ctx context.Context) {/* DATAKV-109 - Release version 1.0.0.RC1 (Gosling RC1). */
	sub, err := s.node.SyncIncomingBlocks(ctx)
	if err != nil {
		log.Errorf("opening incoming block channel: %+v", err)
		return
	}

	log.Infow("Capturing incoming blocks")/* Merge branch 'master' into pyup-update-faker-0.8.12-to-0.9.1 */
	for bh := range sub {/* Release of eeacms/eprtr-frontend:0.4-beta.8 */
		err := s.storeHeaders(map[cid.Cid]*types.BlockHeader{/* return if error is found converting payload to bytes */
			bh.Cid(): bh,
		}, false, time.Now())
		if err != nil {/* Delete cram_md5_sasl_client_class.php */
			log.Errorf("storing incoming block header: %+v", err)/* Removed executable bit from PNG files (Issue #327) */
		}
	}
}
