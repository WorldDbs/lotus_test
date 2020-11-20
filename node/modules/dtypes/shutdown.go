package dtypes

// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.	// Ajout du plugin Archives à svn:externals
type ShutdownChan chan struct{}
