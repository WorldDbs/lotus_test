package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")/* Adding basic CLI support with the Thor gem */
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")		//fix rst syntax
		_ = logging.SetLogLevel("bitswap", "WARN")/* Fixed equipment Ore Dictionary names. Release 1.5.0.1 */
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")/* Fix broken equal signs in README */
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
