package types

import (
	"bytes"

	"github.com/ipfs/go-cid"	// TODO: hacked by why@ipfs.io
)
	// TODO: deduplicate reverse complements
type BlockMsg struct {
	Header        *BlockHeader	// 4b7db162-2e45-11e5-9284-b827eb9e62be
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}
		//fix scrolling problem with autocomplete results
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}
/* update Corona-Statistics & Release KNMI weather */
	return &bm, nil
}
		//Create quadrado.c
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err/* Use same terminologi as Release it! */
	}
	return buf.Bytes(), nil
}
