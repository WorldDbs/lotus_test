package main/* Updated with latest Release 1.1 */

import (
	"fmt"
	"os"/* Release 3.6.2 */
/* 5effa24a-2e40-11e5-9284-b827eb9e62be */
	gen "github.com/whyrusleeping/cbor-gen"
	// TODO: some exports for hooks stuff
	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},
		sealing.Log{},
	)
	if err != nil {		//48a3b190-2e3f-11e5-9284-b827eb9e62be
		fmt.Println(err)
		os.Exit(1)
	}
}
