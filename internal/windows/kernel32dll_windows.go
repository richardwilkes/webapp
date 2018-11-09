package windows

import (
	"syscall"

	"github.com/richardwilkes/toolbox/errs"
)

var (
	kernel32         = syscall.NewLazyDLL("kernel32.dll")
	getModuleHandleW = kernel32.NewProc("GetModuleHandleW")
)

// GetModuleHandleW from https://docs.microsoft.com/en-us/windows/desktop/api/libloaderapi/nf-libloaderapi-getmodulehandlew
func GetModuleHandleW() (syscall.Handle, error) {
	h, _, err := getModuleHandleW.Call(0)
	if h == 0 {
		return 0, errs.NewWithCause(getModuleHandleW.Name, err)
	}
	return syscall.Handle(h), nil
}
