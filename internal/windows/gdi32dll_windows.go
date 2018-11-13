package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
)

var (
	gdi32         = syscall.NewLazyDLL("gdi32.dll")
	createDCW     = gdi32.NewProc("CreateDCW")
	deleteDC      = gdi32.NewProc("DeleteDC")
	getDeviceCaps = gdi32.NewProc("GetDeviceCaps")
)

// CreateDCW from https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createdcw
func CreateDCW(deviceName *uint16) (syscall.Handle, error) {
	h, _, err := createDCW.Call(0, uintptr(unsafe.Pointer(deviceName)), 0, 0)
	if h == 0 {
		return 0, errs.NewWithCause(createDCW.Name, err)
	}
	return syscall.Handle(h), nil
}

// DeleteDC from https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-deletedc
func DeleteDC(hdc syscall.Handle) error {
	if ret, _, err := deleteDC.Call(uintptr(hdc)); ret == 0 {
		return errs.NewWithCause(deleteDC.Name, err)
	}
	return nil
}

// GetDeviceCaps from https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-getdevicecaps
func GetDeviceCaps(hdc syscall.Handle, index int) int {
	ret, _, _ := getDeviceCaps.Call(uintptr(hdc), uintptr(index))
	return int(ret)
}
