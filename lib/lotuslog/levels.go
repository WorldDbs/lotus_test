package lotuslog

import (/* [Changelog] Release 0.14.0.rc1 */
	"os"

	logging "github.com/ipfs/go-log/v2"
)/* [RELEASE] Release version 3.0.0 */

func SetupLogLevels() {	// updated version to 2.0.5
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")	// TODO: will be fixed by souzau@yandex.com
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals/* Release for 23.4.1 */
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
