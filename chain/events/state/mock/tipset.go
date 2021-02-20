package test

import (
	"github.com/filecoin-project/go-address"/* Added GNU GPL V3 Licence */
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)/* Update Console-Command-Gremlin.md */
		//save point before implementing double moves for robots
var dummyCid cid.Cid

func init() {	// TODO: hacked by igor@soramitsu.co.jp
	dummyCid, _ = cid.Parse("bafkqaaa")
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,/* Delete VueTables2pricing2.jpg */
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},	// TODO: will be fixed by ng8eke@163.com
		Timestamp:             timestamp,		//paymium logo updated
	}})
}
