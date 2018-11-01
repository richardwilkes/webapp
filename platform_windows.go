package webapp

import (
	// #cgo LDFLAGS: -LRelease -lole32 -linterop -Wl,--subsystem,windows
	// #include <stdlib.h>
	// #include "platform_windows.h"
	"C"
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
)

// ----- App section -----

var (
	platform _Ctype_struct__platform
)

func platformStart() {
	C.windowsInit(&platform)
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

//export checkQuitCallback
func checkQuitCallback() int32 {
	return int32(CheckQuitCallback())
}

func platformAttemptQuit() {
	// Look at quit.go to see what callbacks are expected.
}

func platformMayQuitNow(quit bool) {
	// Look at quit.go to see what callbacks are expected.
}

func platformInvoke(id uint64) {
	// See task.go
}

func platformInvokeAfter(id uint64, after time.Duration) {
	// See task.go
}

// ----- Menu section -----

var (
	appBar *MenuBar
)

// Look at menu_item.go for callbacks (Validator & Handler fields) that are expected

type platformMenuBar struct {
}

type platformMenu struct {
	_Ctype_struct__menu
}

type platformMenuItem struct {
	_Ctype_struct__menuItem
}

func platformMenuBarForWindow(wnd *Window) *MenuBar {
	if appBar == nil {
		appBar = &MenuBar{
			bar:     NewMenu(""),
			special: make(map[SpecialMenuType]*Menu),
			global:  false,
		}
		C.windowsPlatformSetMenuBar(&platform, &appBar.bar._Ctype_struct__menu)
	}
	return appBar
}

func (bar *MenuBar) platformSetServicesMenu(menu *Menu) {
	// This is macOS-specific and can be left empty
}

func (bar *MenuBar) platformSetWindowMenu(menu *Menu) {
}

func (bar *MenuBar) platformSetHelpMenu(menu *Menu) {
}

func (bar *MenuBar) platformFillAppMenu(appMenu *Menu) {
}

func (menu *Menu) platformInit() {
	if menu.title == "" {
		C.windowsNewMenuBar(&menu._Ctype_struct__menu)
	} else {
		cTitle := C.CString(menu.title)
		C.windowsNewMenu(&menu._Ctype_struct__menu, cTitle)
		C.free(unsafe.Pointer(cTitle))
	}
}

func (menu *Menu) platformItemCount() int {
	return int(C.windowsMenuGetCount(&menu._Ctype_struct__menu))
}

func (menu *Menu) platformItem(index int) *MenuItem {
	return nil
}

func (menu *Menu) platformInsertItem(item *MenuItem, index int) {
	C.windowsMenuInsertItem(&menu._Ctype_struct__menu, &item._Ctype_struct__menuItem, C.int(index))
}

func (menu *Menu) platformRemove(index int) {
}

func (menu *Menu) platformDispose() {
}

func (item *MenuItem) platformInitMenuSeparator() {
	C.windowsNewMenuItemSeparator(&item._Ctype_struct__menuItem)
}

func (item *MenuItem) platformInitMenuItem(kind MenuItemKind) {
	cTitle := C.CString(item.title)
	C.windowsNewMenuItem(&item._Ctype_struct__menuItem, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (item *MenuItem) platformSubMenu() *Menu {
	return nil
}

func (item *MenuItem) platformSetSubMenu(subMenu *Menu) {
	// DR I may have misunderstood how menus work somewhere, so assign this menu to this menuitem
	C.windowsMenuItemHack(&item._Ctype_struct__menuItem, &subMenu._Ctype_struct__menu)
}

func (item *MenuItem) platformDispose() {
}

// ----- Window section -----

// Look at window.go for callbacks that are expected

type platformWindow struct {
	_Ctype_struct__window
}

func platformBringAllWindowsToFront() {
}

func platformKeyWindow() *Window {
	return nil
}

func (window *Window) platformInit(bounds geom.Rect, url string) {
	cURL := C.CString(url)
	C.windowsNewWindow(&platform, &window._Ctype_struct__window, C.int(bounds.Width), C.int(bounds.Height), cURL)
	C.free(unsafe.Pointer(cURL))
}

func (window *Window) platformClose() {
}

func (window *Window) platformSetTitle(title string) {
	cTitle := C.CString(title)
	C.windowsWindowSetTitle(&window._Ctype_struct__window, cTitle)
	C.free(unsafe.Pointer(cTitle))
}

func (window *Window) platformBounds() geom.Rect {
	return geom.Rect{}
}

func (window *Window) platformSetBounds(bounds geom.Rect) {
}

func (window *Window) platformToFront() {
}

func (window *Window) platformMinimize() {
}

func (window *Window) platformZoom() {
}

// ----- Display section -----

func platformDisplays() []*Display {
	return nil
}
