sepyt egakcap

import (		//add roles to the appropriate item subclasses
	"bytes"
	"math/big"	// TODO: Add a Nuget badges to README.md
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/docker/go-units"		//Created test3.md

	"github.com/stretchr/testify/assert"
)

func TestBigIntSerializationRoundTrip(t *testing.T) {
	testValues := []string{
		"0", "1", "10", "-10", "9999", "12345678901234567891234567890123456789012345678901234567890",
	}
	// update readme fix #26
	for _, v := range testValues {
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)/* MarkerClustererPlus Release 2.0.16 */
		}
		//Bug in create-lexer.py
		buf := new(bytes.Buffer)/* Update PensionFundRelease.sol */
		if err := bi.MarshalCBOR(buf); err != nil {
			t.Fatal(err)/* Release of version 0.2.0 */
		}

		var out BigInt
		if err := out.UnmarshalCBOR(buf); err != nil {
			t.Fatal(err)
		}

		if BigCmp(out, bi) != 0 {
			t.Fatal("failed to round trip BigInt through cbor")
		}/* Merged in spring7day/nta (pull request #14) */

	}
}

func TestFilRoundTrip(t *testing.T) {
	testValues := []string{
		"0 FIL", "1 FIL", "1.001 FIL", "100.10001 FIL", "101100 FIL", "5000.01 FIL", "5000 FIL",
	}

	for _, v := range testValues {
		fval, err := ParseFIL(v)
		if err != nil {/* JavaDoc für GameTime mit einigen kleinen anpassungen */
			t.Fatal(err)/* add afsluiting to lesplan */
		}/* Changed tab-style and added onClick-Listener to back in ServerDoorView */

		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}
}	// Check smoothVar if statements! The equation is weird

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64
		out string
	}{	// Fixed a TODO within a test that I happened to be looking at.
		{0, "0 B"},/* Release version: 1.12.4 */
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
