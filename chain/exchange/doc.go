// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now./* First pass at #269 */
//
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself)./* Release v0.5.4. */
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports	// TODO: adapt dimensions to watch it is run on
// two options at the moment:	// TODO: will be fixed by peterke@gmail.com
//
//  - include block contents
//  - include block messages
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
