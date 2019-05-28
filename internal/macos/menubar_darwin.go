package macos

import (
	// #import <stdlib.h>
	// #import "app.h"
	// #import "menus.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/webapp"
)

var (
	emptyCStr          = C.CString("")
	handleMenuItemCStr = C.CString("handleMenuItem:")
)

func (d *driver) MenuBarForWindow(_ *webapp.Window) (*webapp.MenuBar, bool, bool) {
	first := false
	if d.menubar == nil {
		m := C.newMenu(emptyCStr)
		d.menubar = &webapp.MenuBar{PlatformData: m}
		C.setMenuBar(m)
		first = true
	}
	return d.menubar, true, first
}

func (d *driver) MenuBarMenuAtIndex(bar *webapp.MenuBar, index int) *webapp.Menu {
	if p, ok := bar.PlatformData.(C.CMenuPtr); ok {
		if item := C.menuItemAtIndex(p, C.int(index)); item != nil {
			if menu := C.subMenu(item); menu != nil {
				if m, ok := d.menus[menu]; ok {
					return m
				}
			}
		}
	}
	return nil
}

func (d *driver) MenuBarInsert(bar *webapp.MenuBar, beforeIndex int, menu *webapp.Menu) {
	if p, ok := bar.PlatformData.(C.CMenuPtr); ok {
		if m, ok := menu.PlatformData.(C.CMenuPtr); ok {
			cTitle := C.CString(menu.Title)
			mi := C.newMenuItem(C.int(menu.ID), cTitle, handleMenuItemCStr, emptyCStr, 0, true)
			C.free(unsafe.Pointer(cTitle))
			C.setSubMenu(mi, m)
			C.insertMenuItem(p, mi, C.int(beforeIndex))
			switch menu.ID {
			case webapp.MenuIDAppMenu:
				if servicesMenu := bar.Menu(webapp.MenuIDServicesMenu); servicesMenu != nil {
					if sm, ok := servicesMenu.PlatformData.(C.CMenuPtr); ok {
						C.setServicesMenu(sm)
					}
				}
			case webapp.MenuIDWindowMenu:
				C.setWindowMenu(m)
			case webapp.MenuIDHelpMenu:
				C.setHelpMenu(m)
			}
		}
	}
}

func (d *driver) MenuBarRemove(bar *webapp.MenuBar, index int) {
	if p, ok := bar.PlatformData.(C.CMenuPtr); ok {
		C.removeMenuItem(p, C.int(index))
	}
}

func (d *driver) MenuBarCount(bar *webapp.MenuBar) int {
	if p, ok := bar.PlatformData.(C.CMenuPtr); ok {
		return int(C.menuItemCount(p))
	}
	return 0
}

func (d *driver) MenuBarHeightInWindow() float64 {
	return 0
}
