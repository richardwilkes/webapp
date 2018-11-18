package stdmenu

import (
	"fmt"
	"runtime"

	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

// NewAppMenu creates a standard 'App' menu. Really only intended for macOS,
// although other platforms can use it if desired.
func NewAppMenu(aboutHandler, prefsHandler func()) *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuTagAppMenu, cmdline.AppName)
	menu.InsertItem(-1, webapp.MenuTagAboutItem, fmt.Sprintf(i18n.Text("About %s"), cmdline.AppName), 0, 0, nil, aboutHandler)
	if prefsHandler != nil {
		menu.InsertSeparator(-1)
		menu.InsertItem(-1, webapp.MenuTagPreferencesItem, i18n.Text("Preferences…"), keys.VirtualKeyComma, keys.PlatformMenuModifier(), nil, prefsHandler)
	}
	if runtime.GOOS == "darwin" {
		menu.InsertSeparator(-1)
		menu.InsertMenu(-1, webapp.NewMenu(webapp.MenuTagServicesMenu, i18n.Text("Services")))
	}
	if avc, ok := webapp.PlatformDriver().(webapp.AppVisibilityController); ok {
		menu.InsertSeparator(-1)
		menu.InsertItem(-1, webapp.MenuTagHideItem, fmt.Sprintf(i18n.Text("Hide %s"), cmdline.AppName), keys.VirtualKeyH, keys.PlatformMenuModifier(), nil, avc.HideApp)
		menu.InsertItem(-1, webapp.MenuTagHideOthersItem, i18n.Text("Hide Others"), keys.VirtualKeyH, keys.OptionModifier|keys.PlatformMenuModifier(), nil, avc.HideOtherApps)
		menu.InsertItem(-1, webapp.MenuTagShowAllItem, i18n.Text("Show All"), 0, 0, nil, avc.ShowAllApps)
	}
	menu.InsertSeparator(-1)
	InsertQuitItem(menu, -1)
	return menu
}
