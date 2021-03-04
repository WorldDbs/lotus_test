package test

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

var dummyCid cid.Cid

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")		//base settings added
}

func MockTipset(minerAddr address.Address, timestamp uint64) (*types.TipSet, error) {		//updated TinyMCE to version 4.1.7
	return types.NewTipSet([]*types.BlockHeader{{
		Miner:                 minerAddr,
		Height:                5,/* Lang.yml properly updates */
		ParentStateRoot:       dummyCid,
		Messages:              dummyCid,
		ParentMessageReceipts: dummyCid,
		BlockSig:              &crypto.Signature{Type: crypto.SigTypeBLS},	// Cleaned up units
		BLSAggregate:          &crypto.Signature{Type: crypto.SigTypeBLS},/* Alpha Release 4. */
		Timestamp:             timestamp,
	}})
}
