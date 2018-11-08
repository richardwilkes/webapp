package webapp

import (
	"syscall"
	"time"
	"unsafe"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xmath/geom"
)

const (
	winCS_VREDRAW    = 1
	winCS_HREDRAW    = 2
	winWM_DESTROY    = 2
	winWM_CLOSE      = 16
	winIDC_ARROW     = 32512
	winCW_USEDEFAULT = 0x80000000
)

// From https://docs.microsoft.com/en-us/windows/desktop/winmsg/window-styles
const (
	winWS_OVERLAPPED       = 0x00000000
	winWS_MAXIMIZEBOX      = 0x00010000 // Has maximize control
	winWS_MINIMIZEBOX      = 0x00020000 // Has minimize control
	winWS_THICKFRAME       = 0x00040000 // Has a sizing border
	winWS_SYSMENU          = 0x00080000 // Has a window menu on title bar
	winWS_HSCROLL          = 0x00100000 // Has a horizontal scroll bar
	winWS_VSCROLL          = 0x00200000 // Has a vertical scroll bar
	winWS_DLGFRAME         = 0x00400000 // Dialog border
	winWS_BORDER           = 0x00800000 // Thin line border
	winWS_CAPTION          = 0x00C00000 // Has title bar
	winWS_MAXIMIZE         = 0x01000000 // Initially maximized
	winWS_CLIPCHILDREN     = 0x02000000
	winWS_CLIPSIBLINGS     = 0x04000000
	winWS_DISABLED         = 0x08000000
	winWS_VISIBLE          = 0x10000000 // Initially visible
	winWS_MINIMIZE         = 0x20000000 // Initially minimized
	winWS_CHILDWINDOW      = 0x40000000
	winWS_POPUP            = 0x80000000
	winWS_OVERLAPPEDWINDOW = winWS_OVERLAPPED | winWS_CAPTION | winWS_SYSMENU | winWS_SYSMENU | winWS_THICKFRAME | winWS_MINIMIZEBOX | winWS_MAXIMIZEBOX
	winWS_POPUPWINDOW      = winWS_POPUP | winWS_BORDER | winWS_SYSMENU
)

var (
	kernel32dll          = syscall.NewLazyDLL("kernel32.dll")
	user32dll            = syscall.NewLazyDLL("user32.dll")
	createWindowExWProc  = user32dll.NewProc("CreateWindowExW")
	defWindowProcWProc   = user32dll.NewProc("DefWindowProcW")
	destroyWindowProc    = user32dll.NewProc("DestroyWindow")
	getModuleHandleWProc = kernel32dll.NewProc("GetModuleHandleW")
	loadCursorWProc      = user32dll.NewProc("LoadCursorW")
	postQuitMessageProc  = user32dll.NewProc("PostQuitMessage")
	registerClassExWProc = user32dll.NewProc("RegisterClassExW")
	windowClassName      = syscall.StringToUTF16Ptr("wndClass")
)

type wndClassExW struct {
	size       uint32
	style      uint32
	wndProc    uintptr
	clsExtra   int32
	wndExtra   int32
	instance   syscall.Handle
	icon       syscall.Handle
	cursor     syscall.Handle
	background syscall.Handle
	menuName   *uint16
	className  *uint16
	iconSm     syscall.Handle
}

// ----- App section -----

func getModuleHandle() (syscall.Handle, error) {
	h, _, err := getModuleHandleWProc.Call(0)
	if h == 0 {
		return 0, err
	}
	return syscall.Handle(h), nil
}

func registerClassEx(wcx *wndClassExW) (uint16, error) {
	h, _, err := registerClassExWProc.Call(uintptr(unsafe.Pointer(wcx)))
	if h == 0 {
		return 0, err
	}
	return uint16(h), nil
}

func loadCursorResource(cursorName uint32) (syscall.Handle, error) {
	h, _, err := loadCursorWProc.Call(uintptr(0), uintptr(uint16(cursorName)))
	if h == 0 {
		return 0, err
	}
	return syscall.Handle(h), nil
}

func platformPrepareForStart() {
	instance, err := getModuleHandle()
	jot.FatalIfErr(err)
	cursor, err := loadCursorResource(winIDC_ARROW)
	jot.FatalIfErr(err)
	wcx := wndClassExW{
		style:    winCS_HREDRAW | winCS_VREDRAW,
		wndProc:  syscall.NewCallback(wndProc),
		instance: instance,
		// icon: LoadIcon(hInstance, MAKEINTRESOURCE(IDI_CEFCLIENT)),
		cursor: cursor,
		// background: cCOLOR_WINDOW + 1,
		// menuName: MAKEINTRESOURCE(IDC_CEFCLIENT),
		className: windowClassName,
		// iconSm: LoadIcon(wcex.hInstance, MAKEINTRESOURCE(IDI_SMALL)),
	}
	wcx.size = uint32(unsafe.Sizeof(wcx))
	_, err = registerClassEx(&wcx)
	jot.FatalIfErr(err)
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

func defWindowProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	ret, _, _ := defWindowProcWProc.Call(uintptr(hwnd), uintptr(msg), uintptr(wparam), uintptr(lparam))
	return uintptr(ret)
}

func createWindow(windowName string, style uint32, x, y, width, height int32, parent, menu, instance syscall.Handle) (syscall.Handle, error) {
	h, _, err := createWindowExWProc.Call(
		0, // ex style
		uintptr(unsafe.Pointer(windowClassName)),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(windowName))),
		uintptr(style),
		uintptr(x),
		uintptr(y),
		uintptr(width),
		uintptr(height),
		uintptr(parent),
		uintptr(menu),
		uintptr(instance),
		uintptr(0), // param -- can attach data here
	)
	if h == 0 {
		return 0, err
	}
	return syscall.Handle(h), nil
}

func destroyWindow(hwnd syscall.Handle) error {
	h, _, err := destroyWindowProc.Call(uintptr(hwnd))
	if h == 0 {
		return err
	}
	return nil
}

func postQuitMessage(exitCode int32) {
	postQuitMessageProc.Call(uintptr(exitCode))
}

func wndProc(hwnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case winWM_CLOSE:
		destroyWindow(hwnd)
	case winWM_DESTROY:
		postQuitMessage(0)
	default:
		return defWindowProc(hwnd, msg, wparam, lparam)
	}
	return 0
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
