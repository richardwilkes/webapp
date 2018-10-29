package webapp

import (
	"time"

	"github.com/richardwilkes/toolbox/xmath/geom"
)

// ----- App section -----

func platformStart() {
	// Should not return. Look at start.go to see what callbacks are expected.
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

// Look at menu_item.go for callbacks (Validator & Handler fields) that are expected

type platformMenuBar struct {
}

type platformMenu struct {
}

type platformMenuItem struct {
}

func platformMenuBarForWindow(wnd *Window) *MenuBar {
	return nil
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
}

func (menu *Menu) platformItemCount() int {
	return 0
}

func (menu *Menu) platformItem(index int) *MenuItem {
	return nil
}

func (menu *Menu) platformInsertItem(item *MenuItem, index int) {
}

func (menu *Menu) platformRemove(index int) {
}

func (menu *Menu) platformDispose() {
}

func (item *MenuItem) platformInitMenuSeparator() {
}

func (item *MenuItem) platformInitMenuItem(kind MenuItemKind) {
}

func (item *MenuItem) platformSubMenu() *Menu {
	return nil
}

func (item *MenuItem) platformSetSubMenu(subMenu *Menu) {
}

func (item *MenuItem) platformDispose() {
}

// ----- Window section -----

// Look at window.go for callbacks that are expected

type platformWindow struct {
}

func platformBringAllWindowsToFront() {
}

func platformKeyWindow() *Window {
	return nil
}

func (window *Window) platformInit(bounds geom.Rect, url string) {
}

func (window *Window) platformClose() {
}

func (window *Window) platformSetTitle(title string) {
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
