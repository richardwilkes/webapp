package webapp

import "unsafe"

// Menu represents a set of menu items.
type Menu struct {
	PlatformPtr unsafe.Pointer
	title       string
}

// NewMenu creates a new menu.
func NewMenu(title string) *Menu {
	menu := &Menu{title: title}
	driver.MenuInit(menu)
	return menu
}

// Title returns the title of this menu.
func (menu *Menu) Title() string {
	return menu.title
}

// Count of menu items in this menu.
func (menu *Menu) Count() int {
	return driver.MenuCountItems(menu)
}

// Item at the specified index, or nil.
func (menu *Menu) Item(index int) *MenuItem {
	return driver.MenuGetItem(menu, index)
}

// Menu at the specified index, or nil.
func (menu *Menu) Menu(index int) *Menu {
	if item := menu.Item(index); item != nil {
		return item.SubMenu()
	}
	return nil
}

// AppendItem appends a menu item at the end of this menu.
func (menu *Menu) AppendItem(item *MenuItem) {
	menu.InsertItem(item, -1)
}

// InsertItem inserts a menu item at the specified item index within this
// menu. Pass in a negative index to append to the end.
func (menu *Menu) InsertItem(item *MenuItem, index int) {
	max := menu.Count()
	if index < 0 || index > max {
		index = max
	}
	driver.MenuInsertItem(menu, item, index)
}

// AppendMenu appends a menu item with a sub-menu at the end of this menu.
func (menu *Menu) AppendMenu(subMenu *Menu) {
	menu.InsertMenu(subMenu, -1)
}

// InsertMenu inserts a menu item with a sub-menu at the specified item index
// within this menu. Pass in a negative index to append to the end.
func (menu *Menu) InsertMenu(subMenu *Menu, index int) {
	item := NewMenuItem(subMenu.Title())
	driver.MenuItemSetSubMenu(item, subMenu)
	menu.InsertItem(item, index)
}

// Remove the menu item at the specified index from this menu. This does not
// dispose of the menu item.
func (menu *Menu) Remove(index int) {
	if index >= 0 && index < menu.Count() {
		driver.MenuRemove(menu, index)
	}
}

// Dispose releases any operating system resources associated with this menu.
// It will also call Dispose() on all menu items it contains.
func (menu *Menu) Dispose() {
	for i := menu.Count() - 1; i >= 0; i-- {
		if item := menu.Item(i); item != nil {
			item.Dispose()
		}
	}
	driver.MenuDispose(menu)
}
