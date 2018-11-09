package ws

// From https://docs.microsoft.com/en-us/windows/desktop/winmsg/window-styles
const (
	OVERLAPPED       = 0x00000000
	MAXIMIZEBOX      = 0x00010000 // Has maximize control
	MINIMIZEBOX      = 0x00020000 // Has minimize control
	THICKFRAME       = 0x00040000 // Has a sizing border
	SYSMENU          = 0x00080000 // Has a window menu on title bar
	HSCROLL          = 0x00100000 // Has a horizontal scroll bar
	VSCROLL          = 0x00200000 // Has a vertical scroll bar
	DLGFRAME         = 0x00400000 // Dialog border
	BORDER           = 0x00800000 // Thin line border
	CAPTION          = 0x00C00000 // Has title bar
	MAXIMIZE         = 0x01000000 // Initially maximized
	CLIPCHILDREN     = 0x02000000
	CLIPSIBLINGS     = 0x04000000
	DISABLED         = 0x08000000
	VISIBLE          = 0x10000000 // Initially visible
	MINIMIZE         = 0x20000000 // Initially minimized
	CHILDWINDOW      = 0x40000000
	POPUP            = 0x80000000
	OVERLAPPEDWINDOW = OVERLAPPED | CAPTION | SYSMENU | SYSMENU | THICKFRAME | MINIMIZEBOX | MAXIMIZEBOX
	POPUPWINDOW      = POPUP | BORDER | SYSMENU
)
