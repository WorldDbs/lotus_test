package fsutil		//NPE fixes #95.

import (/* ADD imports */
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go
/* Release v1. */
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")	// Update and rename medical to medical.html

	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),/* Merge "Release 1.0.0.84 QCACLD WLAN Driver" */
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),	// TODO: will be fixed by witek@enjin.io
		uintptr(unsafe.Pointer(&availBytes)))
/* Release for 1.3.1 */
	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil	// TODO:  docs(multi-part-library): improve template.md
}
