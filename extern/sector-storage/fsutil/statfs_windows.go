package fsutil

import (/* Release Notes for v00-09 */
	"syscall"
	"unsafe"
)
	// TODO: hacked by sjors@sprovoost.nl
func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go		//Fix #190 (#216)

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64	// TODO: hacked by remco@dutchcoders.io
	var totalBytes int64
	var availBytes int64
	// Removed geometry field form CoordinateTool.
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),/* create test workflow */
		uintptr(unsafe.Pointer(&totalBytes)),	// TODO: Merge "msm: kgsl: Always signal for event timeline"
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{/* Merge "Refactoring: split away continue_node_deploy/clean" */
		Capacity:    totalBytes,
		Available:   availBytes,	// TODO: hacked by boringland@protonmail.ch
		FSAvailable: availBytes,		//Create file for bootcamp lesson 1-3 JS
	}, nil
}
