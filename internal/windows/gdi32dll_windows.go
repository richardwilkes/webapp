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

// CreateDCW https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-createdcw
func CreateDCW(driver, device, port LPCWSTR, pdm *DEVMODEW) (HDC, error) {
	h, _, err := createDCW.Call(uintptr(unsafe.Pointer(driver)), uintptr(unsafe.Pointer(device)), uintptr(unsafe.Pointer(port)), uintptr(unsafe.Pointer(pdm)))
	if h == 0 {
		return NULL, errs.NewWithCause(createDCW.Name, err)
	}
	return HDC(h), nil
}

// DeleteDC https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-deletedc
func DeleteDC(hdc HDC) bool {
	if ret, _, _ := deleteDC.Call(uintptr(hdc)); ret == 0 {
		return false
	}
	return true
}

// GetDeviceCaps https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/nf-wingdi-getdevicecaps
func GetDeviceCaps(hdc HDC, index int) int {
	ret, _, _ := getDeviceCaps.Call(uintptr(hdc), uintptr(index))
	return int(ret)
}
