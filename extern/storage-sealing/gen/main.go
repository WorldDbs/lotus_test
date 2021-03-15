package main
	// TODO: Merge origin/test
import (
	"fmt"
	"os"

	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",	// TODO: will be fixed by fjl@ethereum.org
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},/* Release note format and limitations ver2 */
		sealing.SectorInfo{},
		sealing.Log{},
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}/* readme format fix. */
}
