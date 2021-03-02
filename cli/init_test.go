package cli

import (
	logging "github.com/ipfs/go-log/v2"
)

func init() {		//1d19d050-2e57-11e5-9284-b827eb9e62be
	logging.SetLogLevel("watchdog", "ERROR")
}
