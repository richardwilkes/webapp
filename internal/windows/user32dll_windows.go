package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/xmath/geom"
)

var (
	user32                        = syscall.NewLazyDLL("user32.dll")
	createMenu                    = user32.NewProc("CreateMenu")
	createPopupMenu               = user32.NewProc("CreatePopupMenu")
	createWindowExW               = user32.NewProc("CreateWindowExW")
	defWindowProcW                = user32.NewProc("DefWindowProcW")
	destroyMenu                   = user32.NewProc("DestroyMenu")
	destroyWindow                 = user32.NewProc("DestroyWindow")
	enumDisplayDevicesW           = user32.NewProc("EnumDisplayDevicesW")
	enumDisplayMonitors           = user32.NewProc("EnumDisplayMonitors")
	enumDisplaySettingsExW        = user32.NewProc("EnumDisplaySettingsExW")
	enumWindows                   = user32.NewProc("EnumWindows")
	getClientRect                 = user32.NewProc("GetClientRect")
	getMenu                       = user32.NewProc("GetMenu")
	getMonitorInfoW               = user32.NewProc("GetMonitorInfoW")
	getWindowRect                 = user32.NewProc("GetWindowRect")
	loadCursorW                   = user32.NewProc("LoadCursorW")
	moveWindow                    = user32.NewProc("MoveWindow")
	postQuitMessage               = user32.NewProc("PostQuitMessage")
	registerClassExW              = user32.NewProc("RegisterClassExW")
	registerWindowMessageW        = user32.NewProc("RegisterWindowMessageW")
	setActiveWindow               = user32.NewProc("SetActiveWindow")
	setMenu                       = user32.NewProc("SetMenu")
	setMenuItemInfoW              = user32.NewProc("SetMenuItemInfoW")
	setProcessDpiAwarenessContext = user32.NewProc("SetProcessDpiAwarenessContext")
	setWindowPos                  = user32.NewProc("SetWindowPos")
	setWindowTextW                = user32.NewProc("SetWindowTextW")
	showWindow                    = user32.NewProc("ShowWindow")
)

// CreateMenu from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createmenu
func CreateMenu() (syscall.Handle, error) {
	ret, _, err := createMenu.Call()
	if ret == 0 {
		return 0, errs.NewWithCause(createMenu.Name, err)
	}
	return syscall.Handle(ret), nil
}

// CreatePopupMenu from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createpopupmenu
func CreatePopupMenu() (syscall.Handle, error) {
	ret, _, err := createPopupMenu.Call()
	if ret == 0 {
		return 0, errs.NewWithCause(createPopupMenu.Name, err)
	}
	return syscall.Handle(ret), nil
}

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

// DestroyMenu from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-destroymenu
func DestroyMenu(menu syscall.Handle) error {
	h, _, err := destroyMenu.Call(uintptr(menu))
	if h == 0 {
		return errs.NewWithCause(destroyMenu.Name, err)
	}
	return nil
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

// EnumDisplayMonitors from https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-enumdisplaymonitors
func EnumDisplayMonitors(hdc syscall.Handle, clip *RECT, callback func(monitor, dc syscall.Handle, rect, param uintptr) uintptr, data uintptr) error {
	if ret, _, err := enumDisplayMonitors.Call(uintptr(hdc), uintptr(unsafe.Pointer(clip)), syscall.NewCallback(callback), data); ret == 0 {
		return errs.NewWithCause(enumDisplayMonitors.Name, err)
	}
	return nil
}

// EnumDisplaySettingsExW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaysettingsexw
func EnumDisplaySettingsExW(deviceName *uint16, modeNum uint32, flags uint32) (*DEVMODEW, error) {
	var data DEVMODEW
	if ret, _, err := enumDisplaySettingsExW.Call(uintptr(unsafe.Pointer(deviceName)), uintptr(modeNum), uintptr(unsafe.Pointer(&data)), uintptr(flags)); ret == 0 {
		return nil, errs.NewWithCause(enumDisplaySettingsExW.Name, err)
	}
	return &data, nil
}

// EnumWindows from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumwindows
func EnumWindows(callback func(wnd syscall.Handle, data uintptr) uintptr, data uintptr) error {
	if ret, _, err := enumWindows.Call(syscall.NewCallback(callback), data); ret == 0 {
		return errs.NewWithCause(enumWindows.Name, err)
	}
	return nil
}

// GetClientRect from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getclientrect
func GetClientRect(hwnd syscall.Handle) (geom.Rect, error) {
	var bounds geom.Rect
	var rect RECT
	if ret, _, err := getClientRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect))); ret == 0 {
		return bounds, errs.NewWithCause(getClientRect.Name, err)
	}
	bounds.X = float64(rect.Left)
	bounds.Y = float64(rect.Top)
	bounds.Width = float64(rect.Right - rect.Left)
	bounds.Height = float64(rect.Bottom - rect.Top)
	return bounds, nil
}

// GetMenu from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getmenu
func GetMenu(wnd syscall.Handle) syscall.Handle {
	ret, _, _ := getMenu.Call(uintptr(wnd))
	return syscall.Handle(ret)
}

func GetMonitorInfoW(monitor syscall.Handle) (*MONITORINFO, error) {
	var info MONITORINFO
	info.Size = uint32(unsafe.Sizeof(info))
	if ret, _, err := getMonitorInfoW.Call(uintptr(monitor), uintptr(unsafe.Pointer(&info))); ret == 0 {
		return nil, errs.NewWithCause(getMonitorInfoW.Name, err)
	}
	return &info, nil
}

// GetWindowRect from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getwindowrect
func GetWindowRect(hwnd syscall.Handle) (geom.Rect, error) {
	var bounds geom.Rect
	var rect RECT
	if ret, _, err := getWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&rect))); ret == 0 {
		return bounds, errs.NewWithCause(getWindowRect.Name, err)
	}
	bounds.X = float64(rect.Left)
	bounds.Y = float64(rect.Top)
	bounds.Width = float64(rect.Right - rect.Left)
	bounds.Height = float64(rect.Bottom - rect.Top)
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

// RegisterWindowMessageW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessageW(name string) (uint32, error) {
	str, err := syscall.UTF16PtrFromString(name)
	if err != nil {
		return 0, errs.NewWithCause("Unable to convert name to UTF16", err)
	}
	ret, _, err := registerWindowMessageW.Call(uintptr(unsafe.Pointer(str)))
	if ret == 0 {
		return 0, errs.NewWithCause(registerWindowMessageW.Name, err)
	}
	return uint32(ret), nil
}

// SetActiveWindow from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setactivewindow
func SetActiveWindow(hwnd syscall.Handle) error {
	if ret, _, err := setActiveWindow.Call(uintptr(hwnd)); ret == 0 {
		return errs.NewWithCause(setActiveWindow.Name, err)
	}
	return nil
}

// SetMenu from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setmenu
func SetMenu(wnd, menu syscall.Handle) error {
	if ret, _, err := setMenu.Call(uintptr(wnd), uintptr(menu)); ret == 0 {
		return errs.NewWithCause(setMenu.Name, err)
	}
	return nil
}

// SetMenuItemInfoW from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setmenuiteminfow
func SetMenuItemInfoW(menu syscall.Handle, item uint32, byPosition bool, info *MENUITEMINFOW) error {
	if ret, _, err := setMenuItemInfoW.Call(uintptr(menu)); ret == 0 {
		return errs.NewWithCause(setMenuItemInfoW.Name, err)
	}
	return nil
}

// SetProcessDpiAwarenessContext from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setprocessdpiawarenesscontext
func SetProcessDpiAwarenessContext(value uint32) error {
	if ret, _, err := setProcessDpiAwarenessContext.Call(uintptr(value)); ret == 0 {
		return errs.NewWithCause(setProcessDpiAwarenessContext.Name, err)
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

// ShowWindow from https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-showwindow
func ShowWindow(hwnd syscall.Handle, cmd int) bool {
	ret, _, _ := showWindow.Call(uintptr(hwnd), uintptr(cmd))
	if ret == 0 {
		return false
	}
	return true
}
