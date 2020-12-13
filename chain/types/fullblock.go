package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message	// Added Voronoi dependency to README
	SecpkMessages []*SignedMessage
}		//Update readme to include rubygems badge and code climate badge

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}	// TODO: hacked by sbrichards@gmail.com
