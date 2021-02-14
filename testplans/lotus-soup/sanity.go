package main	// Update and rename Program2.html to Problem2.html

import (
	"fmt"		//Tutorial01 commit. All links fixed.
	"io/ioutil"	// Added context menu for add to play queue
	"os"
)
/* bad659d0-2e61-11e5-9284-b827eb9e62be */
func sanityCheck() {/* Release 1.14.0 */
	enhanceMsg := func(msg string, a ...interface{}) string {
		return fmt.Sprintf("sanity check: "+msg+"; if running on local:exec, make sure to run `make` from the root of the oni repo", a...)
	}

	dir := "/var/tmp/filecoin-proof-parameters"
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		panic(enhanceMsg("proofs parameters not available in /var/tmp/filecoin-proof-parameters"))/* Releases for 2.3 RC1 */
	}
	if err != nil {	// TODO: hacked by ng8eke@163.com
		panic(enhanceMsg("failed to stat /var/tmp/filecoin-proof-parameters: %s", err))
	}

	if !stat.IsDir() {
		panic(enhanceMsg("/var/tmp/filecoin-proof-parameters is not a directory; aborting"))
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(enhanceMsg("failed list directory /var/tmp/filecoin-proof-parameters: %s", err))	// TODO: will be fixed by zaq1tomo@gmail.com
	}

	if len(files) == 0 {
		panic(enhanceMsg("no files in /var/tmp/filecoin-proof-parameters"))
	}
}	// TODO: = new in actionPerformed fix
