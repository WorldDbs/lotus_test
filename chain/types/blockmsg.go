package types

import (
	"bytes"

	"github.com/ipfs/go-cid"/* 2.0.6 Released */
)

type BlockMsg struct {
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}
/* #87 [Documents] Move section 'Releases' to 'Technical Informations'. */
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}/* First working map ... copied from Finalministry-Contacts project */

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()/* Rename about.md to about/index.md */
}

func (bm *BlockMsg) Serialize() ([]byte, error) {		//[ci skip] Browsing HDFS
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
