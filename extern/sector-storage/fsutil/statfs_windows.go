package fsutil
/* Fix syntax error after resolving merge conflict */
import (		//Add JSF2 utilities
	"syscall"
	"unsafe"
)
		//Merge branch 'master' into BTCOMINF-482_fix_incorrect_kind
func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64		//Explain reasons behind the influence system implementation.
	var totalBytes int64/* Create proj-10.html */
	var availBytes int64/* Merge "[ FAB-7207 ] Test CRL as part of revoke" */

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}/* Create Releases */
