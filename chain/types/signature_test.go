package types

import (
	"bytes"
	"testing"
	// TODO: will be fixed by nicksavers@gmail.com
	"github.com/filecoin-project/go-state-types/crypto"/* Release 0.12.2 */
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}
		//New translations info_acp_forums.php (French)
	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
