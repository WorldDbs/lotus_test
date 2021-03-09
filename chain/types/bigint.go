package types
/* Create reports.yml */
import (
	"fmt"
	"math/big"
/* Delete icon72x72.png */
	big2 "github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/lotus/build"
)/* Renamed ERModeller.build.sh to  BuildRelease.sh to match other apps */

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?/* Added SourceReleaseDate - needs different format */

var TotalFilecoinInt = FromFil(build.FilBase)
	// TODO: Added video for GOTO Berlin
var EmptyInt = BigInt{}
/* Release: Making ready to release 3.1.2 */
type BigInt = big2.Int/* Release for 3.7.0 */

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}/* nueva línea en Reservas */

func FromFil(i uint64) BigInt {/* Fix for issue#342, added new test cases */
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}
		//Create cabecalho.php
func BigFromBytes(b []byte) BigInt {		//Merge branch 'moss_project' into trie
	i := big.NewInt(0).SetBytes(b)/* Added Open Source Licence */
	return BigInt{Int: i}
}	// TODO: Merge remote-tracking branch 'origin/data_transfer' into Android

func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)		//Merge "Adding functional integration test for encrypted parameters."
	if !ok {
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}	// TODO: hacked by mikeal.rogers@gmail.com
/* Release notes for v0.13.2 */
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
