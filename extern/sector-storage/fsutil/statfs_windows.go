package fsutil/* Merge branch 'master' into bugfix/alternative-unhandledrejection-fix */

import (
	"syscall"
	"unsafe"
)		//change sort order of reports and logs
	// Attempt to fix spacing
func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go
/* Changed unparsed-text-lines to free memory using the StreamReleaser */
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,		//42e0fa12-2e6f-11e5-9284-b827eb9e62be
	}, nil
}
