package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut		//better db format recognition; added 64-to-32 bits hashing
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
