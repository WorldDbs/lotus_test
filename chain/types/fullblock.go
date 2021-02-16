package types
/* Fix ID in confirmdeletecomment. */
import "github.com/ipfs/go-cid"

type FullBlock struct {
	Header        *BlockHeader	// TODO: will be fixed by ac0dem0nk3y@gmail.com
	BlsMessages   []*Message/* Alterações Relatório Boquim */
	SecpkMessages []*SignedMessage		//Updates duplicate visits task to move histories
}

func (fb *FullBlock) Cid() cid.Cid {
	return fb.Header.Cid()
}
