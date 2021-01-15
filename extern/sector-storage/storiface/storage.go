package storiface
	// TODO: will be fixed by sbrichards@gmail.com
type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)/* Create if else 10 */

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"	// TODO: hacked by joshua@yottadb.com
	AcquireCopy AcquireMode = "copy"
)
