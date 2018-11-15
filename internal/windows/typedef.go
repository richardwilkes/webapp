package windows

import "syscall"

// https://docs.microsoft.com/en-us/windows/desktop/WinProg/windows-data-types
type (
	ATOM                  uint16
	BOOL                  int16
	DWORD                 uint32
	HBITMAP               uintptr
	HCURSOR               uintptr
	HDC                   uintptr
	HMENU                 uintptr
	HINSTANCE             uintptr
	HMONITOR              uintptr
	HWND                  uintptr
	LPCWSTR               *uint16
	LPVOID                uintptr
	LRESULT               uintptr
	LPARAM                uintptr
	WPARAM                uintptr
	DPI_AWARENESS_CONTEXT uint32
)

// RECT https://msdn.microsoft.com/en-us/9439cb6c-f2f7-4c27-b1d7-8ddf16d81fe8
type RECT struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

// MENUITEMINFOW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmenuiteminfow
type MENUITEMINFOW struct {
	Size         uint32
	Mask         uint32
	Type         uint32
	State        uint32
	ID           uint32
	SubMenu      HMENU
	BMPChecked   HBITMAP
	BMPUnchecked HBITMAP
	ItemData     uintptr
	TypeData     uintptr
	CCH          uint32
	BMPItem      HBITMAP
}

// WNDCLASSEXW https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagwndclassexw
type WNDCLASSEXW struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HINSTANCE
	Icon       syscall.Handle
	Cursor     HCURSOR
	Background syscall.Handle
	MenuName   *uint16
	ClassName  *uint16
	IconSm     syscall.Handle
}

// DISPLAY_DEVICEW https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-_display_devicew
type DISPLAY_DEVICEW struct {
	Size         uint32
	DeviceName   [32]uint16
	DeviceString [128]uint16
	Flags        uint32
	DeviceID     [128]uint16
	DeviceKey    [128]uint16
}

// DEVMODEW https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-_devicemodew
type DEVMODEW struct {
	DeviceName    [32]uint16
	SpecVersion   uint16
	DriverVersion uint16
	Size          uint16
	DriverExtra   uint16
	Fields        uint32
	X             int32
	Y             int32
	Orientation   uint32
	FixedOutput   uint32
	Color         int16
	Duplex        int16
	YResolution   int16
	TTOption      int16
	Collate       int16
	FormName      [32]uint16
	LogPixels     uint16
	BitsPerPixel  uint32
	PelsWidth     uint32
	PelsHeight    uint32
	Flags         uint32
	Frequency     uint32
	ICMMethod     uint32
	ICMIntent     uint32
	MediaType     uint32
	DitherType    uint32
	Reserved1     uint32
	Reserved2     uint32
	PanningWidth  uint32
	PanningHeight uint32
}

// MONITORINFO https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmonitorinfo
type MONITORINFO struct {
	Size          DWORD
	MonitorBounds RECT
	WorkBounds    RECT
	Flags         DWORD
}
