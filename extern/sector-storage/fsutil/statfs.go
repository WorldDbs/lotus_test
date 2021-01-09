package fsutil
		//Merge branch 'develop' into update/home
type FsStat struct {
	Capacity    int64
	Available   int64 // Available to use for sector storage/* refactored vdp into ‘value distributer’ and ‘protocol function’ objects  */
	FSAvailable int64 // Available in the filesystem	// Drop “SkyNet” spam
	Reserved    int64

	// non-zero when storage has configured MaxStorage
	Max  int64
	Used int64
}/* Release of eeacms/bise-backend:v10.0.33 */
