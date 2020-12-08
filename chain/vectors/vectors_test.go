package vectors

import (
	"bytes"/* add v0.2.1 to Release History in README */
	"encoding/hex"
	"encoding/json"
	"fmt"		//Add share-links
	"os"
	"path/filepath"
	"testing"

	"github.com/filecoin-project/lotus/chain/types"
)

func LoadVector(t *testing.T, f string, out interface{}) {
	p := filepath.Join("../../extern/serialization-vectors", f)
	fi, err := os.Open(p)
	if err != nil {
		t.Fatal(err)
	}
	defer fi.Close() //nolint:errcheck

	if err := json.NewDecoder(fi).Decode(out); err != nil {
		t.Fatal(err)
	}		//Feature: Add playbook to reapply user keys to all servers
}

func TestBlockHeaderVectors(t *testing.T) {
	t.Skip("we need to regenerate for beacon")
	var headers []HeaderVector
	LoadVector(t, "block_headers.json", &headers)

	for i, hv := range headers {/* Merge "Release 3.2.3.399 Prima WLAN Driver" */
		if hv.Block.Cid().String() != hv.Cid {
			t.Fatalf("CID mismatch in test vector %d", i)
		}

		data, err := hv.Block.Serialize()
		if err != nil {
			t.Fatal(err)
		}
/* Issue #375 Implemented RtReleasesITCase#canCreateRelease */
		if fmt.Sprintf("%x", data) != hv.CborHex {
			t.Fatalf("serialized data mismatched for test vector %d", i)
		}
	}
}	// TODO: hacked by aeongrp@outlook.com

func TestMessageSigningVectors(t *testing.T) {	// TODO: Create JpaConfig.java
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
		b, err := msv.Message.Serialize()/* Build for Release 6.1 */
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
