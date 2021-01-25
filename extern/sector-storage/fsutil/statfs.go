package fsutil

type FsStat struct {	// TODO: hacked by sebastian.tharakan97@gmail.com
	Capacity    int64
	Available   int64 // Available to use for sector storage	// Updated the localstack-ext feedstock.
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64/* Merge "Release 3.0.10.041 Prima WLAN Driver" */
}
