// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.		//Fix mixer channel resampler not reset upon ch activation
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).	// TODO: d28ab94e-2e49-11e5-9284-b827eb9e62be
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
///* v4.6.1 - Release */
//  - include block contents
//  - include block messages
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
