package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
)

var (
	user32                        = syscall.NewLazyDLL("user32.dll")
	createMenu                    = user32.NewProc("CreateMenu")
	createPopupMenu               = user32.NewProc("CreatePopupMenu")
	createWindowExW               = user32.NewProc("CreateWindowExW")
	defWindowProcW                = user32.NewProc("DefWindowProcW")
	deleteMenu                    = user32.NewProc("DeleteMenu")
	destroyMenu                   = user32.NewProc("DestroyMenu")
	destroyWindow                 = user32.NewProc("DestroyWindow")
	enumDisplayDevicesW           = user32.NewProc("EnumDisplayDevicesW")
	enumDisplayMonitors           = user32.NewProc("EnumDisplayMonitors")
	enumDisplaySettingsExW        = user32.NewProc("EnumDisplaySettingsExW")
	enumWindows                   = user32.NewProc("EnumWindows")
	getClientRect                 = user32.NewProc("GetClientRect")
	getMenu                       = user32.NewProc("GetMenu")
	getMenuItemCount              = user32.NewProc("GetMenuItemCount")
	getMenuItemInfoW              = user32.NewProc("GetMenuItemInfoW")
	getMonitorInfoW               = user32.NewProc("GetMonitorInfoW")
	getWindowRect                 = user32.NewProc("GetWindowRect")
	insertMenuItemW               = user32.NewProc("InsertMenuItemW")
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

// CreateMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createmenu
func CreateMenu() (HMENU, error) {
	ret, _, err := createMenu.Call()
	if ret == 0 {
		return NULL, errs.NewWithCause(createMenu.Name, err)
	}
	return HMENU(ret), nil
}

// CreatePopupMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createpopupmenu
func CreatePopupMenu() (HMENU, error) {
	ret, _, err := createPopupMenu.Call()
	if ret == 0 {
		return NULL, errs.NewWithCause(createPopupMenu.Name, err)
	}
	return HMENU(ret), nil
}

// CreateWindowExW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createwindowexw
func CreateWindowExW(exStyle DWORD, className, windowName string, style DWORD, x, y, width, height int32, parent HWND, menu HMENU, instance HINSTANCE, param LPVOID) (HWND, error) {
	cnstr, err := toUTF16PtrOrNilOnEmpty(className)
	if err != nil {
		return NULL, err
	}
	wnstr, err := toUTF16PtrOrNilOnEmpty(windowName)
	if err != nil {
		return NULL, err
	}
	return CreateWindowExW_(exStyle, cnstr, wnstr, style, x, y, width, height, parent, menu, instance, param)
}

// CreateWindowExW_ https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-createwindowexw
func CreateWindowExW_(exStyle DWORD, className, windowName LPCWSTR, style DWORD, x, y, width, height int32, parent HWND, menu HMENU, instance HINSTANCE, param LPVOID) (HWND, error) {
	ret, _, err := createWindowExW.Call(uintptr(exStyle), uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(windowName)), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(parent), uintptr(menu), uintptr(instance), uintptr(param))
	if ret == 0 {
		return NULL, errs.NewWithCause(createWindowExW.Name, err)
	}
	return HWND(ret), nil
}

// DefWindowProcW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-defwindowprocw
func DefWindowProcW(hwnd HWND, msg uint32, wparam WPARAM, lparam LPARAM) LRESULT {
	ret, _, _ := defWindowProcW.Call(uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam))
	return LRESULT(ret)
}

// DeleteMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-deletemenu
func DeleteMenu(hmenu HMENU, position, flags uint32) error {
	if ret, _, err := deleteMenu.Call(uintptr(hmenu), uintptr(position), uintptr(flags)); ret == 0 {
		return errs.NewWithCause(deleteMenu.Name, err)
	}
	return nil
}

// DestroyMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-destroymenu
func DestroyMenu(menu HMENU) error {
	if ret, _, err := destroyMenu.Call(uintptr(menu)); ret == 0 {
		return errs.NewWithCause(destroyMenu.Name, err)
	}
	return nil
}

// DestroyWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-destroywindow
func DestroyWindow(hwnd HWND) error {
	if ret, _, err := destroyWindow.Call(uintptr(hwnd)); ret == 0 {
		return errs.NewWithCause(destroyWindow.Name, err)
	}
	return nil
}

// EnumDisplayDevicesW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevicesW(device string, devNum DWORD, displayDevice *DISPLAY_DEVICEW, flags DWORD) error {
	devstr, err := toUTF16PtrOrNilOnEmpty(device)
	if err != nil {
		return err
	}
	return EnumDisplayDevicesW_(devstr, devNum, displayDevice, flags)
}

// EnumDisplayDevicesW_ https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaydevicesw
func EnumDisplayDevicesW_(device LPCWSTR, devNum DWORD, displayDevice *DISPLAY_DEVICEW, flags DWORD) error {
	if ret, _, err := enumDisplayDevicesW.Call(uintptr(unsafe.Pointer(device)), uintptr(devNum), uintptr(unsafe.Pointer(displayDevice)), uintptr(flags)); ret == 0 {
		return errs.NewWithCause(enumDisplayDevicesW.Name, err)
	}
	return nil
}

// EnumDisplayMonitors https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-enumdisplaymonitors
func EnumDisplayMonitors(hdc HDC, clip *RECT, callback func(monitor HMONITOR, dc HDC, rect *RECT, param LPARAM) BOOL, data LPARAM) error {
	if ret, _, err := enumDisplayMonitors.Call(uintptr(hdc), uintptr(unsafe.Pointer(clip)), syscall.NewCallback(func(monitor HMONITOR, dc HDC, rect *RECT, param LPARAM) uintptr {
		return uintptr(callback(monitor, dc, rect, param))
	}), uintptr(data)); ret == 0 {
		return errs.NewWithCause(enumDisplayMonitors.Name, err)
	}
	return nil
}

// EnumDisplaySettingsExW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaysettingsexw
func EnumDisplaySettingsExW(deviceName string, modeNum DWORD, devMode *DEVMODEW, flags DWORD) error {
	devstr, err := toUTF16PtrOrNilOnEmpty(deviceName)
	if err != nil {
		return err
	}
	return EnumDisplaySettingsExW_(devstr, modeNum, devMode, flags)
}

// EnumDisplaySettingsExW_ https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumdisplaysettingsexw
func EnumDisplaySettingsExW_(deviceName LPCWSTR, modeNum DWORD, devMode *DEVMODEW, flags DWORD) error {
	if ret, _, err := enumDisplaySettingsExW.Call(uintptr(unsafe.Pointer(deviceName)), uintptr(modeNum), uintptr(unsafe.Pointer(devMode)), uintptr(flags)); ret == 0 {
		return errs.NewWithCause(enumDisplaySettingsExW.Name, err)
	}
	return nil
}

// EnumWindows https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumwindows
func EnumWindows(callback func(hwnd HWND, data LPARAM) BOOL, param LPARAM) error {
	if ret, _, err := enumWindows.Call(syscall.NewCallback(func(hwnd HWND, data LPARAM) uintptr {
		return uintptr(callback(hwnd, data))
	}), uintptr(param)); ret == 0 {
		return errs.NewWithCause(enumWindows.Name, err)
	}
	return nil
}

// GetClientRect https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getclientrect
func GetClientRect(hwnd HWND, rect *RECT) error {
	if ret, _, err := getClientRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(rect))); ret == 0 {
		return errs.NewWithCause(getClientRect.Name, err)
	}
	return nil
}

// GetMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getmenu
func GetMenu(hwnd HWND) HMENU {
	ret, _, _ := getMenu.Call(uintptr(hwnd))
	return HMENU(ret)
}

// GetMenuItemCount https://docs.microsoft.com/en-us/windows/desktop/api/Winuser/nf-winuser-getmenuitemcount
func GetMenuItemCount(hmenu HMENU) (int, error) {
	ret, _, err := getMenuItemCount.Call(uintptr(hmenu))
	if ret == ^uintptr(0) { // -1
		return 0, errs.NewWithCause(getMenuItemCount.Name, err)
	}
	return int(ret), nil
}

// GetMenuItemInfoW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getmenuiteminfow
func GetMenuItemInfoW(hmenu HMENU, item uint32, byPosition bool, lpmii *MENUITEMINFOW) error {
	if ret, _, err := getMenuItemInfoW.Call(uintptr(hmenu), uintptr(item), uintptr(toBOOL(byPosition)), uintptr(unsafe.Pointer(lpmii))); ret == 0 {
		return errs.NewWithCause(getMenuItemInfoW.Name, err)
	}
	return nil
}

// GetMonitorInfoW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getmonitorinfow
func GetMonitorInfoW(monitor HMONITOR, pmi *MONITORINFO) error {
	if ret, _, err := getMonitorInfoW.Call(uintptr(monitor), uintptr(unsafe.Pointer(pmi))); ret == 0 {
		return errs.NewWithCause(getMonitorInfoW.Name, err)
	}
	return nil
}

// GetWindowRect https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getwindowrect
func GetWindowRect(hwnd HWND, rect *RECT) error {
	if ret, _, err := getWindowRect.Call(uintptr(hwnd), uintptr(unsafe.Pointer(rect))); ret == 0 {
		return errs.NewWithCause(getWindowRect.Name, err)
	}
	return nil
}

// LoadCursorW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursorw
func LoadCursorW(instance HINSTANCE, cursorName string) (HCURSOR, error) {
	name, err := toUTF16Ptr(cursorName)
	if err != nil {
		return NULL, err
	}
	return LoadCursorW_(instance, name)
}

// InsertMenuItemW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-insertmenuitemw
func InsertMenuItemW(hmenu HMENU, item uint32, byPosition bool, lpmi *MENUITEMINFOW) error {
	if ret, _, err := insertMenuItemW.Call(uintptr(hmenu), uintptr(item), uintptr(toBOOL(byPosition)), uintptr(unsafe.Pointer(lpmi))); ret == 0 {
		return errs.NewWithCause(insertMenuItemW.Name, err)
	}
	return nil
}

// LoadCursorW_ https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursorw
func LoadCursorW_(instance HINSTANCE, cursorName LPCWSTR) (HCURSOR, error) {
	return LoadCursorW__(instance, uintptr(unsafe.Pointer(cursorName)))
}

// LoadCursorW__ https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursorw
func LoadCursorW__(instance HINSTANCE, cursorName uintptr) (HCURSOR, error) {
	h, _, err := loadCursorW.Call(uintptr(instance), cursorName)
	if h == 0 {
		return NULL, errs.NewWithCause(loadCursorW.Name, err)
	}
	return HCURSOR(h), nil
}

// MoveWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-movewindow
func MoveWindow(hwnd HWND, x, y, width, height int32, repaint bool) error {
	if ret, _, err := moveWindow.Call(uintptr(hwnd), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(toBOOL(repaint))); ret == 0 {
		return errs.NewWithCause(moveWindow.Name, err)
	}
	return nil
}

// PostQuitMessage https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-postquitmessage
func PostQuitMessage(exitCode int32) {
	postQuitMessage.Call(uintptr(exitCode))
}

// RegisterClassExW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-registerclassexw
func RegisterClassExW(wndcls *WNDCLASSEXW) (ATOM, error) {
	h, _, err := registerClassExW.Call(uintptr(unsafe.Pointer(wndcls)))
	if h == 0 {
		return 0, errs.NewWithCause(registerClassExW.Name, err)
	}
	return ATOM(h), nil
}

// RegisterWindowMessageW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessageW(name string) (uint32, error) {
	n, err := toUTF16Ptr(name)
	if err != nil {
		return 0, err
	}
	return RegisterWindowMessageW_(n)
}

// RegisterWindowMessageW_ https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-registerwindowmessagew
func RegisterWindowMessageW_(name LPCWSTR) (uint32, error) {
	ret, _, err := registerWindowMessageW.Call(uintptr(unsafe.Pointer(name)))
	if ret == 0 {
		return 0, errs.NewWithCause(registerWindowMessageW.Name, err)
	}
	return uint32(ret), nil
}

// SetActiveWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setactivewindow
func SetActiveWindow(hwnd HWND) error {
	if ret, _, err := setActiveWindow.Call(uintptr(hwnd)); ret == 0 {
		return errs.NewWithCause(setActiveWindow.Name, err)
	}
	return nil
}

// SetMenu https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setmenu
func SetMenu(hwnd HWND, menu HMENU) error {
	if ret, _, err := setMenu.Call(uintptr(hwnd), uintptr(menu)); ret == 0 {
		return errs.NewWithCause(setMenu.Name, err)
	}
	return nil
}

// SetMenuItemInfoW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setmenuiteminfow
func SetMenuItemInfoW(menu HMENU, item uint32, byPosition bool, info *MENUITEMINFOW) error {
	if ret, _, err := setMenuItemInfoW.Call(uintptr(menu), uintptr(item), uintptr(toBOOL(byPosition)), uintptr(unsafe.Pointer(info))); ret == 0 {
		return errs.NewWithCause(setMenuItemInfoW.Name, err)
	}
	return nil
}

// SetProcessDpiAwarenessContext https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setprocessdpiawarenesscontext
func SetProcessDpiAwarenessContext(value DPI_AWARENESS_CONTEXT) error {
	if ret, _, err := setProcessDpiAwarenessContext.Call(uintptr(value)); ret == 0 {
		return errs.NewWithCause(setProcessDpiAwarenessContext.Name, err)
	}
	return nil
}

// SetWindowPos https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowpos
func SetWindowPos(hwnd, hwndInsertAfter HWND, x, y, width, height int32, flags uint32) error {
	if ret, _, err := setWindowPos.Call(uintptr(hwnd), uintptr(hwndInsertAfter), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(flags)); ret == 0 {
		return errs.NewWithCause(setWindowPos.Name, err)
	}
	return nil
}

// SetWindowTextW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowtextw
func SetWindowTextW(hwnd HWND, title string) error {
	str, err := toUTF16Ptr(title)
	if err != nil {
		return err
	}
	return SetWindowTextW_(hwnd, str)
}

// SetWindowTextW_ https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowtextw
func SetWindowTextW_(hwnd HWND, title LPCWSTR) error {
	if ret, _, err := setWindowTextW.Call(uintptr(hwnd), uintptr(unsafe.Pointer(title))); ret == 0 {
		return errs.NewWithCause(setWindowTextW.Name, err)
	}
	return nil
}

// ShowWindow https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-showwindow
func ShowWindow(hwnd HWND, cmd int32) bool {
	if ret, _, _ := showWindow.Call(uintptr(hwnd), uintptr(cmd)); ret == 0 {
		return false
	}
	return true
}
