package webapp

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
