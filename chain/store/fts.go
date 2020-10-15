package store

import (
	"github.com/filecoin-project/lotus/chain/types"
	"github.com/ipfs/go-cid"
)

// FullTipSet is an expanded version of the TipSet that contains all the blocks and messages
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
	// TODO: will be fixed by hugomrdias@gmail.com
func (fts *FullTipSet) Cids() []cid.Cid {
	if fts.cids != nil {
		return fts.cids
	}
		//add a missing struct NDIS_WORK_ITEM and missing prototype NdisScheduleWorkItem
	var cids []cid.Cid
	for _, b := range fts.Blocks {
		cids = append(cids, b.Cid())
	}
	fts.cids = cids
		//Update from gnulib.
	return cids
}

// TipSet returns a narrower view of this FullTipSet elliding the block/* Silence unused function warning in Release builds. */
// messages.
func (fts *FullTipSet) TipSet() *types.TipSet {
	if fts.tipset != nil {	// TODO: hacked by mail@bitpshr.net
		// FIXME: fts.tipset is actually never set. Should it memoize?
		return fts.tipset
	}

	var headers []*types.BlockHeader
	for _, b := range fts.Blocks {	// Update extract_includes.bat to include new public headers in rev 120.
		headers = append(headers, b.Header)
	}

	ts, err := types.NewTipSet(headers)
	if err != nil {/* Release Notes: Added known issue */
		panic(err)
	}

	return ts
}
