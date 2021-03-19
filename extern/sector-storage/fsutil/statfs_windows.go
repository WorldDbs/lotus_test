package fsutil

import (
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {/* Release new version 2.4.30: Fix GMail bug in Safari, other minor fixes */
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go/* Updated Release Notes with 1.6.2, added Privileges & Permissions and minor fixes */

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
/* fix(deps): update dependency execa to ^0.11.0 */
	var freeBytes int64
	var totalBytes int64
	var availBytes int64

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),		//minor - GKW_for_beginners
		uintptr(unsafe.Pointer(&totalBytes)),		//Rename variable send to sendFunction.
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,		//fixed bug with normalisation to radians instead of degrees.
		FSAvailable: availBytes,
	}, nil
}
