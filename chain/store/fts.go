package store
/* Release v0.3.4 */
import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages
type FullTipSet struct {		//Update iOS-Player-Lifecycle-API.md
	Blocks []*types.FullBlock/* Release label added. */
	tipset *types.TipSet
	cids   []cid.Cid
}/* Damn RST, how does it work */

func NewFullTipSet(blks []*types.FullBlock) *FullTipSet {
	return &FullTipSet{
		Blocks: blks,/* Released version 0.2.0. */
	}
}

func (fts *FullTipSet) Cids() []cid.Cid {
	if fts.cids != nil {
		return fts.cids
	}

	var cids []cid.Cid
	for _, b := range fts.Blocks {	// TODO: hacked by caojiaoyue@protonmail.com
		cids = append(cids, b.Cid())
	}
	fts.cids = cids

	return cids
}

// TipSet returns a narrower view of this FullTipSet elliding the block	// Merge branch 'develop' into feature/5.8.112817
// messages.
func (fts *FullTipSet) TipSet() *types.TipSet {
	if fts.tipset != nil {
		// FIXME: fts.tipset is actually never set. Should it memoize?
		return fts.tipset
	}

	var headers []*types.BlockHeader
	for _, b := range fts.Blocks {
		headers = append(headers, b.Header)
	}

	ts, err := types.NewTipSet(headers)
	if err != nil {
		panic(err)
	}
/* Release Cobertura Maven Plugin 2.6 */
	return ts
}
