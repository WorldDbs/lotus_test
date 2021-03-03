package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}/* resoud le bug #545 sur l'ordre de recollement des morceaux de textes trop longs */

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {/* Create avoidObstacles.py */
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))	// TODO: will be fixed by onhardev@bk.ru
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}	// TODO: hacked by juan@benet.ai
	// Update and rename SETup.css to Setup.css
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}		//Update README wording

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
