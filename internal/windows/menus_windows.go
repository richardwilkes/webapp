package windows

import (
	"strings"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/keys"
	"github.com/richardwilkes/win32"
)

type menuBar struct {
	bar                win32.HMENU
	wnd                win32.HWND
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
			win32.DrawMenuBar(bar.wnd)
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
		b, ok := bar.PlatformData.(*menuBar)
		if !ok {
			return nil
		}
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
		if w, ok := wnd.PlatformData.(win32.HWND); ok {
			if m := win32.GetMenu(w); m != win32.NULL {
				return d.menubars[m]
			}
		}
	}
	return nil
}

func (d *driver) MenuBarForWindow(wnd *webapp.Window) (*webapp.MenuBar, bool, bool) {
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		m := win32.GetMenu(w)
		if m == win32.NULL {
			if m = win32.CreateMenu(); m != win32.NULL {
				win32.SetMenu(w, m)
				b := &menuBar{
					bar:      m,
					menuKeys: make(map[string]*menuItem),
				}
				d.menubars[m] = &webapp.MenuBar{PlatformData: b}
				b.markForUpdate()
			}
		}
		return d.menubars[m], false, false
	}
	return nil, false, false
}

func (d *driver) MenuBarMenuAtIndex(bar *webapp.MenuBar, index int) *webapp.Menu {
	if item := d.lookupMenuItem(bar.PlatformData.(*menuBar).bar, index); item != nil {
		return item.SubMenu
	}
	return nil
}

func (d *driver) MenuBarInsert(bar *webapp.MenuBar, beforeIndex int, menu *webapp.Menu) {
	if b, ok := bar.PlatformData.(*menuBar); ok {
		win32.InsertMenuItem(b.bar, uint32(beforeIndex), true, &win32.MENUITEMINFO{
			Size:     uint32(unsafe.Sizeof(win32.MENUITEMINFO{})), //nolint:gosec
			Mask:     win32.MIIM_ID | win32.MIIM_FTYPE | win32.MIIM_STRING | win32.MIIM_SUBMENU,
			Type:     win32.MFT_STRING,
			ID:       uint32(menu.ID),
			TypeData: win32.ToSysWin32Str(menu.Title, false),
			SubMenu:  menu.PlatformData.(win32.HMENU),
		})
		b.markForUpdate()
	}
}

func (d *driver) MenuBarRemove(bar *webapp.MenuBar, index int) {
	if b, ok := bar.PlatformData.(*menuBar); ok {
		win32.DeleteMenu(b.bar, uint32(index), win32.MF_BYPOSITION)
		b.markForUpdate()
	}
}

func (d *driver) MenuBarCount(bar *webapp.MenuBar) int {
	return win32.GetMenuItemCount(bar.PlatformData.(*menuBar).bar)
}

func (d *driver) MenuBarHeightInWindow() float64 {
	return float64(win32.GetSystemMetrics(win32.SM_CYMENU))
}

func (d *driver) MenuInit(menu *webapp.Menu) {
	if m := win32.CreatePopupMenu(); m != win32.NULL {
		menu.PlatformData = m
		d.menus[m] = menu
	}
}

func (d *driver) MenuItemAtIndex(menu *webapp.Menu, index int) *webapp.MenuItem {
	return d.lookupMenuItem(menu.PlatformData.(win32.HMENU), index)
}

func (d *driver) lookupMenuItem(menu win32.HMENU, index int) *webapp.MenuItem {
	var data [512]uint16
	info := &win32.MENUITEMINFO{
		Size:     uint32(unsafe.Sizeof(win32.MENUITEMINFO{})), //nolint:gosec
		Mask:     win32.MIIM_ID | win32.MIIM_FTYPE | win32.MIIM_STRING | win32.MIIM_SUBMENU,
		TypeData: uintptr(unsafe.Pointer(&data[0])), //nolint:gosec
		CCH:      uint32(len(data) - 1),
	}
	if !win32.GetMenuItemInfo(menu, uint32(index), true, info) {
		return nil
	}
	mi := &webapp.MenuItem{
		Owner: d.menus[menu],
		Index: index,
		ID:    int(info.ID),
	}
	if info.Type == win32.MFT_STRING {
		mi.Title = strings.SplitN(syscall.UTF16ToString(data[:info.CCH]), "\t", 2)[0] // Remove any key accelerator info
		mi.SubMenu = d.menus[info.SubMenu]
	}
	return mi
}

func (d *driver) MenuItemAtIndexSetTitle(menu *webapp.Menu, index int, title string) {
	win32.SetMenuItemInfo(menu.PlatformData.(win32.HMENU), uint32(index), true, &win32.MENUITEMINFO{
		Size:     uint32(unsafe.Sizeof(win32.MENUITEMINFO{})), //nolint:gosec
		Mask:     win32.MIIM_STRING,
		TypeData: win32.ToSysWin32Str(title, false),
	})
}

func (d *driver) MenuInsertSeparator(menu *webapp.Menu, beforeIndex int) {
	win32.InsertMenuItem(menu.PlatformData.(win32.HMENU), uint32(beforeIndex), true, &win32.MENUITEMINFO{
		Size: uint32(unsafe.Sizeof(win32.MENUITEMINFO{})), //nolint:gosec
		Mask: win32.MIIM_FTYPE,
		Type: win32.MFT_SEPARATOR,
	})
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, beforeIndex, id int, title string, key *keys.Key, keyModifiers keys.Modifiers, validator func() bool, handler func()) {
	title = strings.SplitN(title, "\t", 2)[0] // Remove any pre-existing key accelerator info
	if key != nil {
		title += "\t" + keyModifiers.String() + key.Name
	}
	win32.InsertMenuItem(menu.PlatformData.(win32.HMENU), uint32(beforeIndex), true, &win32.MENUITEMINFO{
		Size:     uint32(unsafe.Sizeof(win32.MENUITEMINFO{})), //nolint:gosec
		Mask:     win32.MIIM_ID | win32.MIIM_FTYPE | win32.MIIM_STRING,
		Type:     win32.MFT_STRING,
		ID:       uint32(id),
		TypeData: win32.ToSysWin32Str(title, false),
	})
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
	win32.InsertMenuItem(menu.PlatformData.(win32.HMENU), uint32(beforeIndex), true, &win32.MENUITEMINFO{
		Size:     uint32(unsafe.Sizeof(win32.MENUITEMINFO{})), //nolint:gosec
		Mask:     win32.MIIM_ID | win32.MIIM_FTYPE | win32.MIIM_STRING | win32.MIIM_SUBMENU,
		Type:     win32.MFT_STRING,
		ID:       uint32(subMenu.ID),
		TypeData: win32.ToSysWin32Str(subMenu.Title, false),
		SubMenu:  subMenu.PlatformData.(win32.HMENU),
	})
	d.markAllForMenuKeyRefresh()
	return subMenu
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	win32.DeleteMenu(menu.PlatformData.(win32.HMENU), uint32(index), win32.MF_BYPOSITION)
	d.markAllForMenuKeyRefresh()
}

func (d *driver) MenuCount(menu *webapp.Menu) int {
	return win32.GetMenuItemCount(menu.PlatformData.(win32.HMENU))
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	if m, ok := menu.PlatformData.(win32.HMENU); ok {
		delete(d.menus, m)
		win32.DestroyMenu(m)
		d.markAllForMenuKeyRefresh()
	}
}
