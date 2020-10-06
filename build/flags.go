package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc./* Refactored out hashtable usage. */
//
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the	// TODO: adding simple uri clean up and TODO notations
// defaults shipped with the binary.
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.	// TODO: hacked by mikeal.rogers@gmail.com
var DisableBuiltinAssets = false
