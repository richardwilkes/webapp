package webapp

import "github.com/richardwilkes/webapp/keys"

// Pre-defined menu tags. Apps should start their tags at MenuTagUserBase.
const (
	MenuTagAppMenu = 1 + iota
	MenuTagFileMenu
	MenuTagEditMenu
	MenuTagWindowMenu
	MenuTagHelpMenu
	MenuTagServicesMenu
	MenuTagAboutItem
	MenuTagPreferencesItem
	MenuTagQuitItem
	MenuTagCutItem
	MenuTagCopyItem
	MenuTagPasteItem
	MenuTagDeleteItem
	MenuTagSelectAllItem
	MenuTagMinimizeItem
	MenuTagZoomItem
	MenuTagBringAllWindowsToFrontItem
	MenuTagCloseItem
	MenuTagHideItem
	MenuTagHideOthersItem
	MenuTagShowAllItem
	MenuTagUserBase = 1000
)

// Menu represents a set of menu items.
type Menu struct {
	PlatformPtr uintptr
	Tag         int
	Title       string
}

// NewMenu creates a new menu.
func NewMenu(tag int, title string) *Menu {
	menu := &Menu{
		Title: title,
		Tag:   tag,
	}
	driver.MenuInit(menu)
	return menu
}

// ItemAtIndex returns the menu item at the specified index within the menu.
func (menu *Menu) ItemAtIndex(index int) *MenuItem {
	return driver.MenuItemAtIndex(menu, index)
}

// Item returns the menu item with the specified tag anywhere in the menu and
// and its sub-menus.
func (menu *Menu) Item(tag int) *MenuItem {
	return driver.MenuItem(menu, tag)
}

// InsertSeparator inserts a menu separator at the specified item index within
// this menu. Pass in a negative index to append to the end.
func (menu *Menu) InsertSeparator(beforeIndex int) {
	driver.MenuInsertSeparator(menu, beforeIndex)
}

// InsertItem inserts a menu item at the specified item index within this
// menu. Pass in a negative index to append to the end. Both 'validator' and
// 'handler' may be nil for default behavior.
func (menu *Menu) InsertItem(beforeIndex, tag int, title string, keyCode int, keyModifiers keys.Modifiers, validator func() bool, handler func()) {
	if validator == nil {
		validator = func() bool { return true }
	}
	if handler == nil {
		handler = func() {}
	}
	driver.MenuInsertItem(menu, beforeIndex, tag, title, keyCode, keyModifiers, validator, handler)
}

// InsertMenu inserts a sub-menu at the specified item index within this
// menu. Pass in a negative index to append to the end.
func (menu *Menu) InsertMenu(beforeIndex int, subMenu *Menu) {
	driver.MenuInsert(menu, beforeIndex, subMenu)
}

// Remove the menu item at the specified index from this menu.
func (menu *Menu) Remove(index int) {
	if index >= 0 && index < menu.Count() {
		driver.MenuRemove(menu, index)
	}
}

// Count of menu items in this menu.
func (menu *Menu) Count() int {
	return driver.MenuCount(menu)
}

// Dispose releases any operating system resources associated with this menu.
func (menu *Menu) Dispose() {
	driver.MenuDispose(menu)
}
