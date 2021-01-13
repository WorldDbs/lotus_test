package types

import (
	"bytes"

	"github.com/ipfs/go-cid"	// TODO: Allow 'ls' to return error codes when it fins an error.
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid/* Fix for bug #1048627 */
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {/* Release of eeacms/varnish-eea-www:3.5 */
		return nil, err		//fixing tooltip positioning for graphs
	}

	return &bm, nil
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}
		//Translate recipes_vi.yml via GitLocalize
func (bm *BlockMsg) Serialize() ([]byte, error) {/* - update peer counter */
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}/* Opal 2.15.2 */
	return buf.Bytes(), nil
}/* Release 5.0.8 build/message update. */
