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
	"fmt"
	"runtime"

	"github.com/richardwilkes/toolbox"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
)

// NewHelpMenu creates a standard 'Help' menu.
func NewHelpMenu(aboutHandler func(), includeDevTools bool) *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuIDHelpMenu, i18n.Text("Help"))
	if runtime.GOOS != toolbox.MacOS {
		menu.InsertItem(-1, webapp.MenuIDAboutItem, fmt.Sprintf(i18n.Text("About %s"), cmdline.AppName), nil, 0, nil, aboutHandler)
	}
	if includeDevTools {
		if menu.Count() > 0 {
			menu.InsertSeparator(-1)
		}
		menu.InsertItem(-1, webapp.MenuIDDevToolsItem, i18n.Text("Toggle Development Tools"), nil, 0, validateToggleDevTools, toggleDevTools)
	}
	return menu
}

func validateToggleDevTools() bool {
	wnd := webapp.KeyWindow()
	return wnd != nil && wnd.Browser != nil
}

func toggleDevTools() {
	if wnd := webapp.KeyWindow(); wnd != nil && wnd.Browser != nil {
		wnd.ToggleDevTools()
	}
}
