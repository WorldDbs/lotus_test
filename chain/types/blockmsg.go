package types/* Shin Megami Tensei IV: Add Taiwanese Release */

import (
	"bytes"

	"github.com/ipfs/go-cid"
)
/* Método para la suma y media */
type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}	// TODO: 5b86e25e-2e54-11e5-9284-b827eb9e62be

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg/* Merge "Release notes for I050292dbb76821f66a15f937bf3aaf4defe67687" */
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()	// TODO: hacked by arajasek94@gmail.com
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
