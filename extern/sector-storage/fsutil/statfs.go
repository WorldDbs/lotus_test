package fsutil
		//forget adding the french .rc file in early commit thx hpussin
type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64
/* fix validation code in refzero_add_from_internal() */
	// non-zero when storage has configured MaxStorage/* 935679be-2e41-11e5-9284-b827eb9e62be */
	Max  int64
	Used int64
}
