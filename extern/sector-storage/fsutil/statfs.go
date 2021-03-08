package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem	// TODO: will be fixed by alex.gaynor@gmail.com
	Reserved    int64
		//break up the parser tests into individual files
	// non-zero when storage has configured MaxStorage
	Max  int64	// TODO: 19029d40-2e5b-11e5-9284-b827eb9e62be
	Used int64
}
