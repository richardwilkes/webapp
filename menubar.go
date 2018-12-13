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

// Menu returns the menu with the specified id anywhere in the menu bar and
// its sub-menus.
func (bar *MenuBar) Menu(id int) *Menu {
	for i := bar.Count() - 1; i >= 0; i-- {
		menu := bar.MenuAtIndex(i)
		if menu.ID == id {
			return menu
		}
		if item := menu.Item(id); item != nil {
			return item.SubMenu
		}
	}
	return nil
}

// MenuAtIndex returns the menu at the specified index within the menu bar.
func (bar *MenuBar) MenuAtIndex(index int) *Menu {
	return driver.MenuBarMenuAtIndex(bar, index)
}

// MenuItem returns the menu item with the specified id anywhere in the menu
// bar and its sub-menus.
func (bar *MenuBar) MenuItem(id int) *MenuItem {
	for i := bar.Count() - 1; i >= 0; i-- {
		menu := bar.MenuAtIndex(i)
		if menu.ID == id {
			return &MenuItem{
				Index:   i,
				ID:      id,
				Title:   menu.Title,
				SubMenu: menu,
			}
		}
		if item := menu.Item(id); item != nil {
			return item
		}
	}
	return nil
}

// InsertMenu inserts a menu at the specified item index within this menu bar.
// Pass in a negative index to append to the end.
func (bar *MenuBar) InsertMenu(beforeIndex int, menu *Menu) {
	driver.MenuBarInsert(bar, beforeIndex, menu)
}

// IndexOf returns the index of the menu within this menu bar, or -1.
func (bar *MenuBar) IndexOf(menu *Menu) int {
	for i := bar.Count() - 1; i >= 0; i-- {
		if bar.MenuAtIndex(i) == menu {
			return i
		}
	}
	return -1
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

// MenuBarHeight returns the height of the MenuBar when displayed in a window.
func MenuBarHeight() float64 {
	return driver.MenuBarHeightInWindow()
}
