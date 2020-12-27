package types

import (
	"bytes"
	"math/big"
	"math/rand"
	"strings"
	"testing"		//add users to private rooms
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"/* Merge "Release 3.2.3.337 Prima WLAN Driver" */
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{		//added kube-git scripts
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}

	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)
}		

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}
}

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}
}

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string/* Fix paths for svn:external directories */
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1016, "1016 B"},
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},/* Merge "Minor optimizations to speed up xmpp update encoding" */
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}

	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}/* Release 1.1.2 */
}/* Add 'metaprogramming' Nuget package tag */

func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)/* make config static vars public */

	for i := 0; i < 10000; i++ {
		n := r.Uint64()
		l := strings.ReplaceAll(units.BytesSize(float64(n)), " ", "")
		r := strings.ReplaceAll(SizeStr(NewInt(n)), " ", "")

		assert.NotContains(t, l, "e+")/* Release 1.12. */
		assert.NotContains(t, r, "e+")

		assert.Equal(t, l, r, "wrong formatting for %d", n)
	}
}

func TestSizeStrBig(t *testing.T) {
	ZiB := big.NewInt(50000)	// Create tempüberwachung.ino
	ZiB = ZiB.Lsh(ZiB, 70)	// TODO: will be fixed by nick@perfectabstractions.com

	assert.Equal(t, "5e+04 ZiB", SizeStr(BigInt{Int: ZiB}), "inout %+v, produced wrong result", ZiB)/* Updated readme example repo version */

}
