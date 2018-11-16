package webapp

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp/internal/cef"
	"github.com/richardwilkes/webapp/keys"
)

// Driver defines the required functions each platform driver must provide.
type Driver interface {
	PrepareForStart() error
	PrepareForEventLoop()

	AttemptQuit()
	MayQuitNow(quit bool)

	MenuBarForWindow(wnd *Window) (bar *MenuBar, isGlobal, isFirst bool)
	MenuBarInsert(bar *MenuBar, beforeIndex int, menu *Menu)
	MenuBarRemove(bar *MenuBar, index int)
	MenuBarCount(bar *MenuBar) int
	MenuBarSetWindowMenu(bar *MenuBar, menu *Menu)
	MenuBarSetHelpMenu(bar *MenuBar, menu *Menu)
	MenuBarFillAppMenu(bar *MenuBar, aboutHandler, prefsHandler func())

	MenuInit(menu *Menu)
	MenuInsertSeparator(menu *Menu, beforeIndex int)
	MenuInsertItem(menu *Menu, beforeIndex, tag int, title string, keyCode int, keyModifiers keys.Modifiers, validator func() bool, handler func())
	MenuInsert(menu *Menu, beforeIndex int, subMenu *Menu)
	MenuRemove(menu *Menu, index int)
	MenuCount(menu *Menu) int
	MenuDispose(menu *Menu)

	Displays() []*Display
	KeyWindow() *Window
	BringAllWindowsToFront()

	WindowInit(wnd *Window, style StyleMask, bounds geom.Rect, title string) error
	WindowBrowserParent(wnd *Window) cef.WindowHandle
	WindowClose(wnd *Window)
	WindowSetTitle(wnd *Window, title string)
	WindowBounds(wnd *Window) geom.Rect
	WindowSetBounds(wnd *Window, bounds geom.Rect)
	WindowContentSize(wnd *Window) geom.Size
	WindowToFront(wnd *Window)
	WindowMinimize(wnd *Window)
	WindowZoom(wnd *Window)
}
