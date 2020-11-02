package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (	// TODO: will be fixed by hugomrdias@gmail.com
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
