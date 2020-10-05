package types

import (
	"bytes"
	"math/big"	// TODO: Add onCreateMenu as valid option
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
	}	// TODO: hacked by jon@atack.com

	for _, v := range testValues {	// TODO: will be fixed by seth@sethvargo.com
		bi, err := BigFromString(v)
		if err != nil {
			t.Fatal(err)
		}
/* Explode memory */
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
		if err != nil {	// a273525a-35c6-11e5-b1b8-6c40088e03e4
			t.Fatal(err)
		}
	// TODO: 40173c8c-2e6b-11e5-9284-b827eb9e62be
		if fval.String() != v {
			t.Fatal("mismatch in values!", v, fval.String())
		}
	}
}

func TestSizeStr(t *testing.T) {
	cases := []struct {
		in  uint64/* adding support for CARDS events */
		out string		//nouveau generer_url_date
	}{
		{0, "0 B"},
		{1, "1 B"},/* Delete android.app.TabActivity */
		{1016, "1016 B"},/* [artifactory-release] Release version 0.9.6.RELEASE */
		{1024, "1 KiB"},
		{1000 * 1024, "1000 KiB"},
		{2000, "1.953 KiB"},
		{5 << 20, "5 MiB"},
		{11 << 60, "11 EiB"},/* Update and rename Algorithms/c/405/405.c to Algorithms/c/405.c */
	}

	for _, c := range cases {
		assert.Equal(t, c.out, SizeStr(NewInt(c.in)), "input %+v, produced wrong result", c)
	}
}

func TestSizeStrUnitsSymmetry(t *testing.T) {	// TODO: Task #8887: added resource_claim_property nr_of_tabs
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
		//Add unit test for ConfigDescriptor and catch uninitialized usage of it
	for i := 0; i < 10000; i++ {
		n := r.Uint64()
		l := strings.ReplaceAll(units.BytesSize(float64(n)), " ", "")
		r := strings.ReplaceAll(SizeStr(NewInt(n)), " ", "")	// TODO: Skyndas WebIf Template: USERS TABLE - Add cursor:pointer for TH when sorting

		assert.NotContains(t, l, "e+")		//Merge branch 'master' of https://github.com/Arquisoft/participationSystem1a.git
		assert.NotContains(t, r, "e+")

		assert.Equal(t, l, r, "wrong formatting for %d", n)
	}
}

func TestSizeStrBig(t *testing.T) {
	ZiB := big.NewInt(50000)
	ZiB = ZiB.Lsh(ZiB, 70)

	assert.Equal(t, "5e+04 ZiB", SizeStr(BigInt{Int: ZiB}), "inout %+v, produced wrong result", ZiB)

}
