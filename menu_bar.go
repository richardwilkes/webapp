package webapp

import (
	"fmt"

	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp/keys"
)

// Special menu ids.
const (
	ServicesSpecialMenu SpecialMenuType = iota
	WindowSpecialMenu
	HelpSpecialMenu
)

// SpecialMenuType identifies which special menu is being referred to.
type SpecialMenuType int

// MenuBar represents a set of menus.
type MenuBar struct {
	platformMenuBar
	bar     *Menu
	special map[SpecialMenuType]*Menu
	global  bool
}

// MenuBarForWindow returns the menu bar for the given window. On macOS, the
// menu bar is a global entity and the same value will be returned for all
// windows. On macOS, you may pass nil for the window parameter.
func MenuBarForWindow(window *Window) *MenuBar {
	return platformMenuBarForWindow(window)
}

// IsGlobal returns true if this MenuBar is global to the application.
func (bar *MenuBar) IsGlobal() bool {
	return bar.global
}

// Count of the top-level menus in this menu bar.
func (bar *MenuBar) Count() int {
	return bar.bar.Count()
}

// AppendMenu appends a menu at the end of this menu bar.
func (bar *MenuBar) AppendMenu(menu *Menu) {
	bar.InsertMenu(menu, -1)
}

// InsertMenu inserts a menu at the specified item index within this menu bar.
// Pass in a negative index to append to the end.
func (bar *MenuBar) InsertMenu(menu *Menu, index int) {
	bar.bar.InsertMenu(menu, index)
}

// Remove the menu at the specified index from this menu bar. This does not
// dispose of the menu.
func (bar *MenuBar) Remove(index int) {
	bar.bar.Remove(index)
}

// Menu at the specified index, or nil.
func (bar *MenuBar) Menu(index int) *Menu {
	return bar.bar.Menu(index)
}

// SpecialMenu returns the specified special menu, or nil if it has not been
// setup.
func (bar *MenuBar) SpecialMenu(which SpecialMenuType) *Menu {
	return bar.special[which]
}

// SetupSpecialMenu sets up the specified special menu, which must have
// already been installed into the menu bar.
func (bar *MenuBar) SetupSpecialMenu(which SpecialMenuType, menu *Menu) {
	bar.special[which] = menu
	switch which {
	case ServicesSpecialMenu:
		bar.platformSetServicesMenu(menu)
	case WindowSpecialMenu:
		bar.platformSetWindowMenu(menu)
	case HelpSpecialMenu:
		bar.platformSetHelpMenu(menu)
	}
}

// InstallAppMenu adds a standard 'application' menu to the front of the menu
// bar.
func (bar *MenuBar) InstallAppMenu() (appMenu *Menu, aboutItem, prefsItem *MenuItem) {
	appMenu = NewMenu(cmdline.AppName)
	aboutItem = NewMenuItem(fmt.Sprintf(i18n.Text("About %s"), cmdline.AppName))
	appMenu.AppendItem(aboutItem)
	appMenu.AppendItem(NewMenuSeparator())
	prefsItem = NewMenuItemWithKey(i18n.Text("Preferences…"), keys.VirtualKeyComma)
	appMenu.AppendItem(prefsItem)
	bar.platformFillAppMenu(appMenu)
	appMenu.AppendItem(NewMenuSeparator())
	appMenu.AppendItem(NewQuitItem())
	bar.InsertMenu(appMenu, 0)
	return appMenu, aboutItem, prefsItem
}

// InstallEditMenu adds a standard 'Edit' menu to the end of the menu bar.
func (bar *MenuBar) InstallEditMenu() {
	editMenu := NewMenu(i18n.Text("Edit"))
	editMenu.AppendItem(NewCutItem())
	editMenu.AppendItem(NewCopyItem())
	editMenu.AppendItem(NewPasteItem())
	editMenu.AppendItem(NewMenuSeparator())
	editMenu.AppendItem(NewDeleteItem())
	editMenu.AppendItem(NewSelectAllItem())
	bar.AppendMenu(editMenu)
}

// InstallWindowMenu adds a standard 'Window' menu to the end of the menu bar.
func (bar *MenuBar) InstallWindowMenu() {
	windowMenu := NewMenu(i18n.Text("Window"))
	windowMenu.AppendItem(NewMinimizeWindowItem())
	windowMenu.AppendItem(NewZoomWindowItem())
	windowMenu.AppendItem(NewMenuSeparator())
	windowMenu.AppendItem(NewBringAllWindowsToFrontItem())
	bar.AppendMenu(windowMenu)
	bar.SetupSpecialMenu(WindowSpecialMenu, windowMenu)
}

// InstallHelpMenu adds a standard 'Help' menu to the end of the menu bar.
func (bar *MenuBar) InstallHelpMenu() {
	helpMenu := NewMenu(i18n.Text("Help"))
	bar.AppendMenu(helpMenu)
	bar.SetupSpecialMenu(HelpSpecialMenu, helpMenu)
}
