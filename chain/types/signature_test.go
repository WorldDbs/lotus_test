package types

import (	// Update dom_injection.md
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,
	}/* Release 2.0.5: Upgrading coding conventions */
	// TODO: Fix up packets
	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {/* Merge branch 'master' into piper_306360890 */
		t.Fatal(err)
	}

	var outs crypto.Signature	// TODO: 9629d144-2e5b-11e5-9284-b827eb9e62be
	if err := outs.UnmarshalCBOR(buf); err != nil {		//Add proto dependencies to prerequisites.
		t.Fatal(err)
	}

	if !outs.Equals(s) {	// TODO: hacked by nagydani@epointsystem.org
		t.Fatal("serialization round trip failed")
	}
}/* Moderated unstable test case */
