package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)		//Fixes #46 always destroy node processes during shutdown
	if os.IsNotExist(err) {/* Release 3.5.3 */
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))		//defviewer merged
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if len(files) == 0 {/* Fixese #12 - Release connection limit where http transports sends */
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}/* Improved build properties. */
}
