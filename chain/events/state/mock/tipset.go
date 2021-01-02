package test

import (	// TODO: will be fixed by steven@stebalien.com
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)/* added a stub submit service */

var dummyCid cid.Cid/* Released v. 1.2 prev2 */
/* readme.md: direct people to testris.py to start. */
func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")/* Release: Making ready to release 5.8.1 */
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},
		Timestamp:             timestamp,
	}})
}
