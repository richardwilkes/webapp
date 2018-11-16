package webapp

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/toolbox/log/jot"
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
	m.InsertItem(-1, MenuTagCutItem, i18n.Text("Cut"), keys.VirtualKeyX, stdMod, validateCut, handleCut)
	m.InsertItem(-1, MenuTagCopyItem, i18n.Text("Copy"), keys.VirtualKeyC, stdMod, validateCopy, handleCopy)
	m.InsertItem(-1, MenuTagPasteItem, i18n.Text("Paste"), keys.VirtualKeyV, stdMod, validatePaste, handlePaste)
	m.InsertSeparator(-1)
	m.InsertItem(-1, MenuTagDeleteItem, i18n.Text("Delete"), keys.VirtualKeyBackspace, 0, validateDelete, handleDelete)
	m.InsertItem(-1, MenuTagSelectAllItem, i18n.Text("Select All"), keys.VirtualKeyA, stdMod, validateSelectAll, handleSelectAll)
	bar.InsertMenu(-1, m)
}

func validateCut() bool {
	// RAW: Implement
	return true
}

func handleCut() {
	// RAW: Implement
	jot.Info("Cut")
}

func validateCopy() bool {
	// RAW: Implement
	return true
}

func handleCopy() {
	// RAW: Implement
	jot.Info("Copy")
}

func validatePaste() bool {
	// RAW: Implement
	return true
}

func handlePaste() {
	// RAW: Implement
	jot.Info("Paste")
}

func validateDelete() bool {
	// RAW: Implement
	return true
}

func handleDelete() {
	// RAW: Implement
	jot.Info("Delete")
}

func validateSelectAll() bool {
	// RAW: Implement
	return true
}

func handleSelectAll() {
	// RAW: Implement
	jot.Info("Select All")
}

// InstallWindowMenu adds a standard 'Window' menu to the end of the menu bar.
func (bar *MenuBar) InstallWindowMenu() {
	stdMod := keys.PlatformMenuModifier()
	m := NewMenu(MenuTagWindowMenu, i18n.Text("Window"))
	m.InsertItem(-1, MenuTagMinimizeItem, i18n.Text("Minimize"), keys.VirtualKeyM, stdMod, validateMinimize, handleMinimize)
	m.InsertItem(-1, MenuTagZoomItem, i18n.Text("Zoom"), keys.VirtualKeyZ, keys.ShiftModifier|stdMod, validateZoom, handleZoom)
	m.InsertSeparator(-1)
	m.InsertItem(-1, MenuTagBringAllWindowsToFrontItem, i18n.Text("Bring All to Front"), 0, 0, nil, AllWindowsToFront)
	bar.InsertMenu(-1, m)
	driver.MenuBarSetWindowMenu(bar, m)
}

func validateMinimize() bool {
	w := KeyWindow()
	return w != nil && w.Minimizable()
}

func handleMinimize() {
	if wnd := KeyWindow(); wnd != nil {
		wnd.Minimize()
	}
}

func validateZoom() bool {
	w := KeyWindow()
	return w != nil && w.Resizable()
}

func handleZoom() {
	if wnd := KeyWindow(); wnd != nil {
		wnd.Zoom()
	}
}

// InstallHelpMenu adds a standard 'Help' menu to the end of the menu bar.
func (bar *MenuBar) InstallHelpMenu() {
	m := NewMenu(MenuTagHelpMenu, i18n.Text("Help"))
	bar.InsertMenu(-1, m)
	driver.MenuBarSetHelpMenu(bar, m)
}
