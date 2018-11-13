package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/xmath/geom"
)

var (
	user32                 = syscall.NewLazyDLL("user32.dll")
	createWindowExW        = user32.NewProc("CreateWindowExW")
	defWindowProcW         = user32.NewProc("DefWindowProcW")
	destroyWindow          = user32.NewProc("DestroyWindow")
	enumDisplayDevicesW    = user32.NewProc("EnumDisplayDevicesW")
	enumDisplaySettingsExW = user32.NewProc("EnumDisplaySettingsExW")
	getClientRect          = user32.NewProc("GetClientRect")
	getWindowRect          = user32.NewProc("GetWindowRect")
	loadCursorW            = user32.NewProc("LoadCursorW")
	moveWindow             = user32.NewProc("MoveWindow")
	postQuitMessage        = user32.NewProc("PostQuitMessage")
	registerClassExW       = user32.NewProc("RegisterClassExW")
	setActiveWindow        = user32.NewProc("SetActiveWindow")
	setWindowPos           = user32.NewProc("SetWindowPos")
	setWindowTextW         = user32.NewProc("SetWindowTextW")
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

// EnumDisplayDevicesW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevicesW(devNum uint32, flags uint32) (*DISPLAY_DEVICEW, error) {
	var data DISPLAY_DEVICEW
	data.Size = uint32(unsafe.Sizeof(data))
	ret, _, err := enumDisplayDevicesW.Call(0, uintptr(devNum), uintptr(unsafe.Pointer(&data)), uintptr(flags))
	if ret == 0 {
		return nil, errs.NewWithCause(enumDisplayDevicesW.Name, err)
	}
	return &data, nil
}

// EnumDisplaySettingsExW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaysettingsexw
func EnumDisplaySettingsExW(deviceName *uint16, modeNum uint32, flags uint32) (*DEVMODE, error) {
	var data DEVMODE
	if ret, _, err := enumDisplaySettingsExW.Call(uintptr(unsafe.Pointer(deviceName)), uintptr(modeNum), uintptr(unsafe.Pointer(&data)), uintptr(flags)); ret == 0 {
		return nil, errs.NewWithCause(enumDisplaySettingsExW.Name, err)
	}
	return &data, nil
}

// GetClientRect from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getclientrect
func GetClientRect(hwnd syscall.Handle) (geom.Rect, error) {
	var bounds geom.Rect
	var rect [4]int32
	if ret, _, err := getClientRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect[0]))); ret == 0 {
		return bounds, errs.NewWithCause(getClientRect.Name, err)
	}
	bounds.X = float64(rect[0])
	bounds.Y = float64(rect[1])
	bounds.Width = float64(rect[2] - rect[0])
	bounds.Height = float64(rect[3] - rect[1])
	return bounds, nil
}

// GetWindowRect from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getwindowrect
func GetWindowRect(hwnd syscall.Handle) (geom.Rect, error) {
	var bounds geom.Rect
	var rect [4]int32
	if ret, _, err := getWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect[0]))); ret == 0 {
		return bounds, errs.NewWithCause(getWindowRect.Name, err)
	}
	bounds.X = float64(rect[0])
	bounds.Y = float64(rect[1])
	bounds.Width = float64(rect[2] - rect[0])
	bounds.Height = float64(rect[3] - rect[1])
	return bounds, nil
}

// LoadCursorW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursorw
func LoadCursorW(cursorName uint32) (syscall.Handle, error) {
	h, _, err := loadCursorW.Call(0, uintptr(uint16(cursorName)))
	if h == 0 {
		return 0, errs.NewWithCause(loadCursorW.Name, err)
	}
	return syscall.Handle(h), nil
}

// MoveWindow from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-movewindow
func MoveWindow(hwnd syscall.Handle, x, y, width, height int32, repaint bool) error {
	r := 0
	if repaint {
		r = 1
	}
	if ret, _, err := moveWindow.Call(uintptr(hwnd), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(r)); ret == 0 {
		return errs.NewWithCause(moveWindow.Name, err)
	}
	return nil
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

// SetActiveWindow from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setactivewindow
func SetActiveWindow(hwnd syscall.Handle) error {
	if ret, _, err := setActiveWindow.Call(uintptr(hwnd)); ret == 0 {
		return errs.NewWithCause(setActiveWindow.Name, err)
	}
	return nil
}

// SetWindowPos from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowpos
func SetWindowPos(hwnd, hwndInsertAfter syscall.Handle, x, y, width, height int32, flags uint32) error {
	if ret, _, err := setWindowPos.Call(uintptr(hwnd), uintptr(hwndInsertAfter), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(flags)); ret == 0 {
		return errs.NewWithCause(setWindowPos.Name, err)
	}
	return nil
}

// SetWindowTextW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowtextw
func SetWindowTextW(hwnd syscall.Handle, title string) error {
	str, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		return errs.NewWithCause("Unable to convert title to UTF16", err)
	}
	ret, _, err := setWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(str)))
	if ret == 0 {
		return errs.NewWithCause(setWindowTextW.Name, err)
	}
	return nil
}
