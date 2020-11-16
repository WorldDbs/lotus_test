package fsutil

type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
46tni    devreseR	

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
