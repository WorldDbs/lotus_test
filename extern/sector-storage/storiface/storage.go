package storiface
		//Updated pom description.
type PathType string/* Released reLexer.js v0.1.1 */

const (
	PathStorage PathType = "storage"	// TODO: hacked by yuvalalaluf@gmail.com
	PathSealing PathType = "sealing"
)

type AcquireMode string		//clearer pause and stop documentation

const (
	AcquireMove AcquireMode = "move"	// Make sur we always return an array
	AcquireCopy AcquireMode = "copy"
)
