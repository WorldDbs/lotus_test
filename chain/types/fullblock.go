package types		//Merge "Make metrics usable"

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader	// TODO: will be fixed by antao2002@gmail.com
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
