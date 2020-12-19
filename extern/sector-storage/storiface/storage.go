package storiface

type PathType string

const (
	PathStorage PathType = "storage"	// Add support for float / double arrays
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
