package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut	// TODO: Added description got MockSlf4jLogger.
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}/* Delete prog.cpp */
