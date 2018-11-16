package webapp

import (
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp/keys"
)

// InsertCloseKeyWindowItem creates the standard "Close" menu item that will
// close the current key window when chosen.
func InsertCloseKeyWindowItem(m *Menu, beforeIndex int) {
	m.InsertItem(-1, MenuTagCloseItem, i18n.Text("Close"), keys.VirtualKeyW, keys.PlatformMenuModifier(), ValidateClose, HandleClose)
}

func ValidateClose() bool {
	wnd := KeyWindow()
	return wnd != nil && wnd.Closable()
}

func HandleClose() {
	if wnd := KeyWindow(); wnd != nil && wnd.Closable() {
		wnd.AttemptClose()
	}
}
