package stdmenu

import (
	"runtime"

	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

// NewFileMenu creates a standard 'File' menu.
func NewFileMenu() *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuTagFileMenu, i18n.Text("File"))
	InsertCloseKeyWindowItem(menu, -1)
	if runtime.GOOS != "darwin" {
		menu.InsertSeparator(-1)
		InsertQuitItem(menu, -1)
	}
	return menu
}

// InsertCloseKeyWindowItem creates the standard "Close" menu item that will
// close the current key window when chosen.
func InsertCloseKeyWindowItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuTagCloseItem, i18n.Text("Close"), keys.VirtualKeyW, keys.PlatformMenuModifier(), CloseKeyWindowValidator, CloseKeyWindowHandler)
}

// CloseKeyWindowValidator provides the standard validation function for the
// "Close" menu.
func CloseKeyWindowValidator() bool {
	wnd := webapp.KeyWindow()
	return wnd != nil && wnd.Closable()
}

// CloseKeyWindowHandler provides the standard handler function for the
// "Close" menu.
func CloseKeyWindowHandler() {
	if wnd := webapp.KeyWindow(); wnd != nil && wnd.Closable() {
		wnd.AttemptClose()
	}
}

// InsertQuitItem creates the standard "Quit"/"Exit" menu item that will
// issue the Quit command when chosen.
func InsertQuitItem(menu *webapp.Menu, beforeIndex int) {
	var title string
	if runtime.GOOS == "darwin" {
		title = i18n.Text("Quit")
	} else {
		title = i18n.Text("Exit")
	}
	menu.InsertItem(-1, webapp.MenuTagQuitItem, title, keys.VirtualKeyQ, keys.PlatformMenuModifier(), nil, webapp.AttemptQuit)
}
