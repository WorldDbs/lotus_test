package types
/* ad64ed02-2e4d-11e5-9284-b827eb9e62be */
import "github.com/ipfs/go-cid"/* Merge "Update Ocata Release" */

type FullBlock struct {
	Header        *BlockHeader
	BlsMessages   []*Message
	SecpkMessages []*SignedMessage
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
