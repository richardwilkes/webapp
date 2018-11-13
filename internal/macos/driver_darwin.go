package macos

import (
	// #cgo CFLAGS: -x objective-c -I ${SRCDIR}/../../cef
	// #cgo LDFLAGS: -framework Cocoa
	// #include "driver.h"
	"C"
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/internal/cef"
	"github.com/richardwilkes/webapp/internal/uitask"
	"github.com/richardwilkes/webapp/keys"
)

type driver struct {
	menubar   *webapp.MenuBar
	menus     map[C.CMenuPtr]*webapp.Menu
	menuItems map[C.CMenuItemPtr]*webapp.MenuItem
	windows   map[C.CWindowPtr]*webapp.Window
}

var drv = &driver{
	menus:     make(map[C.CMenuPtr]*webapp.Menu),
	menuItems: make(map[C.CMenuItemPtr]*webapp.MenuItem),
	windows:   make(map[C.CWindowPtr]*webapp.Window),
}

// Driver returns the macOS implementation of the driver.
func Driver() *driver {
	return drv
}

func (d *driver) Initialize() error {
	// Nothing to do
	return nil
}

func (d *driver) PrepareForStart() error {
	C.prepareForStart()
	return nil
}

func (d *driver) PrepareForEventLoop() {
	// Nothing to do
}

func (d *driver) AttemptQuit() {
	C.attemptQuit()
}

func (d *driver) MayQuitNow(quit bool) {
	var mayQuit C.int
	if quit {
		mayQuit = 1
	}
	C.mayQuitNow(mayQuit)
}

func (d *driver) Invoke(id uint64) {
	C.invoke(C.ulong(id))
}

func (d *driver) InvokeAfter(id uint64, after time.Duration) {
	C.invokeAfter(C.ulong(id), C.long(after.Nanoseconds()))
}

func (d *driver) MenuBarForWindow(_ *webapp.Window) *webapp.MenuBar {
	if d.menubar == nil {
		d.menubar = &webapp.MenuBar{
			Menu:   webapp.NewMenu(""),
			Global: true,
		}
		C.setMenuBar(C.CMenuPtr(d.menubar.Menu.PlatformPtr))
	}
	return d.menubar
}

func (d *driver) MenuBarSetServicesMenu(_ *webapp.MenuBar, menu *webapp.Menu) {
	C.setServicesMenu(C.CMenuPtr(menu.PlatformPtr))
}

func (d *driver) MenuBarSetWindowMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	C.setWindowMenu(C.CMenuPtr(menu.PlatformPtr))
}

func (d *driver) MenuBarSetHelpMenu(bar *webapp.MenuBar, menu *webapp.Menu) {
	C.setHelpMenu(C.CMenuPtr(menu.PlatformPtr))
}

func (d *driver) MenuBarFillAppMenu(bar *webapp.MenuBar, appMenu *webapp.Menu) {
	appMenu.AppendItem(webapp.NewMenuSeparator())
	servicesMenu := webapp.NewMenu(i18n.Text("Services"))
	appMenu.AppendMenu(servicesMenu)
	bar.SetupSpecialMenu(webapp.ServicesSpecialMenu, servicesMenu)

	appMenu.AppendItem(webapp.NewMenuSeparator())
	hideItem := webapp.NewMenuItemWithKey(fmt.Sprintf(i18n.Text("Hide %s"), cmdline.AppName), keys.VirtualKeyH)
	hideItem.Handler = func() { C.hideApp() }
	appMenu.AppendItem(hideItem)
	hideOthersItem := webapp.NewMenuItemWithKeyAndModifiers(i18n.Text("Hide Others"), keys.VirtualKeyH, keys.OptionModifier|keys.PlatformMenuModifier())
	hideOthersItem.Handler = func() { C.hideOtherApps() }
	appMenu.AppendItem(hideOthersItem)
	showAllItem := webapp.NewMenuItem(i18n.Text("Show All"))
	showAllItem.Handler = func() { C.showAllApps() }
	appMenu.AppendItem(showAllItem)
}

func (d *driver) MenuInit(menu *webapp.Menu) {
	cTitle := C.CString(menu.Title())
	m := C.newMenu(cTitle)
	C.free(unsafe.Pointer(cTitle))
	menu.PlatformPtr = unsafe.Pointer(m)
	d.menus[m] = menu
}

func (d *driver) MenuCountItems(menu *webapp.Menu) int {
	return int(C.menuItemCount(C.CMenuPtr(menu.PlatformPtr)))
}

func (d *driver) MenuGetItem(menu *webapp.Menu, index int) *webapp.MenuItem {
	return d.menuItems[C.menuItem(C.CMenuPtr(menu.PlatformPtr), C.int(index))]
}

func (d *driver) MenuInsertItem(menu *webapp.Menu, item *webapp.MenuItem, index int) {
	C.insertMenuItem(C.CMenuPtr(menu.PlatformPtr), C.CMenuItemPtr(item.PlatformPtr), C.int(index))
}

func (d *driver) MenuRemove(menu *webapp.Menu, index int) {
	C.removeMenuItem(C.CMenuPtr(menu.PlatformPtr), C.int(index))
}

func (d *driver) MenuDispose(menu *webapp.Menu) {
	p := C.CMenuPtr(menu.PlatformPtr)
	C.disposeMenu(p)
	delete(d.menus, p)
}

func (d *driver) MenuItemInitSeparator(item *webapp.MenuItem) {
	mi := C.newMenuSeparator()
	item.PlatformPtr = unsafe.Pointer(mi)
	d.menuItems[mi] = item
}

func (d *driver) MenuItemInit(item *webapp.MenuItem, kind webapp.MenuItemKind) {
	var keyCodeStr string
	if item.KeyCode() != 0 {
		mapping := keys.MappingForKeyCode(item.KeyCode())
		if mapping.KeyChar != 0 {
			keyCodeStr = strings.ToLower(string(mapping.KeyChar))
		}
	}
	cTitle := C.CString(item.Title())
	cKey := C.CString(keyCodeStr)
	cSelector := C.CString(kind.Selector())
	mi := C.newMenuItem(cTitle, cSelector, cKey, C.int(item.KeyModifiers()), kind == webapp.NormalKind)
	C.free(unsafe.Pointer(cSelector))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cTitle))
	item.PlatformPtr = unsafe.Pointer(mi)
	d.menuItems[mi] = item
}

func (d *driver) MenuItemSubMenu(item *webapp.MenuItem) *webapp.Menu {
	if menu, ok := d.menus[C.subMenu(C.CMenuItemPtr(item.PlatformPtr))]; ok {
		return menu
	}
	return nil
}

func (d *driver) MenuItemSetSubMenu(item *webapp.MenuItem, menu *webapp.Menu) {
	C.setSubMenu(C.CMenuItemPtr(item.PlatformPtr), C.CMenuPtr(menu.PlatformPtr))
}

func (d *driver) MenuItemDispose(item *webapp.MenuItem) {
	p := C.CMenuItemPtr(item.PlatformPtr)
	C.disposeMenuItem(p)
	delete(d.menuItems, p)
}

func (d *driver) Displays() []*webapp.Display {
	var count C.ulong
	ptr := unsafe.Pointer(C.displays(&count))
	displays := (*[1 << 30]C.Display)(ptr)
	result := make([]*webapp.Display, count)
	for i := range result {
		d := &webapp.Display{}
		d.Bounds.X = float64(displays[i].bounds.origin.x)
		d.Bounds.Y = float64(displays[i].bounds.origin.y)
		d.Bounds.Width = float64(displays[i].bounds.size.width)
		d.Bounds.Height = float64(displays[i].bounds.size.height)
		d.UsableBounds.X = float64(displays[i].usableBounds.origin.x)
		d.UsableBounds.Y = float64(displays[i].usableBounds.origin.y)
		d.UsableBounds.Width = float64(displays[i].usableBounds.size.width)
		d.UsableBounds.Height = float64(displays[i].usableBounds.size.height)
		d.IsMain = displays[i].isMain != 0
		result[i] = d
	}
	C.free(ptr)
	return result
}

func (d *driver) KeyWindow() *webapp.Window {
	if window, ok := d.windows[C.getKeyWindow()]; ok {
		return window
	}
	return nil
}

func (d *driver) BringAllWindowsToFront() {
	C.bringAllWindowsToFront()
}

func (d *driver) WindowInit(wnd *webapp.Window, style webapp.StyleMask, bounds geom.Rect, title string) error {
	cTitle := C.CString(title)
	w := C.newWindow(C.int(style), C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height), cTitle)
	C.free(unsafe.Pointer(cTitle))
	wnd.PlatformPtr = unsafe.Pointer(w)
	d.windows[w] = wnd
	return nil
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) cef.WindowHandle {
	return cef.WindowHandle(C.contentView(C.CWindowPtr(wnd.PlatformPtr)))
}

func (d *driver) WindowClose(wnd *webapp.Window) {
	p := C.CWindowPtr(wnd.PlatformPtr)
	C.closeWindow(p)
	delete(d.windows, p)
}

func (d *driver) WindowSetTitle(wnd *webapp.Window, title string) {
	cTitle := C.CString(title)
	C.setWindowTitle(C.CWindowPtr(wnd.PlatformPtr), cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (d *driver) WindowBounds(wnd *webapp.Window) geom.Rect {
	var bounds geom.Rect
	C.getWindowBounds(C.CWindowPtr(wnd.PlatformPtr), (*C.double)(&bounds.X), (*C.double)(&bounds.Y), (*C.double)(&bounds.Width), (*C.double)(&bounds.Height))
	return bounds
}

func (d *driver) WindowContentSize(wnd *webapp.Window) geom.Size {
	var size geom.Size
	C.getWindowContentSize(C.CWindowPtr(wnd.PlatformPtr), (*C.double)(&size.Width), (*C.double)(&size.Height))
	return size
}

func (d *driver) WindowSetBounds(wnd *webapp.Window, bounds geom.Rect) {
	C.setWindowBounds(C.CWindowPtr(wnd.PlatformPtr), C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
}

func (d *driver) WindowToFront(wnd *webapp.Window) {
	C.bringWindowToFront(C.CWindowPtr(wnd.PlatformPtr))
}

func (d *driver) WindowMinimize(wnd *webapp.Window) {
	C.minimizeWindow(C.CWindowPtr(wnd.PlatformPtr))
}

func (d *driver) WindowZoom(wnd *webapp.Window) {
	C.zoomWindow(C.CWindowPtr(wnd.PlatformPtr))
}

//export willFinishStartupCallback
func willFinishStartupCallback() {
	webapp.WillFinishStartupCallback()
}

//export didFinishStartupCallback
func didFinishStartupCallback() {
	webapp.DidFinishStartupCallback()
}

//export willActivateCallback
func willActivateCallback() {
	webapp.WillActivateCallback()
}

//export didActivateCallback
func didActivateCallback() {
	webapp.DidActivateCallback()
}

//export willDeactivateCallback
func willDeactivateCallback() {
	webapp.WillDeactivateCallback()
}

//export didDeactivateCallback
func didDeactivateCallback() {
	webapp.DidDeactivateCallback()
}

//export quitAfterLastWindowClosedCallback
func quitAfterLastWindowClosedCallback() bool {
	return webapp.QuitAfterLastWindowClosedCallback()
}

//export checkQuitCallback
func checkQuitCallback() int {
	return int(webapp.CheckQuitCallback())
}

//export quittingCallback
func quittingCallback() {
	webapp.QuittingCallback()
	cef.QuitMessageLoop()
}

//export dispatchUITaskCallback
func dispatchUITaskCallback(id uint64) {
	uitask.Dispatch(id)
}

//export validateMenuItemCallback
func validateMenuItemCallback(menuItem C.CMenuItemPtr) bool {
	if item, ok := drv.menuItems[menuItem]; ok {
		return item.Validate()
	}
	return true
}

//export handleMenuItemCallback
func handleMenuItemCallback(menuItem C.CMenuItemPtr) {
	if item, ok := drv.menuItems[menuItem]; ok && item.Handler != nil {
		item.Handler()
	}
}

//export windowGainedKey
func windowGainedKey(wnd C.CWindowPtr) {
	if w, ok := drv.windows[wnd]; ok {
		w.GainedFocus()
	}
}

//export windowLostKey
func windowLostKey(wnd C.CWindowPtr) {
	if w, ok := drv.windows[wnd]; ok {
		w.LostFocus()
	}
}

//export windowShouldClose
func windowShouldClose(wnd C.CWindowPtr) bool {
	if w, ok := drv.windows[wnd]; ok {
		return w.MayCloseCallback()
	}
	return true
}

//export windowWillClose
func windowWillClose(wnd C.CWindowPtr) {
	if w, ok := drv.windows[wnd]; ok {
		w.WillCloseCallback()
	}
}
