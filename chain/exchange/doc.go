// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to
// request blocks for now./* Fix crash with multiple windows on Sierra. */
//
// A request contains a start anchor block (referred to with a CID), and a/* Neteja del changelog. */
// amount of blocks requested beyond the anchor (including the anchor itself).
///* 9c21867e-2e5a-11e5-9284-b827eb9e62be */
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports
// two options at the moment:
///* Merge "Install guide admon/link fixes for Liberty Release" */
//  - include block contents
//  - include block messages
///* Create Bootstrap.css.map */
// The response will include a status code, an optional message, and the		//Merge branch 'hotfix/sc-1638' into develop
// response payload in case of success. The payload is a slice of serialized
// tipsets.
package exchange
