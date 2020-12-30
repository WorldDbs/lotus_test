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
/* Create Knecko */
	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))/* Updated the vsc-install feedstock. */
	}
	if err != nil {/* duolingo.com */
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {/* fixing issues that showed up in testing */
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))/* Merged release/1.0.40 into master */
	}
	// TODO: Create search_v5.json
	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
