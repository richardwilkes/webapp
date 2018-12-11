package stdmenu

import (
	"runtime"

	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

// NewEditMenu creates a standard 'Edit' menu.
func NewEditMenu(prefsHandler func()) *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuIDEditMenu, i18n.Text("Edit"))
	InsertCutItem(menu, -1)
	InsertCopyItem(menu, -1)
	InsertPasteItem(menu, -1)
	InsertDeleteItem(menu, -1)
	InsertSelectAllItem(menu, -1)
	if runtime.GOOS != "darwin" && prefsHandler != nil {
		menu.InsertSeparator(-1)
		menu.InsertItem(-1, webapp.MenuIDPreferencesItem, i18n.Text("Preferences…"), keys.VirtualKeyComma, keys.PlatformMenuModifier(), nil, prefsHandler)
	}
	return menu
}

// InsertCutItem creates the standard "Cut" menu item that will issue the Cut
// command to the current key window when chosen.
func InsertCutItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuIDCutItem, i18n.Text("Cut"), keys.VirtualKeyX, keys.PlatformMenuModifier(), CutValidator, CutHandler)
}

// CutValidator provides the standard validation function for the "Cut" menu
// item.
func CutValidator() bool {
	// RAW: Implement
	return true
}

// CutHandler provides the standard handler function for the "Cut" menu item.
func CutHandler() {
	if wnd := webapp.KeyWindow(); wnd != nil {
		if ff := wnd.Browser.GetFocusedFrame(); ff != nil {
			ff.Cut()
		}
	}
}

// InsertCopyItem creates the standard "Copy" menu item that will issue the
// Copy command to the current key window when chosen.
func InsertCopyItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuIDCopyItem, i18n.Text("Copy"), keys.VirtualKeyC, keys.PlatformMenuModifier(), CopyValidator, CopyHandler)
}

// CopyValidator provides the standard validation function for the "Copy" menu
// item.
func CopyValidator() bool {
	// RAW: Implement
	return true
}

// CopyHandler provides the standard handler function for the "Copy" menu
// item.
func CopyHandler() {
	if wnd := webapp.KeyWindow(); wnd != nil {
		if ff := wnd.Browser.GetFocusedFrame(); ff != nil {
			ff.Copy()
		}
	}
}

// InsertPasteItem creates the standard "Paste" menu item that will issue the
// Paste command to the current key window when chosen.
func InsertPasteItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuIDPasteItem, i18n.Text("Paste"), keys.VirtualKeyV, keys.PlatformMenuModifier(), PasteValidator, PasteHandler)
}

// PasteValidator provides the standard validation function for the "Paste"
// menu item.
func PasteValidator() bool {
	// RAW: Implement
	return true
}

// PasteHandler provides the standard handler function for the "Paste" menu
// item.
func PasteHandler() {
	if wnd := webapp.KeyWindow(); wnd != nil {
		if ff := wnd.Browser.GetFocusedFrame(); ff != nil {
			ff.Paste()
		}
	}
}

// InsertDeleteItem creates the standard "Delete" menu item that will issue
// the Delete command to the current key window when chosen.
func InsertDeleteItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuIDDeleteItem, i18n.Text("Delete"), keys.VirtualKeyBackspace, 0, DeleteValidator, DeleteHandler)
}

// DeleteValidator provides the standard validation function for the "Delete"
// menu item.
func DeleteValidator() bool {
	// RAW: Implement
	return true
}

// DeleteHandler provides the standard handler function for the "Delete" menu
// item.
func DeleteHandler() {
	if wnd := webapp.KeyWindow(); wnd != nil {
		if ff := wnd.Browser.GetFocusedFrame(); ff != nil {
			ff.Del()
		}
	}
}

// InsertSelectAllItem creates the standard "Select All" menu item that will
// issue the SelectAll command to the current key window when chosen.
func InsertSelectAllItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuIDSelectAllItem, i18n.Text("Select All"), keys.VirtualKeyA, keys.PlatformMenuModifier(), SelectAllValidator, SelectAllHandler)
}

// SelectAllValidator provides the standard validation function for the
// "Select All" menu item.
func SelectAllValidator() bool {
	// RAW: Implement
	return true
}

// SelectAllHandler provides the standard handler function for the
// "Select All" menu item.
func SelectAllHandler() {
	if wnd := webapp.KeyWindow(); wnd != nil {
		if ff := wnd.Browser.GetFocusedFrame(); ff != nil {
			ff.SelectAll()
		}
	}
}
