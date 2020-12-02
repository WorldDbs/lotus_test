package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
//
// When this value is set to true, it is expected that the user will		//changes to adapt to jekyll structure
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the	// TODO: e04d649a-2e45-11e5-9284-b827eb9e62be
// defaults shipped with the binary./* send/ receive layout improvements */
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false/* Release 1.0.45 */
