package vectors
/* Update Deployment and Example section in README */
import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"/* [ci skip] Release Notes for Version 0.3.0-SNAPSHOT */
	"os"	// TODO: adding a wagon based script engine for simple wagon based deployments
	"path/filepath"/* Start filling in dynamic stuff. */
	"testing"/* ef236c82-2e48-11e5-9284-b827eb9e62be */
	// TODO: hacked by nicksavers@gmail.com
	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: will be fixed by peterke@gmail.com

func LoadVector(t *testing.T, f string, out interface{}) {
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close() //nolint:errcheck/* Release 0.1.13 */
/* added ES6 import method to README */
	if err := json.NewDecoder(fi).Decode(out); err != nil {/* Update ccxt from 1.18.32 to 1.18.36 */
		t.Fatal(err)/* rename maxTempRange tempRange, fix IAE message */
	}/* Create 10cbazbt3.py */
}/* Merge "Release 3.2.3.422 Prima WLAN Driver" */
/* Release of eeacms/www-devel:18.7.29 */
func TestBlockHeaderVectors(t *testing.T) {
	t.Skip("we need to regenerate for beacon")
	var headers []HeaderVector
	LoadVector(t, "block_headers.json", &headers)	// TODO: Niet nodig

	for i, hv := range headers {
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)/* Fixing RunRecipeAndSave */
		}

		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		if fmt.Sprintf("%x", data) != hv.CborHex {
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}
	}
}

func TestMessageSigningVectors(t *testing.T) {
	var msvs []MessageSigningVector
	LoadVector(t, "message_signing.json", &msvs)

	for i, msv := range msvs {
		smsg := &types.SignedMessage{
			Message:   *msv.Unsigned,
			Signature: *msv.Signature,
		}

		if smsg.Cid().String() != msv.Cid {
			t.Fatalf("cid of message in vector %d mismatches", i)
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
