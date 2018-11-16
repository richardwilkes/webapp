package webapp

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp/keys"
)

// InsertCloseKeyWindowItem creates the standard "Close" menu item that will
// close the current key window when chosen.
func InsertCloseKeyWindowItem(menu *Menu, beforeIndex int) {
	menu.InsertItem(-1, MenuTagCloseItem, i18n.Text("Close"), keys.VirtualKeyW, keys.PlatformMenuModifier(), func() bool {
		wnd := KeyWindow()
		return wnd != nil && wnd.Closable()
	}, func() {
		if wnd := KeyWindow(); wnd != nil && wnd.Closable() {
			wnd.AttemptClose()
		}
	})
}
