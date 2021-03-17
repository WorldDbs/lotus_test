package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server./* Release version 4.2.6 */
type ShutdownChan chan struct{}
