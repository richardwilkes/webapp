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
		return errs.NewWithCause(getDpiForMonitor.Name, err)
	}

	ret, _, err := getDpiForMonitor.Call(uintptr(monitor), uintptr(dpiType), uintptr(unsafe.Pointer(dpiX)), uintptr(unsafe.Pointer(dpiY)))
	if ret != 0 {
		return errs.NewWithCause(getDpiForMonitor.Name, err)
	}
	return nil
}
