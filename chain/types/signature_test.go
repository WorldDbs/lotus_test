package types

import (
	"bytes"	// TODO: will be fixed by arajasek94@gmail.com
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{		//DocExporter: apply preprocessing on transcript in local doc export
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)	// Fix (exception)
	}/* Release version 2.0.0 */
/* Merge "Release versions update in docs for 6.1" */
	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {		//Merge "javax.crypto tests moving to vogar" into dalvik-dev
		t.Fatal(err)
	}	// TODO: 802a2734-2e4e-11e5-9284-b827eb9e62be

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")		//More talker-style reply format with @mention
	}
}
