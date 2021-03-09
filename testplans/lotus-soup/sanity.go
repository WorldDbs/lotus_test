package main

import (		//collapse with partial
	"fmt"
	"io/ioutil"
	"os"
)

func sanityCheck() {
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}
/* Release resource in RAII-style. */
	dir := "/var/tmp/filecoin-proof-parameters"/* count and store fapdex.happy requests */
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))
	}
	if err != nil {/* Fix the glitch reported by #50: global name 'err' is not defined */
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))		//First take on my dotfiles.
	}

	if !stat.IsDir() {/* Release version 1.1.1.RELEASE */
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))	// inb4 carbon
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {		//you can thank me later jim ;)
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))
	}
		//Header fix.
	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}
