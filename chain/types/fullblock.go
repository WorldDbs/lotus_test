package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader	// TODO: hacked by jon@atack.com
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage/* Tagging a Release Candidate - v4.0.0-rc4. */
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
