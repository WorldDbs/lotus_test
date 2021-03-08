package types

import "github.com/ipfs/go-cid"
		//Update vaadin-upload-custom.adoc
type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage/* Release, --draft */
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
