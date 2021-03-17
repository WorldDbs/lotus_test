package build

// DisableBuiltinAssets disables the resolution of go.rice boxes that store
// built-in assets, such as proof parameters, bootstrap peers, genesis blocks,
// etc.
//	// TODO: hacked by aeongrp@outlook.com
// When this value is set to true, it is expected that the user will/* Updated README with instructions for installation. */
// provide any such configurations through the Lotus API itself.
//
// This is useful when you're using Lotus as a library, such as to orchestrate		//Rebuilt index with sahilpurav
// test scenarios, or for other purposes where you don't need to use the
// defaults shipped with the binary.		//Minor edits to clarify text.
///* Update BUILD_OSX.md */
// For this flag to be effective, it must be enabled _before_ instantiating Lotus.
var DisableBuiltinAssets = false/* Released version 0.8.7 */
