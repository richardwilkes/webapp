package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

type menuBar struct {
	bar        HMENU
	wnd        HWND
	needRedraw bool
}

type menuItem struct {
	validator func() bool
	handler   func()
}

func (bar *menuBar) markForUpdate() {
	if !bar.needRedraw {
		bar.needRedraw = true
		webapp.InvokeUITask(func() {
			bar.needRedraw = false
			DrawMenuBar(bar.wnd)
		})
	}
}

func (d *driver) MenuBarForWindow(wnd *webapp.Window) (*webapp.MenuBar, bool, bool) {
	w := HWND(wnd.PlatformPtr)
	m := GetMenu(w)
	if m == NULL {
		var err error
		if m, err = CreateMenu(); err != nil {
			jot.Error(err)
			return nil, false, false
		}
		if err := SetMenu(w, m); err != nil {
			jot.Error(err)
			return nil, false, false
		}
		b := &menuBar{bar: m}
		d.menubars[m] = &webapp.MenuBar{PlatformData: b}
		b.markForUpdate()
	}
	return d.menubars[m], false, false
}

func (d *driver) MenuBarMenu(bar *webapp.MenuBar, tag int) *webapp.Menu {
	if item := d.lookupMenuItem(bar.PlatformData.(*menuBar).bar, tag, false); item != nil {
		return item.SubMenu
	}
	return nil
}

func (d *driver) MenuBarMenuAtIndex(bar *webapp.MenuBar, index int) *webapp.Menu {
	if item := d.lookupMenuItem(bar.PlatformData.(*menuBar).bar, index, true); item != nil {
		return item.SubMenu
	}
	return nil
}

func (d *driver) MenuBarMenuItem(bar *webapp.MenuBar, tag int) *webapp.MenuItem {
	return d.lookupMenuItem(bar.PlatformData.(*menuBar).bar, tag, false)
}

func (d *driver) MenuBarInsert(bar *webapp.MenuBar, beforeIndex int, menu *webapp.Menu) {
	b := bar.PlatformData.(*menuBar)
	if err := InsertMenuItemW(b.bar, uint32(beforeIndex), true, &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_ID | MIIM_FTYPE | MIIM_STRING | MIIM_SUBMENU,
		Type:     MFT_STRING,
		ID:       uint32(menu.Tag),
		TypeData: uintptr(unsafe.Pointer(mustToUTF16Ptr(menu.Title))),
		SubMenu:  HMENU(menu.PlatformPtr),
	}); err != nil {
		jot.Error(err)
	}
	b.markForUpdate()
}

func (d *driver) MenuBarRemove(bar *webapp.MenuBar, index int) {
	b := bar.PlatformData.(*menuBar)
	if err := DeleteMenu(b.bar, uint32(index), MF_BYPOSITION); err != nil {
		jot.Error(err)
	}
	b.markForUpdate()
}

func (d *driver) MenuBarCount(bar *webapp.MenuBar) int {
	count, err := GetMenuItemCount(bar.PlatformData.(*menuBar).bar)
	if err != nil {
		jot.Error(err)
		return 0
	}
	return count
}

func (d *driver) MenuBarHeightInWindow() float64 {
	return float64(GetSystemMetrics(SM_CYMENU))
}

func (d *driver) MenuInit(menu *webapp.Menu) {
	m, err := CreatePopupMenu()
	if err != nil {
		jot.Error(err)
		return
	}
	menu.PlatformPtr = uintptr(m)
	d.menus[m] = menu
}

func (d *driver) MenuItemAtIndex(menu *webapp.Menu, index int) *webapp.MenuItem {
	return d.lookupMenuItem(HMENU(menu.PlatformPtr), index, true)
}

func (d *driver) MenuItem(menu *webapp.Menu, tag int) *webapp.MenuItem {
	return d.lookupMenuItem(HMENU(menu.PlatformPtr), tag, false)
}

func (d *driver) lookupMenuItem(menu HMENU, item int, byPosition bool) *webapp.MenuItem {
	var data [512]uint16
	info := &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_ID | MIIM_FTYPE | MIIM_STRING | MIIM_SUBMENU,
		TypeData: uintptr(unsafe.Pointer(&data[0])),
		CCH:      uint32(len(data) - 1),
	}
	if err := GetMenuItemInfoW(menu, uint32(item), byPosition, info); err != nil {
		jot.Error(err)
		return nil
	}
	switch info.Type {
	case MFT_STRING:
		return &webapp.MenuItem{
			Tag:     int(info.ID),
			Title:   syscall.UTF16ToString(data[:info.CCH]),
			SubMenu: d.menus[info.SubMenu],
		}
	default:
		return &webapp.MenuItem{Tag: int(info.ID)}
	}
}

func (d *driver) MenuInsertSeparator(menu *webapp.Menu, beforeIndex int) {
	if err := InsertMenuItemW(HMENU(menu.PlatformPtr), uint32(beforeIndex), true, &MENUITEMINFOW{
		Size: uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask: MIIM_FTYPE,
		Type: MFT_SEPARATOR,
	}); err != nil {
		jot.Error(err)
	}
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, beforeIndex, tag int, title string, keyCode int, keyModifiers keys.Modifiers, validator func() bool, handler func()) {
	if err := InsertMenuItemW(HMENU(menu.PlatformPtr), uint32(beforeIndex), true, &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_ID | MIIM_FTYPE | MIIM_STRING,
		Type:     MFT_STRING,
		ID:       uint32(tag),
		TypeData: uintptr(unsafe.Pointer(mustToUTF16Ptr(title))),
	}); err != nil {
		jot.Error(err)
	}
	d.menuitems[tag] = &menuItem{
		validator: validator,
		handler:   handler,
	}
	// RAW: Implement key code support
	// RAW: Implement validator support
}

func (d *driver) MenuInsert(menu *webapp.Menu, beforeIndex int, subMenu *webapp.Menu) {
	if err := InsertMenuItemW(HMENU(menu.PlatformPtr), uint32(beforeIndex), true, &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_ID | MIIM_FTYPE | MIIM_STRING | MIIM_SUBMENU,
		Type:     MFT_STRING,
		ID:       uint32(subMenu.Tag),
		TypeData: uintptr(unsafe.Pointer(mustToUTF16Ptr(subMenu.Title))),
		SubMenu:  HMENU(subMenu.PlatformPtr),
	}); err != nil {
		jot.Error(err)
	}
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	if err := DeleteMenu(HMENU(menu.PlatformPtr), uint32(index), MF_BYPOSITION); err != nil {
		jot.Error(err)
	}
}

func (d *driver) MenuCount(menu *webapp.Menu) int {
	count, err := GetMenuItemCount(HMENU(menu.PlatformPtr))
	if err != nil {
		jot.Error(err)
		return 0
	}
	return count
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	m := HMENU(menu.PlatformPtr)
	delete(d.menus, m)
	if err := DestroyMenu(m); err != nil {
		jot.Error(err)
	}
}
