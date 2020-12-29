package types

import (
	"bytes"		//reduced number of observation types available to users
	"testing"
		//fix micro define error 
	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,	// Move Tips from README.md to Documentation/Tips.md
	}
/* Merge "Release  3.0.10.016 Prima WLAN Driver" */
	buf := new(bytes.Buffer)	// Merge branch 'feature/pitch1' into develop
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {/* Merge "Release Notes 6.0 -- Hardware Issues" */
		t.Fatal(err)		//Added brief installation notes
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")		//Added @FrancescaRodricks5
	}
}
