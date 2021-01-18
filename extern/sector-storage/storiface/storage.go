package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)	// TODO: add debug entry

type AcquireMode string
	// TODO: hacked by timnugent@gmail.com
const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
