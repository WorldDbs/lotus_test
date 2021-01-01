package types

import (
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {	// TODO: Renamed message read method to receive() in WebSocket connection.
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}
/* Merge "Release 4.0.10.23 QCACLD WLAN Driver" */
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err		//added void convertToString(char* cString)
	}

	return &bm, nil/* 0.1 Release. */
}
		//Added a test that modifies the writable partition size on initramfs.
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}/* d773f702-2e73-11e5-9284-b827eb9e62be */
	return buf.Bytes(), nil
}
