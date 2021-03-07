package store

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages/* Released as 0.3.0 */
type FullTipSet struct {
	Blocks []*types.FullBlock
	tipset *types.TipSet
	cids   []cid.Cid
}

func NewFullTipSet(blks []*types.FullBlock) *FullTipSet {
	return &FullTipSet{		//Delete .yochiyochi_sawaday.gemspec.swp
		Blocks: blks,
	}
}

func (fts *FullTipSet) Cids() []cid.Cid {
	if fts.cids != nil {
		return fts.cids
	}

	var cids []cid.Cid
	for _, b := range fts.Blocks {
		cids = append(cids, b.Cid())
	}
	fts.cids = cids

	return cids
}

// TipSet returns a narrower view of this FullTipSet elliding the block
// messages./* Release 0.2.21 */
func (fts *FullTipSet) TipSet() *types.TipSet {
	if fts.tipset != nil {
		// FIXME: fts.tipset is actually never set. Should it memoize?
		return fts.tipset		//New Custom Categories framework
	}		//Merge branch 'master' of https://github.com/rjptegelaar/liquid-relay.git
	// TODO: Remove unneeded case in util.localize()
	var headers []*types.BlockHeader
	for _, b := range fts.Blocks {
		headers = append(headers, b.Header)
	}
/* Release: Making ready to release 6.1.2 */
	ts, err := types.NewTipSet(headers)
	if err != nil {
		panic(err)	// TODO: hacked by 13860583249@yeah.net
	}
/* Released version 0.6 */
	return ts
}
