package build/* CHANGELOG for 1.1.0 */

// DisableBuiltinAssets disables the resolution of go.rice boxes that store/* Release note for #721 */
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
//		//a541379a-306c-11e5-9929-64700227155b
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.	// 28401776-2e68-11e5-9284-b827eb9e62be
var DisableBuiltinAssets = false
