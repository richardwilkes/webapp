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
	m := C.CMenuPtr(menu.PlatformPtr)
	C.setSubMenu(mi, m)
	C.insertMenuItem(bar.PlatformData.(*menuBar).bar, mi, C.int(beforeIndex))
	switch menu.Tag {
	case webapp.MenuTagAppMenu:
		if servicesMenu := d.MenuBarMenu(bar, webapp.MenuTagServicesMenu); servicesMenu != nil {
			C.setServicesMenu(C.CMenuPtr(servicesMenu.PlatformPtr))
		}
	case webapp.MenuTagWindowMenu:
		C.setWindowMenu(m)
	case webapp.MenuTagHelpMenu:
		C.setHelpMenu(m)
	}
}

func (d *driver) MenuBarRemove(bar *webapp.MenuBar, index int) {
	C.removeMenuItem(bar.PlatformData.(*menuBar).bar, C.int(index))
}

func (d *driver) MenuBarCount(bar *webapp.MenuBar) int {
	return int(C.menuItemCount(bar.PlatformData.(*menuBar).bar))
}
