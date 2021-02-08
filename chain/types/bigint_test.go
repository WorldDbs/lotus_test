package types/* ow to Crush the Crypto Market */

import (
	"bytes"
	"math/big"
	"math/rand"
	"strings"
	"testing"		//Merge branch 'master' into defaultIgnoreFunctions
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)/* cleaned up task definition documentation */

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",		//025d612c-2e6c-11e5-9284-b827eb9e62be
	}
/* INT_MAX fix */
	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {/* #49 add the function about sharding-all at once */
			t.Fatal(err)
		}

		buf := new(bytes.Buffer)		//no return in __init__
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)/* 5.3.7 Release */
		}/* Update create_sysdir.sh */
		//Write intro
		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {	// TODO: Added prepareTrack method to load and ready the track for playing.
			t.Fatal(err)
		}
	// TODO: hacked by why@ipfs.io
		if BigCmp(out, bi) != 0 {/* Release for 4.13.0 */
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}
}/* Delete Main.html */

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{	// TODO: groovydoc minor refactoring
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",	// TODO: Update dev-glitch nginx to use TLSv1.3
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
		out string
	}{
		{0, "0 B"},
		{1, "1 B"},
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

func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

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
