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
	menu.PlatformPtr = uintptr(m)
	d.menus[m] = menu
}

func (d *driver) MenuItem(menu *webapp.Menu, tag int) *webapp.MenuItem {
	if item := C.menuItemWithTag(C.CMenuPtr(menu.PlatformPtr), C.int(tag)); item != nil {
		return d.toMenuItem(item)
	}
	return nil
}

func (d *driver) MenuItemAtIndex(menu *webapp.Menu, index int) *webapp.MenuItem {
	if item := C.menuItemAtIndex(C.CMenuPtr(menu.PlatformPtr), C.int(index)); item != nil {
		return d.toMenuItem(item)
	}
	return nil
}

func (d *driver) toMenuItem(item C.CMenuItemPtr) *webapp.MenuItem {
	info := C.menuItemInfo(item)
	mi := &webapp.MenuItem{
		Tag:     int(info.tag),
		Title:   C.GoString(info.title),
		SubMenu: d.menus[info.subMenu],
	}
	C.disposeMenuItemInfo(info)
	return mi
}

func (d *driver) MenuInsertSeparator(menu *webapp.Menu, beforeIndex int) {
	C.insertMenuItem(C.CMenuPtr(menu.PlatformPtr), C.newMenuSeparator(), C.int(beforeIndex))
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, beforeIndex, tag int, title string, keyCode int, keyModifiers keys.Modifiers, validator func() bool, handler func()) {
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
	switch tag {
	case webapp.MenuTagCutItem:
		selector = "cut:"
	case webapp.MenuTagCopyItem:
		selector = "copy:"
	case webapp.MenuTagPasteItem:
		selector = "paste:"
	case webapp.MenuTagDeleteItem:
		selector = "delete:"
	case webapp.MenuTagSelectAllItem:
		selector = "selectAll:"
	default:
		selector = "handleMenuItem:"
		needDelegate = true
	}
	cSelector := C.CString(selector)
	mi := C.newMenuItem(C.int(tag), cTitle, cSelector, cKey, C.int(keyModifiers), C.bool(needDelegate))
	C.free(unsafe.Pointer(cSelector))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cTitle))
	C.insertMenuItem(C.CMenuPtr(menu.PlatformPtr), mi, C.int(beforeIndex))
	d.menuItemValidators[tag] = validator
	d.menuItemHandlers[tag] = handler
}

func (d *driver) MenuInsert(menu *webapp.Menu, beforeIndex int, subMenu *webapp.Menu) {
	cTitle := C.CString(subMenu.Title)
	mi := C.newMenuItem(C.int(subMenu.Tag), cTitle, handleMenuItemCStr, emptyCStr, 0, true)
	C.free(unsafe.Pointer(cTitle))
	C.setSubMenu(mi, C.CMenuPtr(subMenu.PlatformPtr))
	C.insertMenuItem(C.CMenuPtr(menu.PlatformPtr), mi, C.int(beforeIndex))
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	C.removeMenuItem(C.CMenuPtr(menu.PlatformPtr), C.int(index))
}

func (d *driver) MenuCount(menu *webapp.Menu) int {
	return int(C.menuItemCount(C.CMenuPtr(menu.PlatformPtr)))
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	p := C.CMenuPtr(menu.PlatformPtr)
	C.disposeMenu(p)
	delete(d.menus, p)
}

//export validateMenuItemCallback
func validateMenuItemCallback(tag int) bool {
	if validator, ok := drv.menuItemValidators[tag]; ok && validator != nil {
		return validator()
	}
	return true
}

//export handleMenuItemCallback
func handleMenuItemCallback(tag int) {
	if handler, ok := drv.menuItemHandlers[tag]; ok && handler != nil {
		handler()
	}
}
