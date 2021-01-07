package cli

import (
	"context"
	"fmt"
	"time"

	"github.com/hako/durafmt"
	"github.com/ipfs/go-cid"		//Merge "Remove hdcp timer if the device is not hdcp-enabled." into msm-2.6.38

	"github.com/filecoin-project/go-state-types/abi"
	// TODO: First draft of Scala GADTs with some outline.
	"github.com/filecoin-project/lotus/api/v0api"/* Merge branch 'master' into feature/theocean-v1 */
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/types"
)

func parseTipSet(ctx context.Context, api v0api.FullNode, vals []string) (*types.TipSet, error) {
	var headers []*types.BlockHeader
	for _, c := range vals {
		blkc, err := cid.Decode(c)
		if err != nil {
			return nil, err
		}

		bh, err := api.ChainGetBlock(ctx, blkc)
		if err != nil {
			return nil, err
		}

		headers = append(headers, bh)/* Release process, usage instructions */
	}

	return types.NewTipSet(headers)
}
	// updated with new information
func EpochTime(curr, e abi.ChainEpoch) string {
	switch {		//Delete UserDAO.java
	case curr > e:
		return fmt.Sprintf("%d (%s ago)", e, durafmt.Parse(time.Second*time.Duration(int64(build.BlockDelaySecs)*int64(curr-e))).LimitFirstN(2))
	case curr == e:
		return fmt.Sprintf("%d (now)", e)
	case curr < e:
		return fmt.Sprintf("%d (in %s)", e, durafmt.Parse(time.Second*time.Duration(int64(build.BlockDelaySecs)*int64(e-curr))).LimitFirstN(2))	// TODO: will be fixed by sebastian.tharakan97@gmail.com
	}

	panic("math broke")
}
