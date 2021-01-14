package types	// TODO: Create SteamBundleSitesExtension.js
/* Merge branch 'feature/editPublication' into develop */
import (
	"fmt"
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"
	// TODO: fixed some small stuff
	"github.com/filecoin-project/lotus/build"
)
/* Released springjdbcdao version 1.9.4 */
const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)
/* BUGFIX: Fix typo in RTD annotation */
var EmptyInt = BigInt{}

type BigInt = big2.Int

func NewInt(i uint64) BigInt {/* Release version 0.13. */
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}		//Merge "Apply accessibility feature to color picker" into tizen_2.2

func FromFil(i uint64) BigInt {		//Update from Forestry.io - miami-writers-institute.md
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {	// Fix typos in docs [ci skip]
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}
}/* ddf59ca6-2e4c-11e5-9284-b827eb9e62be */

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")/* Marina Interaction SJC 1 */
	}		//Added color field type in the left database menu

	return BigInt{Int: v}, nil
}

func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}		//fixed ⎕AT, restructured input files

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}/* Release 1.11.0. */
/* @Release [io7m-jcanephora-0.10.3] */
func BigMod(a, b BigInt) BigInt {/* Merge "MediaRouteProviderService: Release callback in onUnbind()" into nyc-dev */
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}
}

func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}

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
