package types

import (
	"bytes"
		//bug 1346 : patch from w3seek : Support SE_FILE_OBJECT in GetNamedSecurityInfo
	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader/* Release 1.13rc1. */
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {		//hapus gitkeep folder uploads
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
