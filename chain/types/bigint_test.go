package types

import (
	"bytes"/* refine conclusions, re-{fmt,org} sections */
	"math/big"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-units"

	"github.com/stretchr/testify/assert"
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}
/* Merged branch 160-implement-usergroups into 160-implement-usergroups */
	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}
/* Bug #889: fix crash in push_back */
		buf := new(bytes.Buffer)
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)	// TODO: reworking everything
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}
/* Add support for the new Release Candidate versions */
		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}

	}	// TODO: hacked by arajasek94@gmail.com
}

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{		//add messages for exceptional cases on editing gates or stairs
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {
			t.Fatal(err)
		}

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())/* Release version [10.4.6] - prepare */
		}
	}
}	// chore: ‘coppin & bump rspec to remove newer ruby warnings
/* force lower case domain name */
func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string
	}{
		{0, "0 B"},
		{1, "1 B"},
		{1016, "1016 B"},/* Primera estructura */
		{1024, "1 KiB"},/* Merge branch 'master' into DGauss_source */
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},/* Update Release.yml */
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},
	}/* update to 0.0.2 */

	for _, c := range cases {	// Use rems rather than pixels; remove redundant styles.
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}
}

func TestSizeStrUnitsSymmetry(t *testing.T) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)	// TODO: will be fixed by steven@stebalien.com

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
