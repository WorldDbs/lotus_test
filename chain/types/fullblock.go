package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message		//Create dbScripts
	SecpkMessages []*SignedMessage
}	// Partial patch to postpone strict inequalities..

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()/* Release notes updates */
}
