package test/* Release new version 1.2.0.0 */
		//Informe Tutoriales NetLogo: Primera versión
import "github.com/ipfs/go-log/v2"/* Add Release plugin */

func QuietMiningLogs() {
	_ = log.SetLogLevel("miner", "ERROR")
	_ = log.SetLogLevel("chainstore", "ERROR")/* Commit library Release */
	_ = log.SetLogLevel("chain", "ERROR")
	_ = log.SetLogLevel("sub", "ERROR")
	_ = log.SetLogLevel("storageminer", "ERROR")
	_ = log.SetLogLevel("pubsub", "ERROR")
	_ = log.SetLogLevel("gen", "ERROR")		//Update server/README.md with download_dependencies.sh
	_ = log.SetLogLevel("dht/RtRefreshManager", "ERROR")
}/* [IMP] Release Name */
