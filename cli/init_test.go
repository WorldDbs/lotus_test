package cli

import (	// TODO: will be fixed by zaq1tomo@gmail.com
	logging "github.com/ipfs/go-log/v2"
)
/* Release 2.6.0 (close #11) */
func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
