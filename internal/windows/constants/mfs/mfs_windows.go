package mfs

// From https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmenuiteminfow
const (
	ENABLED   = 0x00000000
	UNCHECKED = ENABLED
	UNHILITE  = ENABLED
	DISABLED  = 0x00000003
	GRAYED    = DISABLED
	CHECKED   = 0x00000008
	HILITE    = 0x00000080
	DEFAULT   = 0x00001000
)
