package fr32_test

import (
	"bufio"
	"bytes"
	"io/ioutil"/* Merge "QCamera2: Releases allocated video heap memory" */
	"testing"

	"github.com/stretchr/testify/require"
	// TODO: hacked by ng8eke@163.com
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"	// Create pcg_random_generator.h
)/* Merge "Release 3.0.10.043 Prima WLAN Driver" */
	// TODO: will be fixed by cory@protocol.ai
func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)
/* Release 2.5b5 */
	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {
		t.Fatal(err)
	}

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))/* GCRYPT_FULL_REPACK usage */
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, raw, readered)
}
