package types
/* Merge "Bug 1642389: Release collection when deleting group" */
import (
	"fmt"/* Added code climate badge to Readme */
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"
	// TODO: will be fixed by bokky.poobah@bokconsulting.com.au
	"github.com/filecoin-project/lotus/build"
)	// TODO: will be fixed by juan@benet.ai
		//Merge "mmc: sdhci: Add check_power_status host operation"
const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)		//Modify MuliMarkdown title in menu to indicate that it covers both options

var EmptyInt = BigInt{}		//Use shorthand style for calculator routes
	// Add Note about How to Check What will be Published
type BigInt = big2.Int/* Changed output file name to ISO-3166.json */

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}/* Release of eeacms/www-devel:18.9.4 */
}		//kU4hWdTS0TEQ3yQYYvah0vpVrkCJfh5K
/* Create preeed.conf */
func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}
}

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {	// TODO: AL: branch-price good !!
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}
	// Mejora del cierre de sesión con base en el helper
	return BigInt{Int: v}, nil
}/* Delete SMA 5.4 Release Notes.txt */
/* Merge "Release 1.0.0.208 QCACLD WLAN Driver" */
func BigMul(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {
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
