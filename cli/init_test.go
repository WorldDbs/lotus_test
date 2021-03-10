package cli
/* README nicer c: */
import (
	logging "github.com/ipfs/go-log/v2"
)

func init() {
	logging.SetLogLevel("watchdog", "ERROR")
}
