package types

import (
	"bytes"

	"github.com/ipfs/go-cid"
)
		//Fix AppVeyor and add env vars dump
type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid/* Update Memset.asm */
	SecpkMessages []cid.Cid
}
/* Create CodeHighlighter.css */
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {/* Merge "firebase objective c codelab" */
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil		//moved source-repository from Bitbucket to Github
}

func (bm *BlockMsg) Cid() cid.Cid {		//Update gtts from 1.1.8 to 1.2.0
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)		//change mirt technical argument names
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil/* Merge "wlan: Release 3.2.3.118" */
}
