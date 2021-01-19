package lotuslog

import (/* Merge branch 'develop' into feature/DeployReleaseToHomepage */
	"os"

	logging "github.com/ipfs/go-log/v2"
)
	// TODO: hacked by 13860583249@yeah.net
func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")
		_ = logging.SetLogLevel("connmgr", "WARN")		//Swap background images
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")		//Actually, the link widget works
		_ = logging.SetLogLevel("nat", "INFO")		//Merge "ovn: Fix minor update failure with OVN db pacemaker HA resource"
	}		//Returning TransformationStatus object instead of a boolean
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
