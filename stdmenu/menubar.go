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
	"github.com/richardwilkes/webapp"
)

// FillMenuBar adds the standard menus to the menu bar.
func FillMenuBar(bar *webapp.MenuBar, aboutHandler, prefsHandler func(), includeDevTools bool) {
	if runtime.GOOS == toolbox.MacOS {
		bar.InsertMenu(-1, NewAppMenu(aboutHandler, prefsHandler))
	}
	bar.InsertMenu(-1, NewFileMenu())
	bar.InsertMenu(-1, NewEditMenu(prefsHandler))
	bar.InsertMenu(-1, NewWindowMenu())
	bar.InsertMenu(-1, NewHelpMenu(aboutHandler, includeDevTools))
}
