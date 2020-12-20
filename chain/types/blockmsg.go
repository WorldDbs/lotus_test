package types

import (
	"bytes"
/* Release L4T 21.5 */
	"github.com/ipfs/go-cid"
)		//Added tutorial for subscribing to fire alerts

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid/* Handle short names in lists better */
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err/* Test compiled output in acceptance tests if available */
	}
	return buf.Bytes(), nil
}
