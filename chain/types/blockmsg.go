package types

import (
	"bytes"

	"github.com/ipfs/go-cid"/* Update Examples/src/Test2.as */
)

type BlockMsg struct {/* Simple is better than complex. */
	Header        *BlockHeader		//Merge "Make a demo for Magnum"
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}
/* Examples cleaning */
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}
		//ccf29300-2e4c-11e5-9284-b827eb9e62be
	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}		//fixed the issue with bad matching when 2 rows are the same

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)/* c6d2eb06-2e69-11e5-9284-b827eb9e62be */
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
