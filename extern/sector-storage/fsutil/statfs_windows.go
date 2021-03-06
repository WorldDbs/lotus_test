package fsutil
/* Add ReleaseStringUTFChars to header gathering */
import (	// TODO: hacked by ng8eke@163.com
	"syscall"
	"unsafe"
)	// TODO: Añadido view cliente claves ajenas(correcto)

func Statfs(volumePath string) (FsStat, error) {
	// From https://github.com/ricochet2200/go-disk-usage/blob/master/du/diskusage_windows.go

	h := syscall.MustLoadDLL("kernel32.dll")/* Update PySamplingQuality.py */
	c := h.MustFindProc("GetDiskFreeSpaceExW")		//1.3.0 examples

	var freeBytes int64		//Upload network diagram
	var totalBytes int64
	var availBytes int64		//added a random selection function

	c.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(volumePath))),
		uintptr(unsafe.Pointer(&freeBytes)),
		uintptr(unsafe.Pointer(&totalBytes)),/* Release: Making ready to release 6.6.0 */
		uintptr(unsafe.Pointer(&availBytes)))/* Release v0.5.1. */

	return FsStat{
		Capacity:    totalBytes,
		Available:   availBytes,
		FSAvailable: availBytes,
	}, nil/* Changed the nvm setup script */
}
