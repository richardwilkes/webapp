package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/internal/cef"
	"github.com/richardwilkes/webapp/internal/windows/constants/cs"
	"github.com/richardwilkes/webapp/internal/windows/constants/sw"
	"github.com/richardwilkes/webapp/internal/windows/constants/swp"
	"github.com/richardwilkes/webapp/internal/windows/constants/wm"
	"github.com/richardwilkes/webapp/internal/windows/constants/ws"
)

const (
	winIDC_ARROW     = 32512
	winCW_USEDEFAULT = 0x80000000
)

const windowClassName = "wndClass"

type driver struct {
	instance             syscall.Handle
	windows              map[syscall.Handle]*webapp.Window
	awaitingQuitDecision bool
}

var drv = &driver{
	windows: make(map[syscall.Handle]*webapp.Window),
}

// Driver returns the Windows implementation of the driver.
func Driver() *driver {
	return drv
}

func (d *driver) Initialize() error {
	var err error
	if d.instance, err = GetModuleHandleW(); err != nil {
		return err
	}
	cef.ExecuteProcess(d.instance)
	return nil
}

func (d *driver) PrepareForStart() error {
	wcx := WNDCLASSEXW{
		Style:    cs.HREDRAW | cs.VREDRAW,
		WndProc:  syscall.NewCallback(d.wndProc),
		Instance: d.instance,
		// Icon: LoadIcon(hInstance, MAKEINTRESOURCE(IDI_CEFCLIENT)),
		// Background: cCOLOR_WINDOW + 1,
		// MenuName: MAKEINTRESOURCE(IDC_CEFCLIENT),
		// IconSm: LoadIcon(wcex.hInstance, MAKEINTRESOURCE(IDI_SMALL)),
	}
	wcx.Size = uint32(unsafe.Sizeof(wcx))
	var err error
	if wcx.Cursor, err = LoadCursorW(winIDC_ARROW); err != nil {
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

func (d *driver) wndProc(wnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case wm.SIZE:
		if w, ok := d.windows[wnd]; ok {
			size := d.WindowContentSize(w)
			SetWindowPos(syscall.Handle(unsafe.Pointer(cef.GetWindowHandle(cef.GetBrowserHost(w.Browser)))), 0, 0, 0, int32(size.Width), int32(size.Height), swp.NOZORDER)
		}
	case wm.CLOSE:
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
	case wm.DESTROY:
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

func (d *driver) MenuBarForWindow(_ *webapp.Window) *webapp.MenuBar {
	// RAW: Implement
	return &webapp.MenuBar{
		Menu: webapp.NewMenu(""),
	}
}

func (d *driver) MenuBarSetServicesMenu(_ *webapp.MenuBar, menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuBarSetWindowMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuBarSetHelpMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuBarFillAppMenu(bar *webapp.MenuBar, appMenu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuInit(menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuCountItems(menu *webapp.Menu) int {
	return 0
}

func (d *driver) MenuGetItem(menu *webapp.Menu, index int) *webapp.MenuItem {
	// RAW: Implement
	return nil
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, item *webapp.MenuItem, index int) {
	// RAW: Implement
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	// RAW: Implement
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuItemInitSeparator(item *webapp.MenuItem) {
	// RAW: Implement
}

func (d *driver) MenuItemInit(item *webapp.MenuItem, kind webapp.MenuItemKind) {
	// RAW: Implement
}

func (d *driver) MenuItemSubMenu(item *webapp.MenuItem) *webapp.Menu {
	// RAW: Implement
	return nil
}

func (d *driver) MenuItemSetSubMenu(item *webapp.MenuItem, menu *webapp.Menu) {
	// RAW: Implement
}

func (d *driver) MenuItemDispose(item *webapp.MenuItem) {
	// RAW: Implement
}

func (d *driver) Displays() []*webapp.Display {
	result := make([]*webapp.Display, 0)
	if err := EnumDisplayMonitors(0, nil, func(monitor, dc syscall.Handle, rect, param uintptr) uintptr {
		d := &webapp.Display{}
		if info, err := GetMonitorInfoW(monitor); err != nil {
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
	// RAW: Implement
}

func (d *driver) WindowInit(wnd *webapp.Window, style webapp.StyleMask, bounds geom.Rect, title string) error {
	w, err := CreateWindowExW(0, windowClassName, title, ws.OVERLAPPEDWINDOW|ws.CLIPCHILDREN, int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), 0, 0, d.instance)
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
	p := syscall.Handle(wnd.PlatformPtr)
	if err := DestroyWindow(p); err != nil {
		jot.Error(err)
	}
	delete(d.windows, p)
}

func (d *driver) WindowSetTitle(wnd *webapp.Window, title string) {
	if err := SetWindowTextW(syscall.Handle(wnd.PlatformPtr), title); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowBounds(wnd *webapp.Window) geom.Rect {
	bounds, err := GetWindowRect(syscall.Handle(wnd.PlatformPtr))
	if err != nil {
		jot.Error(err)
	}
	return bounds
}

func (d *driver) WindowContentSize(wnd *webapp.Window) geom.Size {
	bounds, err := GetClientRect(syscall.Handle(wnd.PlatformPtr))
	if err != nil {
		jot.Error(err)
	}
	return bounds.Size
}

func (d *driver) WindowSetBounds(wnd *webapp.Window, bounds geom.Rect) {
	if err := MoveWindow(syscall.Handle(wnd.PlatformPtr), int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), true); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowToFront(wnd *webapp.Window) {
	ShowWindow(syscall.Handle(wnd.PlatformPtr), sw.SHOWNORMAL)
	if err := SetActiveWindow(syscall.Handle(wnd.PlatformPtr)); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowMinimize(wnd *webapp.Window) {
	ShowWindow(syscall.Handle(wnd.PlatformPtr), sw.MINIMIZE)
}

func (d *driver) WindowZoom(wnd *webapp.Window) {
	ShowWindow(syscall.Handle(wnd.PlatformPtr), sw.MAXIMIZE)
}
