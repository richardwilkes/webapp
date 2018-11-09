package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	createWindowExW  = user32.NewProc("CreateWindowExW")
	defWindowProcW   = user32.NewProc("DefWindowProcW")
	destroyWindow    = user32.NewProc("DestroyWindow")
	loadCursorW      = user32.NewProc("LoadCursorW")
	postQuitMessage  = user32.NewProc("PostQuitMessage")
	registerClassExW = user32.NewProc("RegisterClassExW")
)

// CreateWindowExW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createwindowexw
func CreateWindowExW(exStyle uint32, className, windowName string, style uint32, x, y, width, height int32, parent, menu, instance syscall.Handle) (syscall.Handle, error) {
	cnstr, err := syscall.UTF16PtrFromString(className)
	if err != nil {
		return 0, errs.NewWithCause("Unable to convert className to UTF16", err)
	}
	wnstr, err := syscall.UTF16PtrFromString(windowName)
	if err != nil {
		return 0, errs.NewWithCause("Unable to convert windowName to UTF16", err)
	}
	h, _, err := createWindowExW.Call(uintptr(exStyle), uintptr(unsafe.Pointer(cnstr)), uintptr(unsafe.Pointer(wnstr)), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(parent), uintptr(menu), uintptr(instance), uintptr(0))
	if h == 0 {
		return 0, errs.NewWithCause(createWindowExW.Name, err)
	}
	return syscall.Handle(h), nil
}

// DefWindowProcW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-defwindowprocw
func DefWindowProcW(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	ret, _, _ := defWindowProcW.Call(uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam))
	return uintptr(ret)
}

// DestroyWindow from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-destroywindow
func DestroyWindow(hwnd syscall.Handle) error {
	h, _, err := destroyWindow.Call(uintptr(hwnd))
	if h == 0 {
		return errs.NewWithCause(destroyWindow.Name, err)
	}
	return nil
}

// LoadCursorW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursorw
func LoadCursorW(cursorName uint32) (syscall.Handle, error) {
	h, _, err := loadCursorW.Call(0, uintptr(uint16(cursorName)))
	if h == 0 {
		return 0, errs.NewWithCause(loadCursorW.Name, err)
	}
	return syscall.Handle(h), nil
}

// PostQuitMessage from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	postQuitMessage.Call(uintptr(exitCode))
}

// RegisterClassExW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-registerclassexw
func RegisterClassExW(w *WNDCLASSEXW) (uint16, error) {
	h, _, err := registerClassExW.Call(uintptr(unsafe.Pointer(w)))
	if h == 0 {
		return 0, errs.NewWithCause(registerClassExW.Name, err)
	}
	return uint16(h), nil
}
