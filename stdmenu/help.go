package stdmenu

import (
	"fmt"
	"runtime"

	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
)

// NewHelpMenu creates a standard 'Help' menu.
func NewHelpMenu(aboutHandler func()) *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuIDHelpMenu, i18n.Text("Help"))
	if runtime.GOOS != "darwin" {
		menu.InsertItem(-1, webapp.MenuIDAboutItem, fmt.Sprintf(i18n.Text("About %s"), cmdline.AppName), 0, 0, nil, aboutHandler)
	}
	return menu
}
