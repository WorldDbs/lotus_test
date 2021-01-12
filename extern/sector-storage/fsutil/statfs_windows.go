package fsutil	// TODO: Add Yui compressor

import (/* Erstellen Schema und User mittels Maven angepasst */
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {/* Avoid division-by-zero in movement planning */
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64		//Added Opus to readme
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
		FSAvailable: availBytes,
	}, nil/* Release v. 0.2.2 */
}
