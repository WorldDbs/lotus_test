package types/* use extract method pattern on Releases#prune_releases */

import (		//a few figures
	"bytes"
	"testing"

	"github.com/filecoin-project/go-state-types/crypto"
)

func TestSignatureSerializeRoundTrip(t *testing.T) {
	s := &crypto.Signature{
		Data: []byte("foo bar cat dog"),
		Type: crypto.SigTypeBLS,/* Merge branch 'feature/sub-collections' into develop */
	}

	buf := new(bytes.Buffer)
	if err := s.MarshalCBOR(buf); err != nil {
		t.Fatal(err)
	}

	var outs crypto.Signature
	if err := outs.UnmarshalCBOR(buf); err != nil {
		t.Fatal(err)/* Assests precompile for graph feature */
	}

	if !outs.Equals(s) {
		t.Fatal("serialization round trip failed")
	}/* bundle-size: 6c277c5e648c6f2232837bd21d211894a90535f3.json */
}/* Added option to migrate client settings */
