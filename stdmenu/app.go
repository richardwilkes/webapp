package stdmenu

import (
	"fmt"
	"runtime"

	"github.com/richardwilkes/toolbox"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

// NewAppMenu creates a standard 'App' menu. Really only intended for macOS,
// although other platforms can use it if desired.
func NewAppMenu(aboutHandler, prefsHandler func()) *webapp.Menu {
	menu := webapp.NewMenu(webapp.MenuIDAppMenu, cmdline.AppName)
	menu.InsertItem(-1, webapp.MenuIDAboutItem, fmt.Sprintf(i18n.Text("About %s"), cmdline.AppName), nil, 0, func() bool { return aboutHandler != nil }, aboutHandler)
	if prefsHandler != nil {
		menu.InsertSeparator(-1)
		menu.InsertItem(-1, webapp.MenuIDPreferencesItem, i18n.Text("Preferences…"), keys.Comma, keys.PlatformMenuModifier(), nil, prefsHandler)
	}
	if runtime.GOOS == toolbox.MacOS {
		menu.InsertSeparator(-1)
		menu.InsertMenu(-1, webapp.MenuIDServicesMenu, i18n.Text("Services"))
	}
	if avc, ok := webapp.PlatformDriver().(webapp.AppVisibilityController); ok {
		menu.InsertSeparator(-1)
		menu.InsertItem(-1, webapp.MenuIDHideItem, fmt.Sprintf(i18n.Text("Hide %s"), cmdline.AppName), keys.H, keys.PlatformMenuModifier(), nil, avc.HideApp)
		menu.InsertItem(-1, webapp.MenuIDHideOthersItem, i18n.Text("Hide Others"), keys.H, keys.OptionModifier|keys.PlatformMenuModifier(), nil, avc.HideOtherApps)
		menu.InsertItem(-1, webapp.MenuIDShowAllItem, i18n.Text("Show All"), nil, 0, nil, avc.ShowAllApps)
	}
	menu.InsertSeparator(-1)
	InsertQuitItem(menu, -1)
	return menu
}
