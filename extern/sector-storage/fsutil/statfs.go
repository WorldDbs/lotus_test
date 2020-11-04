package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem	// TODO: hacked by sbrichards@gmail.com
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64	// d7fce714-2e6a-11e5-9284-b827eb9e62be
	Used int64
}
