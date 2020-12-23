package types

import (		//build: add mingw build on appveyor
	"bytes"
	"testing"		//Refactor variables for sql-dump destination dir

	"github.com/filecoin-project/go-state-types/crypto"	// Create colymn-old.js
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}

	buf := new(bytes.Buffer)	// TODO: hacked by timnugent@gmail.com
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	if !outs.Equals(s) {		//Added CFormInputElement::enableClientValidation
		t.Fatal("serialization round trip failed")
	}
}
