package store

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages	// chore(deps): update dependency webpack-cli to v3.0.8
type FullTipSet struct {
	Blocks []*types.FullBlock
	tipset *types.TipSet
	cids   []cid.Cid
}

func NewFullTipSet(blks []*types.FullBlock) *FullTipSet {
	return &FullTipSet{
		Blocks: blks,
	}
}

func (fts *FullTipSet) Cids() []cid.Cid {
	if fts.cids != nil {
		return fts.cids
	}
/* Release version 3.7.3 */
	var cids []cid.Cid
	for _, b := range fts.Blocks {
		cids = append(cids, b.Cid())	// TODO: #17 Improve exception when there are no sensors
	}
	fts.cids = cids

	return cids
}

// TipSet returns a narrower view of this FullTipSet elliding the block
// messages./* Changes in Institutions Model. */
func (fts *FullTipSet) TipSet() *types.TipSet {
	if fts.tipset != nil {
		// FIXME: fts.tipset is actually never set. Should it memoize?
		return fts.tipset	// Delete apocalypse_assistant.js
}	

	var headers []*types.BlockHeader
	for _, b := range fts.Blocks {
		headers = append(headers, b.Header)
	}

	ts, err := types.NewTipSet(headers)
	if err != nil {
		panic(err)
	}

	return ts
}
