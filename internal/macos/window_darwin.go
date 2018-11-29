package macos

import (
	// #import <stdlib.h>
	// #import "windows.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
)

func (d *driver) BringAllWindowsToFront() {
	C.bringAllWindowsToFront()
}

func (d *driver) KeyWindow() *webapp.Window {
	if window, ok := d.windows[C.getKeyWindow()]; ok {
		return window
	}
	return nil
}

func (d *driver) WindowInit(wnd *webapp.Window, style webapp.StyleMask, bounds geom.Rect, title string) error {
	cTitle := C.CString(title)
	w := C.newWindow(C.int(style), C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height), cTitle)
	C.free(unsafe.Pointer(cTitle))
	wnd.PlatformPtr = uintptr(w)
	d.windows[w] = wnd
	return nil
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) unsafe.Pointer {
	return unsafe.Pointer(C.contentView(C.CWindowPtr(wnd.PlatformPtr)))
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
