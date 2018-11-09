package webapp

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp/internal/windows"
	"github.com/richardwilkes/webapp/internal/windows/constants/cs"
)

const (
	winIDC_ARROW     = 32512
	winCW_USEDEFAULT = 0x80000000
)

// ----- App section -----

func platformPrepareForStart() error {
	instance, err := windows.GetModuleHandleW()
	if err != nil {
		return err
	}
	cursor, err := windows.LoadCursorW(winIDC_ARROW)
	if err != nil {
		return err
	}
	cnstr, err := syscall.UTF16PtrFromString("wndClass")
	if err != nil {
		return errs.NewWithCause("Unable to convert className to UTF16", err)
	}
	wcx := windows.WNDCLASSEXW{
		Style:    cs.HREDRAW | cs.VREDRAW,
		WndProc:  syscall.NewCallback(windows.WndProc),
		Instance: instance,
		// Icon: LoadIcon(hInstance, MAKEINTRESOURCE(IDI_CEFCLIENT)),
		Cursor: cursor,
		// Background: cCOLOR_WINDOW + 1,
		// MenuName: MAKEINTRESOURCE(IDC_CEFCLIENT),
		ClassName: cnstr,
		// IconSm: LoadIcon(wcex.hInstance, MAKEINTRESOURCE(IDI_SMALL)),
	}
	wcx.Size = uint32(unsafe.Sizeof(wcx))
	_, err = windows.RegisterClassExW(&wcx)
	return err
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
}

func platformMayQuitNow(quit bool) {
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
