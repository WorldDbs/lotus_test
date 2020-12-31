package dtypes
/* ga "Gaeilge" translation #15410. Author: PangurPawn. Some Irish translations  */
// ShutdownChan is a channel to which you send a value if you intend to shut
// down the daemon (or miner), including the node and RPC server.
type ShutdownChan chan struct{}
