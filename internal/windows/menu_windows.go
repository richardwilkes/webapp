package windows

import (
	"syscall"
)

// From https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmenuiteminfow
type MENUITEMINFOW struct {
	Size         uint32
	Mask         uint32
	Type         uint32
	State        uint32
	ID           uint32
	SubMenu      syscall.Handle
	BMPChecked   syscall.Handle
	BMPUnchecked syscall.Handle
	ItemData     uintptr
	TypeData     uintptr
	CCH          uint32
	BMPItem      syscall.Handle
}
