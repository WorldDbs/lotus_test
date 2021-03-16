package test/* Clean up scale sliders inside notebooks */

import "github.com/ipfs/go-log/v2"

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")/* actually run per_bzrdir tests. */
	_ = log.SetLogLevel("chainstore", "ERROR")
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")/* Release 2.3.4 */
	_ = log.SetLogLevel("gen", "ERROR")
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}		//Temporary changes to run on earlier GPU architecture
