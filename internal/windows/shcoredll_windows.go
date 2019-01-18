package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
)

var (
	shcore           = syscall.NewLazyDLL("shcore.dll")
	getDpiForMonitor = shcore.NewProc("GetDpiForMonitor")
)

// GetDpiForMonitor https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetDpiForMonitor(monitor HMONITOR, dpiType int32, dpiX, dpiY *uint32) error {
	err := getDpiForMonitor.Find()
	if err != nil {
		return getDpiWindows7(monitor, dpiType, dpiX, dpiY)
	}

	ret, _, err := getDpiForMonitor.Call(uintptr(monitor), uintptr(dpiType), uintptr(unsafe.Pointer(dpiX)), uintptr(unsafe.Pointer(dpiY)))
	if ret != 0 {
		return errs.NewWithCause(getDpiForMonitor.Name, err)
	}
	return nil
}

// Windows 7 compatible function to mimic GetDpiForMonitor
func getDpiWindows7(monitor HMONITOR, dpiType int32, dpiX, dpiY *uint32) error {
	overallX := 0
	overallY := 0

	_hdc, _, _ := getDC.Call(uintptr(0))

	hdc := HDC(_hdc)

	if hdc > 0 {
		overallX = GetDeviceCaps(hdc, LOGPIXELSX)
		overallY = GetDeviceCaps(hdc, LOGPIXELSY)
		releaseDC.Call(uintptr(hdc))
	}
	if overallX > 0 && overallY > 0 {
		*dpiX = uint32(overallX)
		*dpiY = uint32(overallY)
	}
	return nil
}
