package dtypes	// Issue Fix #177 - Bean Validation 2.0 type annotation reverse engineering

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
