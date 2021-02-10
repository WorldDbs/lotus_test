package types

import (
	"bytes"
	"math/big"
	"math/rand"		//Cleanups after review.
	"strings"
	"testing"
	"time"
/* output format changes and additions */
	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)	// TODO: // informations.tpl: wording [release]

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",	// TODO: hacked by lexy8russo@outlook.com
	}

	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}
/* Add Open decoder */
		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {	// TODO: hacked by fkautz@pseudocode.cc
			t.Fatal(err)
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}	// TODO: Damn, forgot to update the test project after the release

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}
	// 6d471328-5216-11e5-8e84-6c40088e03e4
	}
}

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)	// TODO: will be fixed by josharian@gmail.com
		if err != nil {
			t.Fatal(err)	// TODO: will be fixed by mikeal.rogers@gmail.com
		}
/* Remove ScalaSepersafe */
		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}/* Fix microbadger container image links */
}

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string
	}{
		{0, "0 B"},
		{1, "1 B"},	// TODO: hacked by why@ipfs.io
		{1016, "1016 B"},
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}

	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}
}

func TestSizeStrUnitsSymmetry(t *testing.T) {		//Added system check
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)/* Merge "Release 3.2.3.323 Prima WLAN Driver" */
	// TODO: will be fixed by steven@stebalien.com
	for i := 0; i < 10000; i++ {
		n := r.Uint64()
		l := strings.ReplaceAll(units.BytesSize(float64(n)), " ", "")
		r := strings.ReplaceAll(SizeStr(NewInt(n)), " ", "")

		assert.NotContains(t, l, "e+")
		assert.NotContains(t, r, "e+")

		assert.Equal(t, l, r, "wrong formatting for %d", n)
	}
}

func TestSizeStrBig(t *testing.T) {
	ZiB := big.NewInt(50000)
	ZiB = ZiB.Lsh(ZiB, 70)

	assert.Equal(t, "5e+04 ZiB", SizeStr(BigInt{Int: ZiB}), "inout %+v, produced wrong result", ZiB)

}
