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
	Menu    *Menu
	special map[SpecialMenuType]*Menu
	Global  bool
}

// MenuBarForWindow returns the menu bar for the given window. On macOS, the
// menu bar is a global entity and the same value will be returned for all
// windows. On macOS, you may pass nil for the window parameter.
func MenuBarForWindow(window *Window) *MenuBar {
	return driver.MenuBarForWindow(window)
}

// SpecialMenu returns the specified special menu, or nil if it has not been
// setup.
func (bar *MenuBar) SpecialMenu(which SpecialMenuType) *Menu {
	return bar.special[which]
}

// SetupSpecialMenu sets up the specified special menu, which must have
// already been installed into the menu bar.
func (bar *MenuBar) SetupSpecialMenu(which SpecialMenuType, menu *Menu) {
	if bar.special == nil {
		bar.special = make(map[SpecialMenuType]*Menu)
	}
	bar.special[which] = menu
	switch which {
	case ServicesSpecialMenu:
		driver.MenuBarSetServicesMenu(bar, menu)
	case WindowSpecialMenu:
		driver.MenuBarSetWindowMenu(bar, menu)
	case HelpSpecialMenu:
		driver.MenuBarSetHelpMenu(bar, menu)
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
	driver.MenuBarFillAppMenu(bar, appMenu)
	appMenu.AppendItem(NewMenuSeparator())
	appMenu.AppendItem(NewQuitItem())
	bar.Menu.InsertMenu(appMenu, 0)
	return appMenu, aboutItem, prefsItem
}

// InstallEditMenu adds a standard 'Edit' menu to the end of the menu bar.
func (bar *MenuBar) InstallEditMenu() {
	editMenu := NewMenu(i18n.Text("Edit"))
	editMenu.AppendItem(NewSpecialMenuItem(CutKind))
	editMenu.AppendItem(NewSpecialMenuItem(CopyKind))
	editMenu.AppendItem(NewSpecialMenuItem(PasteKind))
	editMenu.AppendItem(NewMenuSeparator())
	editMenu.AppendItem(NewSpecialMenuItem(DeleteKind))
	editMenu.AppendItem(NewSpecialMenuItem(SelectAllKind))
	bar.Menu.AppendMenu(editMenu)
}

// InstallWindowMenu adds a standard 'Window' menu to the end of the menu bar.
func (bar *MenuBar) InstallWindowMenu() {
	windowMenu := NewMenu(i18n.Text("Window"))
	windowMenu.AppendItem(NewMinimizeWindowItem())
	windowMenu.AppendItem(NewZoomWindowItem())
	windowMenu.AppendItem(NewMenuSeparator())
	windowMenu.AppendItem(NewBringAllWindowsToFrontItem())
	bar.Menu.AppendMenu(windowMenu)
	bar.SetupSpecialMenu(WindowSpecialMenu, windowMenu)
}

// InstallHelpMenu adds a standard 'Help' menu to the end of the menu bar.
func (bar *MenuBar) InstallHelpMenu() {
	helpMenu := NewMenu(i18n.Text("Help"))
	bar.Menu.AppendMenu(helpMenu)
	bar.SetupSpecialMenu(HelpSpecialMenu, helpMenu)
}
