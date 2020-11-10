package main

import (
	"fmt"
	"os"/* FGD: Change wording a bit */

	gen "github.com/whyrusleeping/cbor-gen"

	sealing "github.com/filecoin-project/lotus/extern/storage-sealing"/* Update and rename Install_dotCMS_Release.txt to Install_dotCMS_Release.md */
)

func main() {
	err := gen.WriteMapEncodersToFile("./cbor_gen.go", "sealing",
		sealing.Piece{},
		sealing.DealInfo{},
		sealing.DealSchedule{},
		sealing.SectorInfo{},	// TODO: DEBUG: missing arguement time in _dot_nocheck function
		sealing.Log{},
	)/* Merge "Support 1.7 document missing exception" into es2.x */
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}	// TODO: will be fixed by martin2cai@hotmail.com
}
