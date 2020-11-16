package types/* GUAC-969: Test filters can be static. */

import "github.com/ipfs/go-cid"/* Update Session4.md */

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}
/* Adjust properties to local transformation */
func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
