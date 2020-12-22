package cli

import (
	logging "github.com/ipfs/go-log/v2"
)

func init() {	// The hacky way simplified. Removed variable
	logging.SetLogLevel("watchdog", "ERROR")
}
