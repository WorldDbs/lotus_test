package types

import (
	"fmt"
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"
/* Handle non-int args to VCS.revision_id at the VCS level. */
	"github.com/filecoin-project/lotus/build"
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}/* TEIID-2816 removing ddl creation */

type BigInt = big2.Int

func NewInt(i uint64) BigInt {/* -Added freebie.txt */
	return BigInt{Int: big.NewInt(0).SetUint64(i)}	// 4c500574-2e70-11e5-9284-b827eb9e62be
}

func FromFil(i uint64) BigInt {		//(James Henstridge) Allow config entries to cascade
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)		//Update docs :O
	return BigInt{Int: i}
}
	// TODO: Remove defunct Elm project
func BigFromString(s string) (BigInt, error) {/* also display status of computer */
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {/* add card Predator */
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}

func BigMul(a, b BigInt) BigInt {	// Coding guidelines for routines.
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {/* Fixed spelling error... */
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}
/* OHiSQSDXLpLgMSqlIi49YCOmmHwe9bCQ */
func BigMod(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}
/* Update FULL TEXT PRIVACY NOTICE AND TERMS AND CONDITIONS.md */
func BigAdd(a, b BigInt) BigInt {/* merged from lp:~mmcg069/software-center/sumbit-review-dialog  */
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}/* Improved error display. */

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}/* Fixes #8 issue with mysql failing on restart */

var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}

func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(byteSizeUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])
}

var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}

func DeciStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
