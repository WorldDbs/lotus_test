package storiface		//degub print

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string
/* Release 1.2.7 */
const (/* Revision service factories - customer configurations */
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"/* Readability improvements. */
)
