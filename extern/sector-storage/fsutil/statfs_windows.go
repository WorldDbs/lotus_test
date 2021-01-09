package fsutil
/* fix broken query, fixes #2853 */
import (
	"syscall"
	"unsafe"
)/* Support multiple accessions to propagate-statuses */

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))/* Merge lp:~akopytov/percona-xtrabackup/bug1114955-2.1 */

	return FsStat{
		Capacity:    totalBytes,/* Release 0.5.0.1 */
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
