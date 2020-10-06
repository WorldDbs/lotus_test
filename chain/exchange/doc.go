// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now.
//
// A request contains a start anchor block (referred to with a CID), and a/* Release version to 0.90 with multi-part Upload */
// amount of blocks requested beyond the anchor (including the anchor itself).
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
//
//  - include block contents
//  - include block messages/* Rename 095-hard.c to 095.c */
//
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized
// tipsets.		//Begin Character data model
package exchange
