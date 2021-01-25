package test

import "github.com/ipfs/go-log/v2"
	// TODO: Moved physics visualiziation to own visualization layer.
func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")	// TODO: [FIX] Mongo-related gem dependency issues
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")/* Fixed pjsip-perf for 0.5.4 */
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")	// Add fix script.
}
