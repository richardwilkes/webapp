package webapp

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp/keys"
)

// MenuBar represents a set of menus.
type MenuBar struct {
	PlatformData interface{}
}

// MenuBarForWindow returns the menu bar for the given window. On macOS, the
// menu bar is a global entity and the same value will be returned for all
// windows. On macOS, you may pass nil for the window parameter. If isGlobal
// is true, the first time this function is called, isFirst will be true as
// well, allowing you to only initialize the global menu bar once.
func MenuBarForWindow(window *Window) (bar *MenuBar, isGlobal, isFirst bool) {
	return driver.MenuBarForWindow(window)
}

// Menu returns the menu with the specified tag anywhere in the menu bar and
// its sub-menus.
func (bar *MenuBar) Menu(tag int) *Menu {
	return driver.MenuBarMenu(bar, tag)
}

// MenuAtIndex returns the menu at the specified index within the menu bar.
func (bar *MenuBar) MenuAtIndex(index int) *Menu {
	return driver.MenuBarMenuAtIndex(bar, index)
}

// MenuItem returns the menu item with the specified tag anywhere in the menu
// bar and its sub-menus.
func (bar *MenuBar) MenuItem(tag int) *MenuItem {
	return driver.MenuBarMenuItem(bar, tag)
}

// InsertMenu inserts a menu at the specified item index within this menu bar.
// Pass in a negative index to append to the end.
func (bar *MenuBar) InsertMenu(beforeIndex int, menu *Menu) {
	driver.MenuBarInsert(bar, beforeIndex, menu)
}

// Remove the menu at the specified index from this menu bar.
func (bar *MenuBar) Remove(index int) {
	if index >= 0 && index < bar.Count() {
		driver.MenuBarRemove(bar, index)
	}
}

// Count of menus in this menu bar.
func (bar *MenuBar) Count() int {
	return driver.MenuBarCount(bar)
}

// InstallAppMenu adds a standard 'application' menu to the front of the menu
// bar. The handlers may be nil.
func (bar *MenuBar) InstallAppMenu(aboutHandler, prefsHandler func()) {
	driver.MenuBarFillAppMenu(bar, aboutHandler, prefsHandler)
}

// InstallEditMenu adds a standard 'Edit' menu to the end of the menu bar.
func (bar *MenuBar) InstallEditMenu() {
	stdMod := keys.PlatformMenuModifier()
	m := NewMenu(MenuTagEditMenu, i18n.Text("Edit"))
	m.InsertItem(-1, MenuTagCutItem, i18n.Text("Cut"), keys.VirtualKeyX, stdMod, func() bool {
		// RAW: Implement
		return true
	}, func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Browser.Cut()
		}
	})
	m.InsertItem(-1, MenuTagCopyItem, i18n.Text("Copy"), keys.VirtualKeyC, stdMod, func() bool {
		// RAW: Implement
		return true
	}, func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Browser.Copy()
		}
	})
	m.InsertItem(-1, MenuTagPasteItem, i18n.Text("Paste"), keys.VirtualKeyV, stdMod, func() bool {
		// RAW: Implement
		return true
	}, func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Browser.Paste()
		}
	})
	m.InsertSeparator(-1)
	m.InsertItem(-1, MenuTagDeleteItem, i18n.Text("Delete"), keys.VirtualKeyBackspace, 0, func() bool {
		// RAW: Implement
		return true
	}, func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Browser.Delete()
		}
	})
	m.InsertItem(-1, MenuTagSelectAllItem, i18n.Text("Select All"), keys.VirtualKeyA, stdMod, func() bool {
		// RAW: Implement
		return true
	}, func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Browser.SelectAll()
		}
	})
	bar.InsertMenu(-1, m)
}

// InstallWindowMenu adds a standard 'Window' menu to the end of the menu bar.
func (bar *MenuBar) InstallWindowMenu() {
	stdMod := keys.PlatformMenuModifier()
	m := NewMenu(MenuTagWindowMenu, i18n.Text("Window"))
	m.InsertItem(-1, MenuTagMinimizeItem, i18n.Text("Minimize"), keys.VirtualKeyM, stdMod, func() bool {
		w := KeyWindow()
		return w != nil && w.Minimizable()
	}, func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Minimize()
		}
	})
	m.InsertItem(-1, MenuTagZoomItem, i18n.Text("Zoom"), keys.VirtualKeyZ, keys.ShiftModifier|stdMod, func() bool {
		w := KeyWindow()
		return w != nil && w.Resizable()
	}, func() {
		if wnd := KeyWindow(); wnd != nil {
			wnd.Zoom()
		}
	})
	m.InsertSeparator(-1)
	m.InsertItem(-1, MenuTagBringAllWindowsToFrontItem, i18n.Text("Bring All to Front"), 0, 0, nil, AllWindowsToFront)
	bar.InsertMenu(-1, m)
	driver.MenuBarSetWindowMenu(bar, m)
}

// InstallHelpMenu adds a standard 'Help' menu to the end of the menu bar.
func (bar *MenuBar) InstallHelpMenu() {
	m := NewMenu(MenuTagHelpMenu, i18n.Text("Help"))
	bar.InsertMenu(-1, m)
	driver.MenuBarSetHelpMenu(bar, m)
}
