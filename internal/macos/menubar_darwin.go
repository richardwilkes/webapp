package macos

import (
	// #import <stdlib.h>
	// #import "app.h"
	// #import "menus.h"
	"C"
	"fmt"
	"unsafe"

	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

var (
	emptyCStr          = C.CString("")
	handleMenuItemCStr = C.CString("handleMenuItem:")
)

type menuBar struct {
	bar C.CMenuPtr
}

func (d *driver) MenuBarForWindow(_ *webapp.Window) (*webapp.MenuBar, bool, bool) {
	first := false
	if d.menubar == nil {
		m := C.newMenu(emptyCStr)
		d.menubar = &webapp.MenuBar{PlatformData: &menuBar{bar: m}}
		C.setMenuBar(m)
		first = true
	}
	return d.menubar, true, first
}

func (d *driver) MenuBarMenu(bar *webapp.MenuBar, tag int) *webapp.Menu {
	if item := C.menuItemWithTag(bar.PlatformData.(*menuBar).bar, C.int(tag)); item != nil {
		if menu := C.subMenu(item); menu != nil {
			if m, ok := d.menus[menu]; ok {
				return m
			}
		}
	}
	return nil
}

func (d *driver) MenuBarMenuAtIndex(bar *webapp.MenuBar, index int) *webapp.Menu {
	if item := C.menuItemAtIndex(bar.PlatformData.(*menuBar).bar, C.int(index)); item != nil {
		if menu := C.subMenu(item); menu != nil {
			if m, ok := d.menus[menu]; ok {
				return m
			}
		}
	}
	return nil
}

func (d *driver) MenuBarMenuItem(bar *webapp.MenuBar, tag int) *webapp.MenuItem {
	if item := C.menuItemWithTag(bar.PlatformData.(*menuBar).bar, C.int(tag)); item != nil {
		return d.toMenuItem(item)
	}
	return nil
}

func (d *driver) MenuBarInsert(bar *webapp.MenuBar, beforeIndex int, menu *webapp.Menu) {
	cTitle := C.CString(menu.Title)
	mi := C.newMenuItem(C.int(menu.Tag), cTitle, handleMenuItemCStr, emptyCStr, 0, true)
	C.free(unsafe.Pointer(cTitle))
	C.setSubMenu(mi, C.CMenuPtr(menu.PlatformPtr))
	C.insertMenuItem(bar.PlatformData.(*menuBar).bar, mi, C.int(beforeIndex))
}

func (d *driver) MenuBarRemove(bar *webapp.MenuBar, index int) {
	C.removeMenuItem(bar.PlatformData.(*menuBar).bar, C.int(index))
}

func (d *driver) MenuBarCount(bar *webapp.MenuBar) int {
	return int(C.menuItemCount(bar.PlatformData.(*menuBar).bar))
}

func (d *driver) MenuBarSetWindowMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	C.setWindowMenu(C.CMenuPtr(menu.PlatformPtr))
}

func (d *driver) MenuBarSetHelpMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	C.setHelpMenu(C.CMenuPtr(menu.PlatformPtr))
}

func (d *driver) MenuBarFillAppMenu(bar *webapp.MenuBar, aboutHandler, prefsHandler func()) {
	stdMod := keys.PlatformMenuModifier()
	m := webapp.NewMenu(webapp.MenuTagAppMenu, cmdline.AppName)
	m.InsertItem(-1, webapp.MenuTagAboutItem, fmt.Sprintf(i18n.Text("About %s"), cmdline.AppName), 0, 0, nil, aboutHandler)
	if prefsHandler != nil {
		m.InsertSeparator(-1)
		m.InsertItem(-1, webapp.MenuTagPreferencesItem, i18n.Text("Preferences…"), keys.VirtualKeyComma, stdMod, nil, prefsHandler)
	}
	m.InsertSeparator(-1)
	servicesMenu := webapp.NewMenu(webapp.MenuTagServicesMenu, i18n.Text("Services"))
	m.InsertMenu(-1, servicesMenu)
	m.InsertSeparator(-1)
	m.InsertItem(-1, webapp.MenuTagHideItem, fmt.Sprintf(i18n.Text("Hide %s"), cmdline.AppName), keys.VirtualKeyH, stdMod, nil, func() { C.hideApp() })
	m.InsertItem(-1, webapp.MenuTagHideOthersItem, i18n.Text("Hide Others"), keys.VirtualKeyH, keys.OptionModifier|stdMod, nil, func() { C.hideOtherApps() })
	m.InsertItem(-1, webapp.MenuTagShowAllItem, i18n.Text("Show All"), 0, 0, nil, func() { C.showAllApps() })
	m.InsertSeparator(-1)
	m.InsertItem(-1, webapp.MenuTagQuitItem, fmt.Sprintf(i18n.Text("Quit %s"), cmdline.AppName), keys.VirtualKeyQ, stdMod, nil, webapp.AttemptQuit)
	bar.InsertMenu(0, m)
	C.setServicesMenu(C.CMenuPtr(servicesMenu.PlatformPtr))
}
