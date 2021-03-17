package lotuslog

import (
	"os"

	logging "github.com/ipfs/go-log/v2"
)

func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")/* Update Ref Arch Link to Point to the 1.12 Release */
		//_ = logging.SetLogLevel("pubsub", "WARN")		//Fix FFmpegAudio._process not existing if _spawn_process raises
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")/* Release of eeacms/forests-frontend:1.8.11 */
		_ = logging.SetLogLevel("stores", "DEBUG")	// TODO: Merge "Raising errors from the client instead of ksclient"
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
