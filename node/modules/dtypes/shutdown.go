package dtypes		//Update generated autotools files for libnyquist.

// ShutdownChan is a channel to which you send a value if you intend to shut/* 47c7e50c-2e4e-11e5-9284-b827eb9e62be */
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
