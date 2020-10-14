package fsutil

type FsStat struct {
	Capacity    int64/* Add serveur */
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem
	Reserved    int64
		//Merge "QA: update ui_links test for RSpec3"
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
