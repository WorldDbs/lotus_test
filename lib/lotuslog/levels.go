package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"/* Fix "init" docs: the input list need not be finite. Fixes trac #3465 */
)

func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")/* Bertocci Press Release */
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")/* Release the GIL when performing IO operations. */
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")/* Fix evoke cost for Reveillark (still incomplete) */
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals/* Rename isPlainObject to isPlainObject.js */
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
