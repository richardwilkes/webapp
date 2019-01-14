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
	wnd.PlatformData = w
	d.windows[w] = wnd
	return nil
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) unsafe.Pointer {
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		return unsafe.Pointer(C.contentView(w))
	}
	return nil
}

func (d *driver) WindowClose(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		C.closeWindow(w)
		delete(d.windows, w)
	}
}

func (d *driver) WindowSetTitle(wnd *webapp.Window, title string) {
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		cTitle := C.CString(title)
		C.setWindowTitle(w, cTitle)
		C.free(unsafe.Pointer(cTitle))
	}
}

func (d *driver) WindowBounds(wnd *webapp.Window) geom.Rect {
	var bounds geom.Rect
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		C.getWindowBounds(w, (*C.double)(&bounds.X), (*C.double)(&bounds.Y), (*C.double)(&bounds.Width), (*C.double)(&bounds.Height))
	}
	return bounds
}

func (d *driver) WindowContentSize(wnd *webapp.Window) geom.Size {
	var size geom.Size
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		C.getWindowContentSize(w, (*C.double)(&size.Width), (*C.double)(&size.Height))
	}
	return size
}

func (d *driver) WindowSetBounds(wnd *webapp.Window, bounds geom.Rect) {
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		C.setWindowBounds(w, C.double(bounds.X), C.double(bounds.Y), C.double(bounds.Width), C.double(bounds.Height))
	}
}

func (d *driver) WindowToFront(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		C.bringWindowToFront(w)
	}
}

func (d *driver) WindowMinimize(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		C.minimizeWindow(w)
	}
}

func (d *driver) WindowZoom(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		C.zoomWindow(w)
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
