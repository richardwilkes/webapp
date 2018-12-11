package macos

import (
	// #import <stdlib.h>
	// #import "menus.h"
	"C"
	"strings"
	"unsafe"

	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

type menuItem struct {
	validator func() bool
	handler   func()
	item      C.CMenuItemPtr
}

func (d *driver) MenuInit(menu *webapp.Menu) {
	cTitle := C.CString(menu.Title)
	m := C.newMenu(cTitle)
	C.free(unsafe.Pointer(cTitle))
	menu.PlatformData = m
	d.menus[m] = menu
}

func (d *driver) MenuItem(menu *webapp.Menu, id int) *webapp.MenuItem {
	if item := C.menuItemWithID(menu.PlatformData.(C.CMenuPtr), C.int(id)); item != nil {
		return d.toMenuItem(item)
	}
	return nil
}

func (d *driver) MenuItemAtIndex(menu *webapp.Menu, index int) *webapp.MenuItem {
	if item := C.menuItemAtIndex(menu.PlatformData.(C.CMenuPtr), C.int(index)); item != nil {
		return d.toMenuItem(item)
	}
	return nil
}

func (d *driver) toMenuItem(item C.CMenuItemPtr) *webapp.MenuItem {
	info := C.menuItemInfo(item)
	mi := &webapp.MenuItem{
		ID:      int(info.id),
		Title:   C.GoString(info.title),
		SubMenu: d.menus[info.subMenu],
	}
	C.disposeMenuItemInfo(info)
	return mi
}

func (d *driver) MenuInsertSeparator(menu *webapp.Menu, beforeIndex int) {
	C.insertMenuItem(menu.PlatformData.(C.CMenuPtr), C.newMenuSeparator(), C.int(beforeIndex))
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, beforeIndex, id int, title string, keyCode int, keyModifiers keys.Modifiers, validator func() bool, handler func()) {
	var keyCodeStr string
	if keyCode != 0 {
		mapping := keys.MappingForKeyCode(keyCode)
		if mapping.KeyChar != 0 {
			keyCodeStr = strings.ToLower(string(mapping.KeyChar))
		}
	}
	cTitle := C.CString(title)
	cKey := C.CString(keyCodeStr)
	var needDelegate bool
	var selector string
	switch id {
	case webapp.MenuIDCutItem:
		selector = "cut:"
	case webapp.MenuIDCopyItem:
		selector = "copy:"
	case webapp.MenuIDPasteItem:
		selector = "paste:"
	case webapp.MenuIDDeleteItem:
		selector = "delete:"
	case webapp.MenuIDSelectAllItem:
		selector = "selectAll:"
	default:
		selector = "handleMenuItem:"
		needDelegate = true
	}
	cSelector := C.CString(selector)
	mi := C.newMenuItem(C.int(id), cTitle, cSelector, cKey, C.int(keyModifiers), C.bool(needDelegate))
	C.free(unsafe.Pointer(cSelector))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cTitle))
	C.insertMenuItem(menu.PlatformData.(C.CMenuPtr), mi, C.int(beforeIndex))
	d.menuItemValidators[id] = validator
	d.menuItemHandlers[id] = handler
}

func (d *driver) MenuInsert(menu *webapp.Menu, beforeIndex int, subMenu *webapp.Menu) {
	cTitle := C.CString(subMenu.Title)
	mi := C.newMenuItem(C.int(subMenu.ID), cTitle, handleMenuItemCStr, emptyCStr, 0, true)
	C.free(unsafe.Pointer(cTitle))
	C.setSubMenu(mi, subMenu.PlatformData.(C.CMenuPtr))
	C.insertMenuItem(menu.PlatformData.(C.CMenuPtr), mi, C.int(beforeIndex))
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	C.removeMenuItem(menu.PlatformData.(C.CMenuPtr), C.int(index))
}

func (d *driver) MenuCount(menu *webapp.Menu) int {
	return int(C.menuItemCount(menu.PlatformData.(C.CMenuPtr)))
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	p := menu.PlatformData.(C.CMenuPtr)
	C.disposeMenu(p)
	delete(d.menus, p)
}

//export validateMenuItemCallback
func validateMenuItemCallback(id int) bool {
	if validator, ok := drv.menuItemValidators[id]; ok && validator != nil {
		return validator()
	}
	return true
}

//export handleMenuItemCallback
func handleMenuItemCallback(id int) {
	if handler, ok := drv.menuItemHandlers[id]; ok && handler != nil {
		handler()
	}
}
