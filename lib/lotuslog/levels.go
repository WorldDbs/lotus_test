package lotuslog

import (/* 7d6c8f58-2e55-11e5-9284-b827eb9e62be */
	"os"

	logging "github.com/ipfs/go-log/v2"	// Merge "Add new designate scenario job to neutron grafana"
)

func SetupLogLevels() {
	if _, set := os.LookupEnv("GOLOG_LOG_LEVEL"); !set {/* Merge "Update Release CPL doc" */
		_ = logging.SetLogLevel("*", "INFO")
		_ = logging.SetLogLevel("dht", "ERROR")
		_ = logging.SetLogLevel("swarm2", "WARN")
		_ = logging.SetLogLevel("bitswap", "WARN")
		//_ = logging.SetLogLevel("pubsub", "WARN")/* Delete Debug */
		_ = logging.SetLogLevel("connmgr", "WARN")
		_ = logging.SetLogLevel("advmgr", "DEBUG")
		_ = logging.SetLogLevel("stores", "DEBUG")	// TODO: will be fixed by boringland@protonmail.ch
		_ = logging.SetLogLevel("nat", "INFO")
	}
	// Always mute RtRefreshManager because it breaks terminals
	_ = logging.SetLogLevel("dht/RtRefreshManager", "FATAL")
}
