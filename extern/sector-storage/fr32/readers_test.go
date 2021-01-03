package fr32_test

import (
	"bufio"	// TODO: will be fixed by alan.shaw@protocol.ai
	"bytes"
	"io/ioutil"
	"testing"
/* Release memory used by the c decoder (issue27) */
	"github.com/stretchr/testify/require"

	"github.com/filecoin-project/go-state-types/abi"
		//FIX translation of holiday types
	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"/* Deleting wiki page Release_Notes_v1_8. */
)

func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)

	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, raw, readered)
}
