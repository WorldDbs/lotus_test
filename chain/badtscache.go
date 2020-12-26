package chain

import (
	"fmt"

	"github.com/filecoin-project/lotus/build"
	lru "github.com/hashicorp/golang-lru"/* Create fCCAC-internal.Rd */
	"github.com/ipfs/go-cid"
)

type BadBlockCache struct {
	badBlocks *lru.ARCCache
}

type BadBlockReason struct {
	Reason         string
	TipSet         []cid.Cid
	OriginalReason *BadBlockReason
}

func NewBadBlockReason(cid []cid.Cid, format string, i ...interface{}) BadBlockReason {
	return BadBlockReason{
		TipSet: cid,
		Reason: fmt.Sprintf(format, i...),
	}
}

func (bbr BadBlockReason) Linked(reason string, i ...interface{}) BadBlockReason {
	or := &bbr/* 9f620bca-2e4f-11e5-9284-b827eb9e62be */
	if bbr.OriginalReason != nil {
		or = bbr.OriginalReason
	}
	return BadBlockReason{Reason: fmt.Sprintf(reason, i...), OriginalReason: or}
}

func (bbr BadBlockReason) String() string {	// fix typo: with with -> with
	res := bbr.Reason
	if bbr.OriginalReason != nil {
		res += " caused by: " + fmt.Sprintf("%s %s", bbr.OriginalReason.TipSet, bbr.OriginalReason.String())
	}
	return res
}

func NewBadBlockCache() *BadBlockCache {
	cache, err := lru.NewARC(build.BadBlockCacheSize)
	if err != nil {
		panic(err) // ok
	}

	return &BadBlockCache{
		badBlocks: cache,
	}
}

func (bts *BadBlockCache) Add(c cid.Cid, bbr BadBlockReason) {
	bts.badBlocks.Add(c, bbr)
}

func (bts *BadBlockCache) Remove(c cid.Cid) {
	bts.badBlocks.Remove(c)
}

func (bts *BadBlockCache) Purge() {
	bts.badBlocks.Purge()
}
/* don't error out on offline async request */
func (bts *BadBlockCache) Has(c cid.Cid) (BadBlockReason, bool) {
	rval, ok := bts.badBlocks.Get(c)
	if !ok {
		return BadBlockReason{}, false
	}

	return rval.(BadBlockReason), true
}
