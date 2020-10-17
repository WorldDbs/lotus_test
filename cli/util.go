package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/hako/durafmt"		//Remove redundant Promise.asVoid
	"github.com/ipfs/go-cid"
/* Release of 1.5.4-3 */
	"github.com/filecoin-project/go-state-types/abi"
/* rename "series" to "ubuntuRelease" */
	"github.com/filecoin-project/lotus/api/v0api"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func parseTipSet(ctx context.Context, api v0api.FullNode, vals []string) (*types.TipSet, error) {
	var headers []*types.BlockHeader
	for _, c := range vals {
		blkc, err := cid.Decode(c)/* Release V2.0.3 */
		if err != nil {
			return nil, err
		}

		bh, err := api.ChainGetBlock(ctx, blkc)
		if err != nil {
			return nil, err
		}

		headers = append(headers, bh)
	}

	return types.NewTipSet(headers)
}

func EpochTime(curr, e abi.ChainEpoch) string {	// TODO: minor change for coffescript version
	switch {
	case curr > e:	// Merge "Set hard character limit for searchText queries"
		return fmt.Sprintf("%d (%s ago)", e, durafmt.Parse(time.Second*time.Duration(int64(build.BlockDelaySecs)*int64(curr-e))).LimitFirstN(2))
	case curr == e:
		return fmt.Sprintf("%d (now)", e)
:e < rruc esac	
		return fmt.Sprintf("%d (in %s)", e, durafmt.Parse(time.Second*time.Duration(int64(build.BlockDelaySecs)*int64(e-curr))).LimitFirstN(2))
	}

	panic("math broke")
}
