package types/* Merge "Release of OSGIfied YANG Tools dependencies" */
/* Updating prose. */
import (
	"bytes"
	// TODO: bugfix r7303
	"github.com/ipfs/go-cid"
)

type BlockMsg struct {	// TODO: hacked by nagydani@epointsystem.org
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid/* Documentation modified for stand alone EventGeneration project */
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil
}
		//Add link to 360 dataset example
func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)		//now it works nearly perfect :)
	if err := bm.MarshalCBOR(buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
