// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package macos

import (
	"unsafe"

	"github.com/richardwilkes/webapp"
)

import (
	// #import <stdlib.h>
	// #import "app.h"
	// #import "menus.h"
	"C"
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
				if m, ok2 := d.menus[menu]; ok2 {
					return m
				}
			}
		}
	}
	return nil
}

func (d *driver) MenuBarInsert(bar *webapp.MenuBar, beforeIndex int, menu *webapp.Menu) {
	if p, ok := bar.PlatformData.(C.CMenuPtr); ok {
		if m, ok2 := menu.PlatformData.(C.CMenuPtr); ok2 {
			cTitle := C.CString(menu.Title)
			mi := C.newMenuItem(C.int(menu.ID), cTitle, handleMenuItemCStr, emptyCStr, 0, true)
			C.free(unsafe.Pointer(cTitle))
			C.setSubMenu(mi, m)
			C.insertMenuItem(p, mi, C.int(beforeIndex))
			switch menu.ID {
			case webapp.MenuIDAppMenu:
				if servicesMenu := bar.Menu(webapp.MenuIDServicesMenu); servicesMenu != nil {
					if sm, ok3 := servicesMenu.PlatformData.(C.CMenuPtr); ok3 {
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
