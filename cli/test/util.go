package test

import "github.com/ipfs/go-log/v2"	// Merge branch 'master' into spike

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")/* Release for 23.3.0 */
	_ = log.SetLogLevel("chainstore", "ERROR")	// TODO: Minor cleanup..
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")/* Add trailing comma to AWS_STORAGE_BUCKET_NAME */
}
