package windows

import (
	"time"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/internal/cef"
)

type driver struct {
}

var drv = &driver{}

// Driver returns the Windows implementation of the driver.
func Driver() *driver {
	return drv
}

func (d *driver) PrepareForStart() error {
	// RAW: Implement
	return nil
}

func (d *driver) AttemptQuit() {
	// RAW: Implement
}

func (d *driver) MayQuitNow(quit bool) {
	// RAW: Implement
}

func (d *driver) Invoke(id uint64) {
	// RAW: Implement
}

func (d *driver) InvokeAfter(id uint64, after time.Duration) {
	// RAW: Implement
}

func (d *driver) MenuBarForWindow(_ *webapp.Window) *webapp.MenuBar {
	// RAW: Implement
	return nil
}

func (d *driver) MenuBarSetServicesMenu(_ *webapp.MenuBar, menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuBarSetWindowMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuBarSetHelpMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuBarFillAppMenu(bar *webapp.MenuBar, appMenu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuInit(menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuCountItems(menu *webapp.Menu) int {
	return 0
}

func (d *driver) MenuGetItem(menu *webapp.Menu, index int) *webapp.MenuItem {
	// RAW: Implement
	return nil
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, item *webapp.MenuItem, index int) {
	// RAW: Implement
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	// RAW: Implement
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuItemInitSeparator(item *webapp.MenuItem) {
	// RAW: Implement
}

func (d *driver) MenuItemInit(item *webapp.MenuItem, kind webapp.MenuItemKind) {
	// RAW: Implement
}

func (d *driver) MenuItemSubMenu(item *webapp.MenuItem) *webapp.Menu {
	// RAW: Implement
	return nil
}

func (d *driver) MenuItemSetSubMenu(item *webapp.MenuItem, menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuItemDispose(item *webapp.MenuItem) {
	// RAW: Implement
}

func (d *driver) Displays() []*webapp.Display {
	// RAW: Implement
	return nil
}

func (d *driver) KeyWindow() *webapp.Window {
	// RAW: Implement
	return nil
}

func (d *driver) BringAllWindowsToFront() {
	// RAW: Implement
}

func (d *driver) WindowInit(wnd *webapp.Window, style webapp.StyleMask, bounds geom.Rect) {
	// RAW: Implement
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) cef.WindowHandle {
	// RAW: Implement
	return nil
}

func (d *driver) WindowClose(wnd *webapp.Window) {
	// RAW: Implement
}

func (d *driver) WindowSetTitle(wnd *webapp.Window, title string) {
	// RAW: Implement
}

func (d *driver) WindowBounds(wnd *webapp.Window) geom.Rect {
	var bounds geom.Rect
	// RAW: Implement
	return bounds
}

func (d *driver) WindowContentSize(wnd *webapp.Window) geom.Size {
	var size geom.Size
	// RAW: Implement
	return size
}

func (d *driver) WindowSetBounds(wnd *webapp.Window, bounds geom.Rect) {
	// RAW: Implement
}

func (d *driver) WindowToFront(wnd *webapp.Window) {
	// RAW: Implement
}

func (d *driver) WindowMinimize(wnd *webapp.Window) {
	// RAW: Implement
}

func (d *driver) WindowZoom(wnd *webapp.Window) {
	// RAW: Implement
}
