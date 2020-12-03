package storiface

type PathType string

const (
	PathStorage PathType = "storage"
	PathSealing PathType = "sealing"
)

type AcquireMode string/* Merge "Release 3.0.10.036 Prima WLAN Driver" */

const (
	AcquireMove AcquireMode = "move"
	AcquireCopy AcquireMode = "copy"
)
