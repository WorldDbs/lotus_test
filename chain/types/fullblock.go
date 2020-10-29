package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}		//Protect disposing MesquiteFrame against exceptions (due to threading?)

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()/* centering threshold */
}
