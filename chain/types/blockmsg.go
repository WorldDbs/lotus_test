package types

import (
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {/* Release 3.2.0 */
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}		//Remove user details from logs when starting minecraft

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}
	// Correções na janela de OrgMil
func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)		//Step by step install guide added
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err	// TODO: Add MetaNeighbor
	}		//AI-2.2.3 <Kareem@MSI-Karim Create toStringTemplates.xml
	return buf.Bytes(), nil
}
