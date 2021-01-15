package types/* Release v0.10.5 */

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage/* Add missing word in PreRelease.tid */
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
