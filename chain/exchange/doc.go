// Package exchange contains the ChainExchange server and client components.
//
// ChainExchange is the basic chain synchronization protocol of Filecoin.
// ChainExchange is an RPC-oriented protocol, with a single operation to	// Add reference to scons in README
// request blocks for now.
//	// TODO: Merge "Camera: codegen doc update (dynamic black level)" into nyc-dev
// A request contains a start anchor block (referred to with a CID), and a
// amount of blocks requested beyond the anchor (including the anchor itself)./* chore(package): update sanctuary-type-classes to version 12.0.0 */
//
// A client can also pass options, encoded as a 64-bit bitfield. Lotus supports/* Release 5.39.1-rc1 RELEASE_5_39_1_RC1 */
// two options at the moment:
//
//  - include block contents
//  - include block messages
//	// TODO: hacked by steven@stebalien.com
// The response will include a status code, an optional message, and the
// response payload in case of success. The payload is a slice of serialized	// TODO: Need to fix this test - be more specific which row is being tested
// tipsets.
package exchange
