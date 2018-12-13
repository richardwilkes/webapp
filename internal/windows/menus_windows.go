package windows

import (
	"strings"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
)

type menuBar struct {
	bar                HMENU
	wnd                HWND
	menuKeys           map[string]*menuItem
	needMenuKeyRefresh bool
	needRedraw         bool
}

type menuItem struct {
	validator func() bool
	handler   func()
	key       *keys.Key
	modifiers keys.Modifiers
}

func (bar *menuBar) markForUpdate() {
	bar.needMenuKeyRefresh = true
	if !bar.needRedraw {
		bar.needRedraw = true
		webapp.InvokeUITask(func() {
			bar.needRedraw = false
			DrawMenuBar(bar.wnd)
		})
	}
}

func (d *driver) markAllForMenuKeyRefresh() {
	for _, bar := range d.menubars {
		bar.PlatformData.(*menuBar).needMenuKeyRefresh = true
	}
}

func (d *driver) refreshMenuKeyForWindow(wnd *webapp.Window) map[string]*menuItem {
	if bar := d.menuBarForWindow(wnd); bar != nil {
		b := bar.PlatformData.(*menuBar)
		if b.needMenuKeyRefresh {
			b.needMenuKeyRefresh = false
			b.menuKeys = make(map[string]*menuItem)
			for i := bar.Count() - 1; i >= 0; i-- {
				d.refreshMenuKeysForMenu(bar, bar.MenuAtIndex(i))
			}
		}
		return b.menuKeys
	}
	return nil
}

func (d *driver) refreshMenuKeysForMenu(bar *webapp.MenuBar, m *webapp.Menu) {
	for i := m.Count() - 1; i >= 0; i-- {
		mi := m.ItemAtIndex(i)
		if mi.SubMenu != nil {
			d.refreshMenuKeysForMenu(bar, mi.SubMenu)
		} else if mi, exists := d.menuitems[mi.ID]; exists {
			if mi.key != nil {
				bar.PlatformData.(*menuBar).menuKeys[mi.modifiers.String()+mi.key.Name] = mi
			}
		}
	}
}

func (d *driver) menuBarForWindow(wnd *webapp.Window) *webapp.MenuBar {
	if wnd != nil {
		if m := GetMenu(HWND(wnd.PlatformPtr)); m != NULL {
			return d.menubars[m]
		}
	}
	return nil
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
		b := &menuBar{
			bar:      m,
			menuKeys: make(map[string]*menuItem),
		}
		d.menubars[m] = &webapp.MenuBar{PlatformData: b}
		b.markForUpdate()
	}
	return d.menubars[m], false, false
}

func (d *driver) MenuBarMenuAtIndex(bar *webapp.MenuBar, index int) *webapp.Menu {
	if item := d.lookupMenuItem(bar.PlatformData.(*menuBar).bar, index); item != nil {
		return item.SubMenu
	}
	return nil
}

func (d *driver) MenuBarInsert(bar *webapp.MenuBar, beforeIndex int, menu *webapp.Menu) {
	b := bar.PlatformData.(*menuBar)
	if err := InsertMenuItemW(b.bar, uint32(beforeIndex), true, &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_ID | MIIM_FTYPE | MIIM_STRING | MIIM_SUBMENU,
		Type:     MFT_STRING,
		ID:       uint32(menu.ID),
		TypeData: uintptr(unsafe.Pointer(mustToUTF16Ptr(menu.Title))),
		SubMenu:  menu.PlatformData.(HMENU),
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
	menu.PlatformData = m
	d.menus[m] = menu
}

func (d *driver) MenuItemAtIndex(menu *webapp.Menu, index int) *webapp.MenuItem {
	return d.lookupMenuItem(menu.PlatformData.(HMENU), index)
}

func (d *driver) lookupMenuItem(menu HMENU, index int) *webapp.MenuItem {
	var data [512]uint16
	info := &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_ID | MIIM_FTYPE | MIIM_STRING | MIIM_SUBMENU,
		TypeData: uintptr(unsafe.Pointer(&data[0])),
		CCH:      uint32(len(data) - 1),
	}
	if err := GetMenuItemInfoW(menu, uint32(index), true, info); err != nil {
		jot.Error(err)
		return nil
	}
	mi := &webapp.MenuItem{
		Owner: d.menus[menu],
		Index: index,
		ID:    int(info.ID),
	}
	if info.Type == MFT_STRING {
		mi.Title = strings.SplitN(syscall.UTF16ToString(data[:info.CCH]), "\t", 2)[0] // Remove any key accelerator info
		mi.SubMenu = d.menus[info.SubMenu]
	}
	return mi
}

func (d *driver) MenuItemAtIndexSetTitle(menu *webapp.Menu, index int, title string) {
	cTitle := syscall.StringToUTF16(title)
	info := &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_STRING,
		TypeData: uintptr(unsafe.Pointer(&cTitle[0])),
	}
	if err := SetMenuItemInfoW(menu.PlatformData.(HMENU), uint32(index), true, info); err != nil {
		jot.Error(err)
	}
}

func (d *driver) MenuInsertSeparator(menu *webapp.Menu, beforeIndex int) {
	if err := InsertMenuItemW(menu.PlatformData.(HMENU), uint32(beforeIndex), true, &MENUITEMINFOW{
		Size: uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask: MIIM_FTYPE,
		Type: MFT_SEPARATOR,
	}); err != nil {
		jot.Error(err)
	}
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, beforeIndex, id int, title string, key *keys.Key, keyModifiers keys.Modifiers, validator func() bool, handler func()) {
	title = strings.SplitN(title, "\t", 2)[0] // Remove any pre-existing key accelerator info
	if key != nil {
		title += "\t" + keyModifiers.String() + key.Name
	}
	if err := InsertMenuItemW(menu.PlatformData.(HMENU), uint32(beforeIndex), true, &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_ID | MIIM_FTYPE | MIIM_STRING,
		Type:     MFT_STRING,
		ID:       uint32(id),
		TypeData: uintptr(unsafe.Pointer(mustToUTF16Ptr(title))),
	}); err != nil {
		jot.Error(err)
	}
	d.menuitems[id] = &menuItem{
		validator: validator,
		handler:   handler,
		key:       key,
		modifiers: keyModifiers,
	}
	d.markAllForMenuKeyRefresh()
}

func (d *driver) MenuInsertMenu(menu *webapp.Menu, beforeIndex, id int, title string) *webapp.Menu {
	subMenu := webapp.NewMenu(id, title)
	if err := InsertMenuItemW(menu.PlatformData.(HMENU), uint32(beforeIndex), true, &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_ID | MIIM_FTYPE | MIIM_STRING | MIIM_SUBMENU,
		Type:     MFT_STRING,
		ID:       uint32(subMenu.ID),
		TypeData: uintptr(unsafe.Pointer(mustToUTF16Ptr(subMenu.Title))),
		SubMenu:  subMenu.PlatformData.(HMENU),
	}); err != nil {
		jot.Error(err)
	}
	d.markAllForMenuKeyRefresh()
	return subMenu
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	if err := DeleteMenu(menu.PlatformData.(HMENU), uint32(index), MF_BYPOSITION); err != nil {
		jot.Error(err)
	}
	d.markAllForMenuKeyRefresh()
}

func (d *driver) MenuCount(menu *webapp.Menu) int {
	count, err := GetMenuItemCount(menu.PlatformData.(HMENU))
	if err != nil {
		jot.Error(err)
		return 0
	}
	return count
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	m := menu.PlatformData.(HMENU)
	delete(d.menus, m)
	if err := DestroyMenu(m); err != nil {
		jot.Error(err)
	}
	d.markAllForMenuKeyRefresh()
}
