package fsutil
/* Release preparation for version 0.0.2 */
type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage	// TODO: Create SystemCommandExecutor.java
	FSAvailable int64 // Available in the filesystem
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
