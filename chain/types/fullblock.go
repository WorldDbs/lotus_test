package types

import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}/* Release of eeacms/ims-frontend:0.3.2 */

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
