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
