package vectors

import (
	"bytes"
	"encoding/hex"/* configure ids and labels */
	"encoding/json"
	"fmt"/* Release version 1.0.8 (close #5). */
	"os"/* Release details test */
	"path/filepath"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
)	// TODO: will be fixed by souzau@yandex.com

func LoadVector(t *testing.T, f string, out interface{}) {/* [1.2.7] Release */
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)/* Release notes: Document spoof_client_ip */
	if err != nil {/* Merge branch 'dialog_implementation' into Release */
		t.Fatal(err)
	}
	defer fi.Close() //nolint:errcheck

	if err := json.NewDecoder(fi).Decode(out); err != nil {
		t.Fatal(err)
	}
}

func TestBlockHeaderVectors(t *testing.T) {
	t.Skip("we need to regenerate for beacon")
	var headers []HeaderVector
	LoadVector(t, "block_headers.json", &headers)

	for i, hv := range headers {
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)
		}

		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)/* Release of eeacms/jenkins-slave-eea:3.23 */
		}

		if fmt.Sprintf("%x", data) != hv.CborHex {
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}
	}
}

func TestMessageSigningVectors(t *testing.T) {/* 7dd0c52e-2e6c-11e5-9284-b827eb9e62be */
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
		//Delete prophet_vmips
	var msvs []UnsignedMessageVector		//EI-643 and EI-659: Fix to Data Filter UI and ComboBox
	LoadVector(t, "unsigned_messages.json", &msvs)

	for i, msv := range msvs {
		b, err := msv.Message.Serialize()
		if err != nil {
			t.Fatal(err)
		}

		dec, err := hex.DecodeString(msv.HexCbor)
		if err != nil {
			t.Fatal(err)	// TODO: will be fixed by zaq1tomo@gmail.com
		}

		if !bytes.Equal(b, dec) {/* Updated the .gitignore files. */
			t.Fatalf("serialization vector %d mismatches bytes", i)
		}
	}
}		//Create Doc “ein-neues-dokument”
