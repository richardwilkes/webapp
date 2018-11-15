package webapp

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp/internal/cef"
)

// Driver defines the required functions each platform driver must provide.
type Driver interface {
	PrepareForStart() error
	PrepareForEventLoop()

	AttemptQuit()
	MayQuitNow(quit bool)

	MenuBarForWindow(wnd *Window) *MenuBar
	MenuBarSetServicesMenu(bar *MenuBar, menu *Menu)
	MenuBarSetWindowMenu(bar *MenuBar, menu *Menu)
	MenuBarSetHelpMenu(bar *MenuBar, menu *Menu)
	MenuBarFillAppMenu(bar *MenuBar, appMenu *Menu)

	MenuInit(menu *Menu)
	MenuCountItems(menu *Menu) int
	MenuGetItem(menu *Menu, index int) *MenuItem
	MenuInsertItem(menu *Menu, item *MenuItem, index int)
	MenuRemove(menu *Menu, index int)
	MenuDispose(menu *Menu)

	MenuItemInitSeparator(item *MenuItem)
	MenuItemInit(item *MenuItem, kind MenuItemKind)
	MenuItemSubMenu(item *MenuItem) *Menu
	MenuItemSetSubMenu(item *MenuItem, menu *Menu)
	MenuItemDispose(item *MenuItem)

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
