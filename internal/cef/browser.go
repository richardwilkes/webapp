package cef

import (
	// #include "browser.h"
	"C"
)

type Browser struct {
	native *C.cef_browser_t
}

// Host retrieves the BrowserHost.
func (b *Browser) Host() BrowserHost {
	return BrowserHost(C.get_cef_browser_host(b.native))
}

func (b *Browser) Cut() {
	C.cef_browser_cut(b.native)
}

func (b *Browser) Copy() {
	C.cef_browser_copy(b.native)
}

func (b *Browser) Paste() {
	C.cef_browser_paste(b.native)
}

func (b *Browser) Delete() {
	C.cef_browser_delete(b.native)
}

func (b *Browser) SelectAll() {
	C.cef_browser_select_all(b.native)
}
