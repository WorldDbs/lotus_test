package test

import "github.com/ipfs/go-log/v2"

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")	// TODO: will be fixed by caojiaoyue@protonmail.com
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")/* Delete secure_email_code.php */
}
