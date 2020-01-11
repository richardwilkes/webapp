// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package stdmenu

import (
	"runtime"

	"github.com/richardwilkes/toolbox"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

// NewFileMenu creates a standard 'File' menu.
func NewFileMenu() *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuIDFileMenu, i18n.Text("File"))
	InsertCloseKeyWindowItem(menu, -1)
	if runtime.GOOS != toolbox.MacOS {
		menu.InsertSeparator(-1)
		InsertQuitItem(menu, -1)
	}
	return menu
}

// InsertCloseKeyWindowItem creates the standard "Close" menu item that will
// close the current key window when chosen.
func InsertCloseKeyWindowItem(menu *webapp.Menu, beforeIndex int) {
	menu.InsertItem(-1, webapp.MenuIDCloseItem, i18n.Text("Close"), keys.W, keys.PlatformMenuModifier(), CloseKeyWindowValidator, CloseKeyWindowHandler)
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
	if runtime.GOOS == toolbox.MacOS {
		title = i18n.Text("Quit")
	} else {
		title = i18n.Text("Exit")
	}
	menu.InsertItem(-1, webapp.MenuIDQuitItem, title, keys.Q, keys.PlatformMenuModifier(), nil, webapp.AttemptQuit)
}
