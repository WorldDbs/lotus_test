package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}
/* include ssh-server-key in package */
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}		//Create youtube-e-podcasts.md
