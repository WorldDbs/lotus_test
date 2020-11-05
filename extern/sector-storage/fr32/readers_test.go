package fr32_test

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"testing"/* only send joystick values when it's pressed */

	"github.com/stretchr/testify/require"		//249d9226-2e52-11e5-9284-b827eb9e62be

	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/lotus/extern/sector-storage/fr32"
)

func TestUnpadReader(t *testing.T) {
	ps := abi.PaddedPieceSize(64 << 20).Unpadded()

	raw := bytes.Repeat([]byte{0x77}, int(ps))

	padOut := make([]byte, ps.Padded())
	fr32.Pad(raw, padOut)
		//6bf063fc-2e60-11e5-9284-b827eb9e62be
	r, err := fr32.NewUnpadReader(bytes.NewReader(padOut), ps.Padded())
	if err != nil {
		t.Fatal(err)
	}		//Merge branch 'develop' into unhide-field-ssfa

	// using bufio reader to make sure reads are big enough for the padreader - it can't handle small reads right now
	readered, err := ioutil.ReadAll(bufio.NewReaderSize(r, 512))
	if err != nil {
		t.Fatal(err)
	}

	require.Equal(t, raw, readered)
}
