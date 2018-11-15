package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/internal/cef"
)

const windowClassName = "wndClass"

type driver struct {
	instance             HINSTANCE
	windows              map[HWND]*webapp.Window
	menubars             map[HMENU]*webapp.MenuBar
	menus                map[HMENU]*webapp.Menu
	menuitems            map[HMENU]*webapp.MenuItem
	awaitingQuitDecision bool
	forMenuBar           bool
}

var drv = &driver{
	windows:   make(map[HWND]*webapp.Window),
	menubars:  make(map[HMENU]*webapp.MenuBar),
	menus:     make(map[HMENU]*webapp.Menu),
	menuitems: make(map[HMENU]*webapp.MenuItem),
}

// Driver returns the Windows implementation of the driver.
func Driver() *driver {
	return drv
}

func (d *driver) PrepareForStart() error {
	var err error
	if d.instance, err = GetModuleHandleW(); err != nil {
		return err
	}
	wcx := WNDCLASSEXW{
		Style:    CS_HREDRAW | CS_VREDRAW,
		WndProc:  syscall.NewCallback(d.wndProc),
		Instance: d.instance,
		// Icon: LoadIcon(hInstance, MAKEINTRESOURCE(IDI_CEFCLIENT)),
		// Background: cCOLOR_WINDOW + 1,
		// MenuName: MAKEINTRESOURCE(IDC_CEFCLIENT),
		// IconSm: LoadIcon(wcex.hInstance, MAKEINTRESOURCE(IDI_SMALL)),
	}
	wcx.Size = uint32(unsafe.Sizeof(wcx))
	if wcx.Cursor, err = LoadCursorW__(NULL, IDC_ARROW); err != nil {
		return err
	}
	if wcx.ClassName, err = syscall.UTF16PtrFromString(windowClassName); err != nil {
		return errs.NewWithCause("Unable to convert className to UTF16", err)
	}
	_, err = RegisterClassExW(&wcx)
	return err
}

func (d *driver) PrepareForEventLoop() {
	webapp.WillFinishStartupCallback()
	webapp.DidFinishStartupCallback()
}

func (d *driver) wndProc(wnd HWND, msg uint32, wparam WPARAM, lparam LPARAM) LRESULT {
	switch msg {
	case WM_SIZE:
		if w, ok := d.windows[wnd]; ok {
			size := d.WindowContentSize(w)
			SetWindowPos(HWND(unsafe.Pointer(cef.GetWindowHandle(cef.GetBrowserHost(w.Browser)))), 0, 0, 0, int32(size.Width), int32(size.Height), SWP_NOZORDER)
		}
	case WM_CLOSE:
		if w, ok := d.windows[wnd]; ok {
			w.AttemptClose()
		} else {
			if err := DestroyWindow(wnd); err != nil {
				jot.Error(err)
			}
		}
		if len(d.windows) == 0 && webapp.QuitAfterLastWindowClosedCallback() {
			webapp.AttemptQuit()
		}
	case WM_DESTROY:
		PostQuitMessage(0)
	default:
		return DefWindowProcW(wnd, msg, wparam, lparam)
	}
	return 0
}

func (d *driver) AttemptQuit() {
	switch webapp.CheckQuitCallback() {
	case webapp.Cancel:
		return
	case webapp.Now:
		d.quit()
	case webapp.Later:
		d.awaitingQuitDecision = true
	}
}

func (d *driver) MayQuitNow(quit bool) {
	if d.awaitingQuitDecision {
		d.awaitingQuitDecision = false
		if quit {
			d.quit()
		}
	} else {
		jot.Error("Call to MayQuitNow without AttemptQuit")
	}
}

func (d *driver) quit() {
	webapp.QuittingCallback()
	PostQuitMessage(0)
	cef.QuitMessageLoop()
}

func (d *driver) MenuBarForWindow(wnd *webapp.Window) *webapp.MenuBar {
	w := HWND(wnd.PlatformPtr)
	m := GetMenu(w)
	if m == NULL {
		d.forMenuBar = true
		bar := &webapp.MenuBar{Menu: webapp.NewMenu("")}
		d.forMenuBar = false
		m = HMENU(bar.Menu.PlatformPtr)
		if err := SetMenu(w, m); err != nil {
			jot.Error(err)
			return nil
		}
		d.menubars[m] = bar
	}
	return d.menubars[m]
}

func (d *driver) MenuBarSetServicesMenu(_ *webapp.MenuBar, menu *webapp.Menu) {
	// Not relevant to Windows
}

func (d *driver) MenuBarSetWindowMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	// Not relevant to Windows
}

func (d *driver) MenuBarSetHelpMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	// Not relevant to Windows
}

func (d *driver) MenuBarFillAppMenu(bar *webapp.MenuBar, appMenu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuInit(menu *webapp.Menu) {
	var m HMENU
	var err error
	if d.forMenuBar {
		m, err = CreateMenu()
	} else {
		m, err = CreatePopupMenu()
	}
	if err != nil {
		jot.Error(err)
		return
	}
	menu.PlatformPtr = unsafe.Pointer(m)
	if !d.forMenuBar {
		d.menus[m] = menu
	}
}

func (d *driver) MenuCountItems(menu *webapp.Menu) int {
	count, err := GetMenuItemCount(HMENU(menu.PlatformPtr))
	if err != nil {
		jot.Error(err)
		return 0
	}
	return count
}

func (d *driver) MenuGetItem(menu *webapp.Menu, index int) *webapp.MenuItem {
	// RAW: Implement
	return nil
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, item *webapp.MenuItem, index int) {
	// RAW: Implement
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	// RAW: Cleanup menu item that was removed
	if err := DeleteMenu(HMENU(menu.PlatformPtr), uint32(index), MF_BYPOSITION); err != nil {
		jot.Error(err)
	}
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	m := HMENU(menu.PlatformPtr)
	delete(d.menus, m)
	if err := DestroyMenu(m); err != nil {
		jot.Error(err)
	}
}

func (d *driver) MenuItemInitSeparator(item *webapp.MenuItem) {
	m, err := CreateMenu()
	if err != nil {
		jot.Error(err)
		return
	}
	if err = SetMenuItemInfoW(m, 0, true, &MENUITEMINFOW{
		Size: uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask: MIIM_FTYPE,
		Type: MFT_SEPARATOR,
	}); err != nil {
		jot.Error(err)
	}
	item.PlatformPtr = unsafe.Pointer(m)
	d.menuitems[m] = item
}

func (d *driver) MenuItemInit(item *webapp.MenuItem, kind webapp.MenuItemKind) {
	m, err := CreateMenu()
	if err != nil {
		jot.Error(err)
		return
	}
	title, err := toUTF16Ptr(item.Title())
	if err != nil {
		jot.Error(err)
		var empty [1]uint16
		title = &empty[0]
	}
	data := &MENUITEMINFOW{
		Size:     uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:     MIIM_FTYPE | MIIM_STRING,
		Type:     MFT_STRING,
		TypeData: uintptr(unsafe.Pointer(title)),
	}
	if item.Enabled() {
		data.Mask |= MIIM_STATE
		data.State = MFS_DISABLED
	}
	if err = SetMenuItemInfoW(m, 0, true, data); err != nil {
		jot.Error(err)
	}
	item.PlatformPtr = unsafe.Pointer(m)
	d.menuitems[m] = item
}

func (d *driver) MenuItemSubMenu(item *webapp.MenuItem) *webapp.Menu {
	// RAW: Implement
	return nil
}

func (d *driver) MenuItemSetSubMenu(item *webapp.MenuItem, menu *webapp.Menu) {
	if err := SetMenuItemInfoW(HMENU(item.PlatformPtr), 0, true, &MENUITEMINFOW{
		Size:    uint32(unsafe.Sizeof(MENUITEMINFOW{})),
		Mask:    MIIM_SUBMENU,
		SubMenu: HMENU(menu.PlatformPtr),
	}); err != nil {
		jot.Error(err)
	}
}

func (d *driver) MenuItemDispose(item *webapp.MenuItem) {
	m := HMENU(item.PlatformPtr)
	delete(d.menuitems, m)
	if err := DestroyMenu(m); err != nil {
		jot.Error(err)
	}
}

func (d *driver) Displays() []*webapp.Display {
	result := make([]*webapp.Display, 0)
	if err := EnumDisplayMonitors(0, nil, func(monitor HMONITOR, dc HDC, rect *RECT, param LPARAM) BOOL {
		d := &webapp.Display{}
		var info MONITORINFO
		info.Size = DWORD(unsafe.Sizeof(info))
		if err := GetMonitorInfoW(monitor, &info); err != nil {
			jot.Error(err)
		} else {
			d.Bounds.X = float64(info.MonitorBounds.Left)
			d.Bounds.Y = float64(info.MonitorBounds.Top)
			d.Bounds.Width = float64(info.MonitorBounds.Right - info.MonitorBounds.Left)
			d.Bounds.Height = float64(info.MonitorBounds.Bottom - info.MonitorBounds.Top)
			d.UsableBounds.X = float64(info.WorkBounds.Left)
			d.UsableBounds.Y = float64(info.WorkBounds.Top)
			d.UsableBounds.Width = float64(info.WorkBounds.Right - info.WorkBounds.Left)
			d.UsableBounds.Height = float64(info.WorkBounds.Bottom - info.WorkBounds.Top)
			d.IsMain = info.Flags&MONITORINFOF_PRIMARY != 0
			result = append(result, d)
		}
		return 1
	}, 0); err != nil {
		jot.Error(err)
	}
	return result
}

func (d *driver) KeyWindow() *webapp.Window {
	// RAW: Implement
	return nil
}

func (d *driver) BringAllWindowsToFront() {
	list := make([]*webapp.Window, 0)
	if err := EnumWindows(func(wnd HWND, data LPARAM) BOOL {
		if one, ok := d.windows[wnd]; ok {
			list = append(list, one)
		}
		return 1
	}, 0); err != nil {
		jot.Error(err)
		return
	}
	for i, one := range list {
		after := HWND_TOP
		flags := uint32(SWP_NOMOVE | SWP_NOSIZE)
		if i != 0 {
			flags |= SWP_NOACTIVATE
			after = HWND(list[i-1].PlatformPtr)
		}
		if err := SetWindowPos(HWND(one.PlatformPtr), after, 0, 0, 0, 0, flags); err != nil {
			jot.Error(err)
		}
	}
}

func (d *driver) WindowInit(wnd *webapp.Window, style webapp.StyleMask, bounds geom.Rect, title string) error {
	w, err := CreateWindowExW(0, windowClassName, title, WS_OVERLAPPEDWINDOW|WS_CLIPCHILDREN, int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), NULL, NULL, d.instance, NULL)
	if err != nil {
		return err
	}
	wnd.PlatformPtr = unsafe.Pointer(w)
	d.windows[w] = wnd
	return nil
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) cef.WindowHandle {
	return cef.WindowHandle(wnd.PlatformPtr)
}

func (d *driver) WindowClose(wnd *webapp.Window) {
	wnd.WillCloseCallback()
	hwnd := HWND(wnd.PlatformPtr)
	if err := DestroyWindow(hwnd); err != nil {
		jot.Error(err)
	}
	delete(d.windows, hwnd)
}

func (d *driver) WindowSetTitle(wnd *webapp.Window, title string) {
	if err := SetWindowTextW(HWND(wnd.PlatformPtr), title); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowBounds(wnd *webapp.Window) geom.Rect {
	var rect RECT
	if err := GetWindowRect(HWND(wnd.PlatformPtr), &rect); err != nil {
		jot.Error(err)
	}
	var bounds geom.Rect
	bounds.X = float64(rect.Left)
	bounds.Y = float64(rect.Top)
	bounds.Width = float64(rect.Right - rect.Left)
	bounds.Height = float64(rect.Bottom - rect.Top)
	return bounds
}

func (d *driver) WindowContentSize(wnd *webapp.Window) geom.Size {
	var rect RECT
	if err := GetClientRect(HWND(wnd.PlatformPtr), &rect); err != nil {
		jot.Error(err)
	}
	return geom.Size{
		Width:  float64(rect.Right - rect.Left),
		Height: float64(rect.Bottom - rect.Top),
	}
}

func (d *driver) WindowSetBounds(wnd *webapp.Window, bounds geom.Rect) {
	if err := MoveWindow(HWND(wnd.PlatformPtr), int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), true); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowToFront(wnd *webapp.Window) {
	ShowWindow(HWND(wnd.PlatformPtr), SW_SHOWNORMAL)
	if err := SetActiveWindow(HWND(wnd.PlatformPtr)); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowMinimize(wnd *webapp.Window) {
	ShowWindow(HWND(wnd.PlatformPtr), SW_MINIMIZE)
}

func (d *driver) WindowZoom(wnd *webapp.Window) {
	ShowWindow(HWND(wnd.PlatformPtr), SW_MAXIMIZE)
}
