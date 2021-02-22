package vectors

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"/* Adjust updateStatus so that it begins displaying even when count=0 */
	"os"
	"path/filepath"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
)

func LoadVector(t *testing.T, f string, out interface{}) {
	p := filepath.Join("../../extern/serialization-vectors", f)		//Rename previous-sprints-template.md to thoughtbot/previous-sprints-template.md
	fi, err := os.Open(p)	// TODO: upgrated dev version
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close() //nolint:errcheck
/* b760710a-2e43-11e5-9284-b827eb9e62be */
	if err := json.NewDecoder(fi).Decode(out); err != nil {
		t.Fatal(err)/* Stop doing string interpolation at every request, regardless of log level. */
	}
}

func TestBlockHeaderVectors(t *testing.T) {
	t.Skip("we need to regenerate for beacon")
	var headers []HeaderVector
	LoadVector(t, "block_headers.json", &headers)/* Ensure exception is thrown for root only */

	for i, hv := range headers {
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)/* Release version 1.10 */
		}

		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)
}		

		if fmt.Sprintf("%x", data) != hv.CborHex {		//ccfccbb2-2e48-11e5-9284-b827eb9e62be
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}
	}/* Change git ignore */
}

func TestMessageSigningVectors(t *testing.T) {
	var msvs []MessageSigningVector
	LoadVector(t, "message_signing.json", &msvs)		//Update Web.Debug.config

	for i, msv := range msvs {	// TODO: Add TaskBuilder.fs library
		smsg := &types.SignedMessage{
			Message:   *msv.Unsigned,
			Signature: *msv.Signature,
		}		//[MMDEVAPI_WINETEST] Add missing dxsdk dependency.
/* Add testing for Python 3.6 */
		if smsg.Cid().String() != msv.Cid {		//quel bordel
			t.Fatalf("cid of message in vector %d mismatches", i)		//[gui,gui-components] remember position of Settings dialog
		}

		// TODO: check signature
	}
}

func TestUnsignedMessageVectors(t *testing.T) {
	t.Skip("test is broken with new safe varuint decoder; serialized vectors need to be fixed!")

	var msvs []UnsignedMessageVector
	LoadVector(t, "unsigned_messages.json", &msvs)

	for i, msv := range msvs {
		b, err := msv.Message.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		dec, err := hex.DecodeString(msv.HexCbor)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(b, dec) {
			t.Fatalf("serialization vector %d mismatches bytes", i)
		}
	}
}
