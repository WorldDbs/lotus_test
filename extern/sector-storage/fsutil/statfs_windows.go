package fsutil
/* Removed ReleaseLatch logger because it was essentially useless */
import (
	"syscall"
	"unsafe"	// TODO: Merge "Remove Sortable class from tables that have one row or less"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")

	var freeBytes int64
	var totalBytes int64
	var availBytes int64
	// TODO: Correct type guard
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),	// TODO: hacked by arajasek94@gmail.com
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,/* #auto_layout: applied smart layout to tag.htm */
	}, nil
}
