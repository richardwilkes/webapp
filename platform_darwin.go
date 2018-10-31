package webapp

import (
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa -framework WebKit -framework Security
	// #include "platform_darwin.h"
	"C"
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/i18n"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp/keys"
)

// ----- App section -----

func platformStart() {
	C.start()
}

//export willFinishStartupCallback
func willFinishStartupCallback() {
	WillFinishStartupCallback()
}

//export didFinishStartupCallback
func didFinishStartupCallback() {
	DidFinishStartupCallback()
}

//export willActivateCallback
func willActivateCallback() {
	WillActivateCallback()
}

//export didActivateCallback
func didActivateCallback() {
	DidActivateCallback()
}

//export willDeactivateCallback
func willDeactivateCallback() {
	WillDeactivateCallback()
}

//export didDeactivateCallback
func didDeactivateCallback() {
	DidDeactivateCallback()
}

//export quitAfterLastWindowClosedCallback
func quitAfterLastWindowClosedCallback() bool {
	return QuitAfterLastWindowClosedCallback()
}

func platformAttemptQuit() {
	C.attemptQuit()
}

//export checkQuitCallback
func checkQuitCallback() int {
	return int(CheckQuitCallback())
}

func platformMayQuitNow(quit bool) {
	var mayQuit C.int
	if quit {
		mayQuit = 1
	}
	C.mayQuitNow(mayQuit)
}

//export quittingCallback
func quittingCallback() {
	QuittingCallback()
	atexit.Exit(0)
}

func platformInvoke(id uint64) {
	C.invoke(C.ulong(id))
}

func platformInvokeAfter(id uint64, after time.Duration) {
	C.invokeAfter(C.ulong(id), C.long(after.Nanoseconds()))
}

//export dispatchUITaskCallback
func dispatchUITaskCallback(id uint64) {
	dispatchUITask(id)
}

// ----- Menu section -----

var (
	appBar      *MenuBar
	menuMap     = make(map[C.Menu]*Menu)
	menuItemMap = make(map[C.MenuItem]*MenuItem)
)

type platformMenuBar struct {
}

type platformMenu struct {
	menu C.Menu
}

type platformMenuItem struct {
	item C.MenuItem
}

func platformMenuBarForWindow(_ *Window) *MenuBar {
	if appBar == nil {
		appBar = &MenuBar{
			platformMenuBar: platformMenuBar{}, // Here solely to shut the linter up
			bar:             NewMenu(""),
			special:         make(map[SpecialMenuType]*Menu),
			global:          true,
		}
		C.setMenuBar(appBar.bar.menu)
	}
	return appBar
}

func (bar *MenuBar) platformSetServicesMenu(menu *Menu) {
	C.setServicesMenu(menu.menu)
}

func (bar *MenuBar) platformSetWindowMenu(menu *Menu) {
	C.setWindowMenu(menu.menu)
}

func (bar *MenuBar) platformSetHelpMenu(menu *Menu) {
	C.setHelpMenu(menu.menu)
}

func (bar *MenuBar) platformFillAppMenu(appMenu *Menu) {
	appMenu.AppendItem(NewMenuSeparator())
	servicesMenu := NewMenu(i18n.Text("Services"))
	appMenu.AppendMenu(servicesMenu)
	bar.SetupSpecialMenu(ServicesSpecialMenu, servicesMenu)

	appMenu.AppendItem(NewMenuSeparator())
	hideItem := NewMenuItemWithKey(fmt.Sprintf(i18n.Text("Hide %s"), cmdline.AppName), keys.VirtualKeyH)
	hideItem.Handler = func() { C.hideApp() }
	appMenu.AppendItem(hideItem)
	hideOthersItem := NewMenuItemWithKeyAndModifiers(i18n.Text("Hide Others"), keys.VirtualKeyH, keys.OptionModifier|keys.PlatformMenuModifier())
	hideOthersItem.Handler = func() { C.hideOtherApps() }
	appMenu.AppendItem(hideOthersItem)
	showAllItem := NewMenuItem(i18n.Text("Show All"))
	showAllItem.Handler = func() { C.showAllApps() }
	appMenu.AppendItem(showAllItem)
}

func (menu *Menu) platformInit() {
	cTitle := C.CString(menu.title)
	menu.menu = C.newMenu(cTitle)
	menuMap[menu.menu] = menu
	C.free(unsafe.Pointer(cTitle))
}

func (menu *Menu) platformItemCount() int {
	return int(C.menuItemCount(menu.menu))
}

func (menu *Menu) platformItem(index int) *MenuItem {
	return menuItemMap[C.menuItem(menu.menu, C.int(index))]
}

func (menu *Menu) platformInsertItem(item *MenuItem, index int) {
	C.insertMenuItem(menu.menu, item.item, C.int(index))
}

func (menu *Menu) platformRemove(index int) {
	C.removeMenuItem(menu.menu, C.int(index))
}

func (menu *Menu) platformDispose() {
	C.disposeMenu(menu.menu)
	delete(menuMap, menu.menu)
}

func (item *MenuItem) platformInitMenuSeparator() {
	item.item = C.newMenuSeparator()
	menuItemMap[item.item] = item
}

func (item *MenuItem) platformInitMenuItem(kind MenuItemKind) {
	var keyCodeStr string
	if item.keyCode != 0 {
		mapping := keys.MappingForKeyCode(item.keyCode)
		if mapping.KeyChar != 0 {
			keyCodeStr = strings.ToLower(string(mapping.KeyChar))
		}
	}
	cTitle := C.CString(item.title)
	cKey := C.CString(keyCodeStr)
	cSelector := C.CString(kind.selector())
	item.item = C.newMenuItem(cTitle, cSelector, cKey, C.int(item.keyModifiers), kind == NormalKind)
	C.free(unsafe.Pointer(cSelector))
	C.free(unsafe.Pointer(cKey))
	C.free(unsafe.Pointer(cTitle))
	menuItemMap[item.item] = item
}

func (item *MenuItem) platformSubMenu() *Menu {
	if menu, ok := menuMap[C.subMenu(item.item)]; ok {
		return menu
	}
	return nil
}

func (item *MenuItem) platformSetSubMenu(subMenu *Menu) {
	C.setSubMenu(item.item, subMenu.menu)
}

func (item *MenuItem) platformDispose() {
	C.disposeMenuItem(item.item)
	delete(menuItemMap, item.item)
}

//export validateMenuItemCallback
func validateMenuItemCallback(menuItem C.MenuItem) bool {
	if item, ok := menuItemMap[menuItem]; ok && item.Validator != nil {
		item.enabled = item.Validator()
		return item.enabled
	}
	return true
}

//export handleMenuItemCallback
func handleMenuItemCallback(menuItem C.MenuItem) {
	if item, ok := menuItemMap[menuItem]; ok && item.Handler != nil {
		item.Handler()
	}
}

// ----- Window section -----

var (
	windowMap = make(map[C.Window]*Window)
)

type platformWindow struct {
	window C.Window
}

func platformBringAllWindowsToFront() {
	C.bringAllWindowsToFront()
}

func platformKeyWindow() *Window {
	if window, ok := windowMap[C.getKeyWindow()]; ok {
		return window
	}
	return nil
}

func (window *Window) platformInit(bounds geom.Rect, url string) {
	cURL := C.CString(url)
	window.window = C.newWindow(C.int(window.style), C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height), cURL)
	C.free(unsafe.Pointer(cURL))
	windowMap[window.window] = window
}

func (window *Window) platformClose() {
	C.closeWindow(window.window)
	delete(windowMap, window.window)
}

func (window *Window) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.setWindowTitle(window.window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Window) platformBounds() geom.Rect {
	var bounds geom.Rect
	C.getWindowBounds(window.window, (*C.double)(&bounds.X), (*C.double)(&bounds.Y), (*C.double)(&bounds.Width), (*C.double)(&bounds.Height))
	return bounds
}

func (window *Window) platformSetBounds(bounds geom.Rect) {
	C.setWindowBounds(window.window, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
}

func (window *Window) platformToFront() {
	C.bringWindowToFront(window.window)
}

func (window *Window) platformMinimize() {
	C.minimizeWindow(window.window)
}

func (window *Window) platformZoom() {
	C.zoomWindow(window.window)
}

//export windowGainedKey
func windowGainedKey(cWindow C.Window) {
	if window, ok := windowMap[cWindow]; ok {
		window.GainedFocus()
	}
}

//export windowLostKey
func windowLostKey(cWindow C.Window) {
	if window, ok := windowMap[cWindow]; ok {
		window.LostFocus()
	}
}

//export windowShouldClose
func windowShouldClose(cWindow C.Window) bool {
	if window, ok := windowMap[cWindow]; ok {
		return window.MayCloseCallback()
	}
	return true
}

//export windowWillClose
func windowWillClose(cWindow C.Window) {
	if window, ok := windowMap[cWindow]; ok {
		window.WillCloseCallback()
	}
}

// ----- Display section -----

func platformDisplays() []*Display {
	var count C.ulong
	ptr := unsafe.Pointer(C.displays(&count))
	displays := (*[1 << 30]C.Display)(ptr)
	result := make([]*Display, count)
	for i := range result {
		d := &Display{}
		d.Bounds.X = float64(displays[i].bounds.origin.x)
		d.Bounds.Y = float64(displays[i].bounds.origin.y)
		d.Bounds.Width = float64(displays[i].bounds.size.width)
		d.Bounds.Height = float64(displays[i].bounds.size.height)
		d.UsableBounds.X = float64(displays[i].usableBounds.origin.x)
		d.UsableBounds.Y = float64(displays[i].usableBounds.origin.y)
		d.UsableBounds.Width = float64(displays[i].usableBounds.size.width)
		d.UsableBounds.Height = float64(displays[i].usableBounds.size.height)
		d.ScaleFactor = float64(displays[i].scaleFactor)
		d.IsMain = displays[i].isMain != 0
		result[i] = d
	}
	C.free(ptr)
	return result
}
