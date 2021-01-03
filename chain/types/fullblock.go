package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}
/* Add web browser requirements */
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}	// No need to mention the obvious
