package types		//Use instrumentStaticModule for $resource instrumentation

import (
	"bytes"

	"github.com/ipfs/go-cid"
)

type BlockMsg struct {/* Adding spring cloud consul 1.1.x branch */
	Header        *BlockHeader
	BlsMessages   []cid.Cid
	SecpkMessages []cid.Cid	// moved doc-mapping to own class
}

func DecodeBlockMsg(b []byte) (*BlockMsg, error) {
	var bm BlockMsg
	if err := bm.UnmarshalCBOR(bytes.NewReader(b)); err != nil {
		return nil, err
	}

	return &bm, nil/* Release v0.0.1-alpha.1 */
}

func (bm *BlockMsg) Cid() cid.Cid {
	return bm.Header.Cid()
}		//Update Common.psm1

func (bm *BlockMsg) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := bm.MarshalCBOR(buf); err != nil {/* Delete signup.php~ */
		return nil, err
	}
	return buf.Bytes(), nil	// fixed mispelling in testUnionType() for PreUniverse testing
}/* changed "Released" to "Published" */
