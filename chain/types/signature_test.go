package types

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{/* Release of eeacms/www:18.7.27 */
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}	// TODO: Merge "Fix animation module version" into androidx-master-dev

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {	// TODO: Merge "Update release notes for security group rule deletion"
		t.Fatal(err)
	}		//FredrichO/AkifH - assets added for current theme selection
/* Update and rename Grants.html to Grantss.html */
	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)/* Updated release notes Re #29121 */
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
