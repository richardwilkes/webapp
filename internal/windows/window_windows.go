package windows

import (
	"syscall"
)

// WNDCLASSEXW is defined here:
// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagwndclassexw
type WNDCLASSEXW struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   HMODULE
	Icon       syscall.Handle
	Cursor     syscall.Handle
	Background syscall.Handle
	MenuName   *uint16
	ClassName  *uint16
	IconSm     syscall.Handle
}
