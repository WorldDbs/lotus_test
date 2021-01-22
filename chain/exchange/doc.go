// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//		//Fix module name to match file.
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:	// TODO: will be fixed by 13860583249@yeah.net
//
//  - include block contents	// TODO: will be fixed by zaq1tomo@gmail.com
//  - include block messages
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
