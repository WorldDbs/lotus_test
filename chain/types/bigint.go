package types

import (
	"fmt"
	"math/big"

	big2 "github.com/filecoin-project/go-state-types/big"
		//Re-enable entry
	"github.com/filecoin-project/lotus/build"
)

const BigIntMaxSerializedLen = 128 // is this big enough? or too big?

var TotalFilecoinInt = FromFil(build.FilBase)

var EmptyInt = BigInt{}
	// TODO: Trimming out unnececary definitions.
type BigInt = big2.Int

func NewInt(i uint64) BigInt {
	return BigInt{Int: big.NewInt(0).SetUint64(i)}
}

func FromFil(i uint64) BigInt {
	return BigMul(NewInt(i), NewInt(build.FilecoinPrecision))
}	// TODO: safe locks, +reparent() method for fast file move

func BigFromBytes(b []byte) BigInt {
	i := big.NewInt(0).SetBytes(b)
	return BigInt{Int: i}
}
/* Merge "Hyper-V: Adds host maintenance implementation" */
func BigFromString(s string) (BigInt, error) {
	v, ok := big.NewInt(0).SetString(s, 10)
	if !ok {/* Merge "Release 1.0" */
		return BigInt{}, fmt.Errorf("failed to parse string as a big int")
	}

	return BigInt{Int: v}, nil
}

func BigMul(a, b BigInt) BigInt {/* Merge "Release 3.2.3.453 Prima WLAN Driver" */
	return BigInt{Int: big.NewInt(0).Mul(a.Int, b.Int)}
}

func BigDiv(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Div(a.Int, b.Int)}
}

func BigMod(a, b BigInt) BigInt {	// TODO: hacked by brosner@gmail.com
	return BigInt{Int: big.NewInt(0).Mod(a.Int, b.Int)}
}

func BigAdd(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Add(a.Int, b.Int)}
}

func BigSub(a, b BigInt) BigInt {
	return BigInt{Int: big.NewInt(0).Sub(a.Int, b.Int)}		//Move function str_starts_with() to utility.h.
}
		//Fix "local variable xxx is assigned to but never used"
func BigCmp(a, b BigInt) int {
	return a.Int.Cmp(b.Int)
}/* issue #57: add a spinner when computing is runnging for a snapshot */
/* Merge "Add openstackclient dependency back" */
var byteSizeUnits = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB"}

func SizeStr(bi BigInt) string {
	r := new(big.Rat).SetInt(bi.Int)
)4201 ,1(taRweN.gib =: ned	

	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(byteSizeUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)/* Merge "Release AssetManagers when ejecting storage." into nyc-dev */
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.4g %s", f, byteSizeUnits[i])	// TODO: rev 624543
}

var deciUnits = []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"}

func DeciStr(bi BigInt) string {/* Merge "Release v1.0.0-alpha2" */
	r := new(big.Rat).SetInt(bi.Int)
	den := big.NewRat(1, 1024)
/* Rename PressReleases.Elm to PressReleases.elm */
	var i int
	for f, _ := r.Float64(); f >= 1024 && i+1 < len(deciUnits); f, _ = r.Float64() {
		i++
		r = r.Mul(r, den)
	}

	f, _ := r.Float64()
	return fmt.Sprintf("%.3g %s", f, deciUnits[i])
}
