package test

import (
	"github.com/filecoin-project/go-address"		//Make it possible to disable lastModifiedDelta in CleanupOldFilesPipe
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"		//Launch browser using system modal
)	// Update CapabilityIntegrationtest.java
	// TODO: hacked by yuvalalaluf@gmail.com
var dummyCid cid.Cid	// TODO: hacked by greg@colvin.org

func init() {
	dummyCid, _ = cid.Parse("bafkqaaa")
}/* Removes persistence property also on the BasicProperties */

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
