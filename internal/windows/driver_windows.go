package windows

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/internal/cef"
	"github.com/richardwilkes/webapp/internal/windows/constants/cs"
	"github.com/richardwilkes/webapp/internal/windows/constants/display"
	"github.com/richardwilkes/webapp/internal/windows/constants/sw"
	"github.com/richardwilkes/webapp/internal/windows/constants/ws"
)

const (
	winIDC_ARROW     = 32512
	winCW_USEDEFAULT = 0x80000000
)

const windowClassName = "wndClass"

type driver struct {
	instance syscall.Handle
}

var drv = &driver{}

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
	cursor, err := LoadCursorW(winIDC_ARROW)
	if err != nil {
		return err
	}
	cnstr, err := syscall.UTF16PtrFromString(windowClassName)
	if err != nil {
		return errs.NewWithCause("Unable to convert className to UTF16", err)
	}
	wcx := WNDCLASSEXW{
		Style:    cs.HREDRAW | cs.VREDRAW,
		WndProc:  syscall.NewCallback(WndProc),
		Instance: d.instance,
		// Icon: LoadIcon(hInstance, MAKEINTRESOURCE(IDI_CEFCLIENT)),
		Cursor: cursor,
		// Background: cCOLOR_WINDOW + 1,
		// MenuName: MAKEINTRESOURCE(IDC_CEFCLIENT),
		ClassName: cnstr,
		// IconSm: LoadIcon(wcex.hInstance, MAKEINTRESOURCE(IDI_SMALL)),
	}
	wcx.Size = uint32(unsafe.Sizeof(wcx))
	_, err = RegisterClassExW(&wcx)
	return err
}

func (d *driver) PrepareForEventLoop() {
	webapp.WillFinishStartupCallback()
	webapp.DidFinishStartupCallback()
}

func (d *driver) AttemptQuit() {
	// RAW: Implement
}

func (d *driver) MayQuitNow(quit bool) {
	// RAW: Implement
}

func (d *driver) Invoke(id uint64) {
	// RAW: Implement
}

func (d *driver) InvokeAfter(id uint64, after time.Duration) {
	// RAW: Implement
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
	var devNum uint32
	result := make([]*webapp.Display, 0)
	for {
		one, err := EnumDisplayDevicesW(devNum, 0)
		if err != nil {
			break
		}
		if one.Flags&display.DEVICE_ACTIVE != 0 {
			dc, err := CreateDCW(&one.DeviceName[0])
			if err != nil {
				jot.Error(err)
			} else {
				d := &webapp.Display{}
				d.Bounds.X = 0 // RAW: Implement
				d.Bounds.Y = 0 // RAW: Implement
				d.Bounds.Width = float64(GetDeviceCaps(dc, HORZRES))
				d.Bounds.Height = float64(GetDeviceCaps(dc, VERTRES))
				d.UsableBounds = d.Bounds // RAW: Account for task bar
				d.ScaleFactor = 1         // RAW: Implement
				d.IsMain = one.Flags&display.DEVICE_PRIMARY_DEVICE != 0
				result = append(result, d)
				if err = DeleteDC(dc); err != nil {
					jot.Error(err)
				}
			}
		}
		devNum++
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
	w, err := CreateWindowExW(0, windowClassName, title, ws.OVERLAPPEDWINDOW|ws.CLIPCHILDREN, int32(bounds.X), int32(bounds.Y), int32(bounds.Height), int32(bounds.Width), 0, 0, d.instance)
	if err != nil {
		return err
	}
	wnd.PlatformPtr = unsafe.Pointer(w)
	return nil
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) cef.WindowHandle {
	return cef.WindowHandle(wnd.PlatformPtr)
}

func (d *driver) WindowClose(wnd *webapp.Window) {
	p := syscall.Handle(wnd.PlatformPtr)
	if err := DestroyWindow(p); err != nil {
		jot.Error(err)
	}
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
