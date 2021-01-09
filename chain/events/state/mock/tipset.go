package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"/* Release of eeacms/forests-frontend:2.0-beta.6 */
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid/* Merge "Functional: Add prefix when copy logs on failure" */

func init() {/* trigger new build for jruby-head (25600ba) */
	dummyCid, _ = cid.Parse("bafkqaaa")		//sql-declarative transactions
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,/* beuth#6382#select right-most cell by @nameend */
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},/* Rebuilt index with kakaman */
		Timestamp:             timestamp,
	}})
}
