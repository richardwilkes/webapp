package webapp

// Menu represents a set of menu items.
type Menu struct {
	platformMenu
	title string
}

// NewMenu creates a new menu.
func NewMenu(title string) *Menu {
	menu := &Menu{title: title}
	menu.platformInit()
	return menu
}

// Title returns the title of this menu.
func (menu *Menu) Title() string {
	return menu.title
}

// Count of menu items in this menu.
func (menu *Menu) Count() int {
	return menu.platformItemCount()
}

// Item at the specified index, or nil.
func (menu *Menu) Item(index int) *MenuItem {
	return menu.platformItem(index)
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
	menu.platformInsertItem(item, index)
}

// AppendMenu appends a menu item with a sub-menu at the end of this menu.
func (menu *Menu) AppendMenu(subMenu *Menu) {
	menu.InsertMenu(subMenu, -1)
}

// InsertMenu inserts a menu item with a sub-menu at the specified item index
// within this menu. Pass in a negative index to append to the end.
func (menu *Menu) InsertMenu(subMenu *Menu, index int) {
	item := NewMenuItem(subMenu.Title())
	item.platformSetSubMenu(subMenu)
	menu.InsertItem(item, index)
}

// Remove the menu item at the specified index from this menu. This does not
// dispose of the menu item.
func (menu *Menu) Remove(index int) {
	if index >= 0 && index < menu.Count() {
		menu.platformRemove(index)
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
	menu.platformDispose()
}
