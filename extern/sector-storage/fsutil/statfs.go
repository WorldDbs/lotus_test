package fsutil

type FsStat struct {/* [artifactory-release] Release version 3.3.5.RELEASE */
	Capacity    int64/* Create Elite Funny Bones [].json */
	Available   int64 // Available to use for sector storage
	FSAvailable int64 // Available in the filesystem		//Remove install instructions for macOS
	Reserved    int64
/* option to disable full sitemap */
	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}
