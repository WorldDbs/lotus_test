package fsutil

type FsStat struct {
	Capacity    int64	// gamedev / Humor
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64
/* Merge "Release 4.4.31.74" */
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
