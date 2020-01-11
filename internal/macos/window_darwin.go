// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package macos

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
)

import (
	// #import <stdlib.h>
	// #import "windows.h"
	"C"
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

func (d *driver) WindowThemeIsDark(wnd *webapp.Window) bool {
	if w, ok := wnd.PlatformData.(C.CWindowPtr); ok {
		return C.themeIsDark(w) != 0
	}
	return false
}

//nolint:gocritic
//export windowGainedKey
func windowGainedKey(wnd C.CWindowPtr) {
	if w, ok := drv.windows[wnd]; ok {
		w.GainedFocus()
	}
}

//nolint:gocritic
//export windowLostKey
func windowLostKey(wnd C.CWindowPtr) {
	if w, ok := drv.windows[wnd]; ok {
		w.LostFocus()
	}
}

//nolint:gocritic
//export windowShouldClose
func windowShouldClose(wnd C.CWindowPtr) bool {
	if w, ok := drv.windows[wnd]; ok {
		return w.MayCloseCallback()
	}
	return true
}

//nolint:gocritic
//export windowWillClose
func windowWillClose(wnd C.CWindowPtr) {
	if w, ok := drv.windows[wnd]; ok {
		w.WillCloseCallback()
	}
}
