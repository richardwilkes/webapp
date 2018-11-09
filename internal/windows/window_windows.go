package windows

import (
	"syscall"

	"github.com/richardwilkes/webapp/internal/windows/constants/wm"
)

// WNDCLASSEXW is defined here:
// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagwndclassexw
type WNDCLASSEXW struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   syscall.Handle
	Icon       syscall.Handle
	Cursor     syscall.Handle
	Background syscall.Handle
	MenuName   *uint16
	ClassName  *uint16
	IconSm     syscall.Handle
}

// WndProc provides standard handling of window messages.
func WndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case wm.CLOSE:
		DestroyWindow(hwnd)
	case wm.DESTROY:
		PostQuitMessage(0)
	default:
		return DefWindowProcW(hwnd, msg, wparam, lparam)
	}
	return 0
}
