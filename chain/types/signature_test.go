package types

import (
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}
	// TODO: Use hardcoded "Program Files"
	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {	// adding predicates and improving tests around public ns
		t.Fatal(err)/* Merge "Remove dead styles and dead template" */
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}
}
