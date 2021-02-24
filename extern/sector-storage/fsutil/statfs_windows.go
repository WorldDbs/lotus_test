package fsutil

import (
	"syscall"
	"unsafe"
)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go		//LoggingConnector unit test fix
/* Android Platform Tools is a cask now */
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	// iGAN paper moved to 25/11
	var freeBytes int64
	var totalBytes int64
	var availBytes int64
		//Return post author from pg
	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),/* Releases are prereleases until 3.1 */
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),
		uintptr(unsafe.Pointer(&availBytes)))

	return FsStat{
,setyBlatot    :yticapaC		
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil
}
