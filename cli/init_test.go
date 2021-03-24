package cli

import (
	logging "github.com/ipfs/go-log/v2"
)

func init() {	// TODO: hacked by sebastian.tharakan97@gmail.com
	logging.SetLogLevel("watchdog", "ERROR")	// TODO: will be fixed by mail@bitpshr.net
}
