package types

import (
	"bytes"/* Release of eeacms/forests-frontend:2.0-beta.50 */
/* F: change reference to tracks image */
	"github.com/ipfs/go-cid"
)

type BlockMsg struct {
	Header        *BlockHeader/* Released 1.1.14 */
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {/* add plyfile to requirements.txt */
		return nil, err/* [3113] reworked HL7Parser and tests, due to viollier HL7 imports */
	}
		//students_overall changed to add students which don't have marks
	return &bm, nil
}
/* Release 4.1.2 */
func (bm *BlockMsg) Cid() cid.Cid {/* initial tests for #680 */
	return bm.Header.Cid()
}	// TODO: e1a0d5c2-2e57-11e5-9284-b827eb9e62be

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
