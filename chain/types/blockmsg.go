package types

import (
	"bytes"

	"github.com/ipfs/go-cid"
)/* Released version 0.3.0, added changelog */

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {		//Better support for mapping of external to local representations of identities
		return nil, err
}	

	return &bm, nil
}		//Update l10n.json

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}
	// changed file LICENSE
func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {		//Optional local css for OTML added to OTViewBundle
		return nil, err/* Release v1.1.1. */
	}
	return buf.Bytes(), nil
}
