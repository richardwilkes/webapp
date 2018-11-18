package cef

import (
	// #include "browser.h"
	"C"
)

type Browser struct {
	native *C.cef_browser_t
}

// Host retrieves the BrowserHost.
func (b *Browser) Host() *BrowserHost {
	return &BrowserHost{native: C.gocef_get_browser_host(b.native)}
}

// FocusedFrame returns the currently focused frame.
func (b *Browser) FocusedFrame() *Frame {
	if f := C.gocef_get_focused_frame(b.native); f != nil {
		return &Frame{native: f}
	}
	return nil
}
