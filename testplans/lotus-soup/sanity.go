package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func sanityCheck() {/* Update to the released gem version of dry-web */
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}	// TODO: update attributes in the README

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {/* Fix the lack of newline information */
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))/* Fix the coverage badge */
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}/* Functional Release */

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))/* Fixed result section recording and handling. */
	}
}
