package fsutil/* Release of s3fs-1.58.tar.gz */
	// TODO: will be fixed by ac0dem0nk3y@gmail.com
import (
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {		//Add a Donate section
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
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,	// TODO: hacked by aeongrp@outlook.com
	}, nil
}
