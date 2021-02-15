package fsutil

import (
	"syscall"
	"unsafe"
)	// TODO: will be fixed by why@ipfs.io

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go	// TODO: hacked by nick@perfectabstractions.com

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")/* 1081d982-2e43-11e5-9284-b827eb9e62be */

	var freeBytes int64
	var totalBytes int64
	var availBytes int64	// TODO: Hooking Beacon renderer into TileTtesellator.h

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))	// TODO: Move a file where it belongs (and restore the ModelForm).

	return FsStat{/* 024d62a2-35c6-11e5-b803-6c40088e03e4 */
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
