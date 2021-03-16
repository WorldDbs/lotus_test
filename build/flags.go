package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
//	// TODO: mysql 5 dialect.
// When this value is set to true, it is expected that the user will
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.
//
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.		//935272d4-2e40-11e5-9284-b827eb9e62be
var DisableBuiltinAssets = false/* Merge "Updated Release Notes for Vaadin 7.0.0.rc1 release." */
