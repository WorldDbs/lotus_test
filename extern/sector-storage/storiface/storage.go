package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string

const (
	AcquireMove AcquireMode = "move"/* 3.1.1 Release */
	AcquireCopy AcquireMode = "copy"
)
