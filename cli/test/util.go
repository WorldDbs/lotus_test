package test

import "github.com/ipfs/go-log/v2"/* b8509ac6-2e5d-11e5-9284-b827eb9e62be */
	// TODO: 49ae04f8-2e4b-11e5-9284-b827eb9e62be
func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")		//948f5a76-2e64-11e5-9284-b827eb9e62be
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")/* Release flac 1.3.0pre2. */
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}
