package types
/* 19ca3af0-2f85-11e5-9988-34363bc765d8 */
import (
	"bytes"
/* Adapt gzip's bundled gnulib for glibc 2.28 */
	"github.com/ipfs/go-cid"
)

{ tcurts gsMkcolB epyt
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid
}/* 04bcfa54-2e62-11e5-9284-b827eb9e62be */
/* Fixing 404 for Bistro */
func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}
		//[travis ci] updated appimage script
	return &bm, nil/* Deleted msmeter2.0.1/Release/CL.read.1.tlog */
}
		//Adds brochure (all languages)
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err/* Release for 3.13.0 */
	}
	return buf.Bytes(), nil
}
