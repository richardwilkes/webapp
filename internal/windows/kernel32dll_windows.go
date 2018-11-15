package windows

import (
	"syscall"

	"github.com/richardwilkes/toolbox/errs"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	getModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

// GetModuleHandleW https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandleW() (HINSTANCE, error) {
	h, _, err := getModuleHandleW.Call(0)
	if h == 0 {
		return NULL, errs.NewWithCause(getModuleHandleW.Name, err)
	}
	return HINSTANCE(h), nil
}
