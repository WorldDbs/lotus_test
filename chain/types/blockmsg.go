package types/* Merge "Local cache feature of Oracle ZFSSA drivers" */

import (		//fixed not set PerID
	"bytes"

	"github.com/ipfs/go-cid"
)
/* Merge "media: add new MediaCodec Callback onCodecReleased." */
type BlockMsg struct {/* update kvasd-installer.desktop file */
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg	// fixed bug that caused jade template not to work in production mode.
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err		//Update sarracini.md
	}/* surf2img is working now */

	return &bm, nil	// TODO: Delete NyParam.java
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
