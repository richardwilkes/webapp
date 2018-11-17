#include "browser.h"

cef_browser_host_t *get_cef_browser_host(cef_browser_t *browser) {
	return browser->get_host(browser);
}

void cef_browser_cut(cef_browser_t *browser) {
	cef_frame_t *f = browser->get_focused_frame(browser);
	if (f) {
		f->cut(f);
	}
}

void cef_browser_copy(cef_browser_t *browser) {
	cef_frame_t *f = browser->get_focused_frame(browser);
	if (f) {
		f->copy(f);
	}
}

void cef_browser_paste(cef_browser_t *browser) {
	cef_frame_t *f = browser->get_focused_frame(browser);
	if (f) {
		f->paste(f);
	}
}

void cef_browser_delete(cef_browser_t *browser) {
	cef_frame_t *f = browser->get_focused_frame(browser);
	if (f) {
		f->del(f);
	}
}

void cef_browser_select_all(cef_browser_t *browser) {
	cef_frame_t *f = browser->get_focused_frame(browser);
	if (f) {
		f->select_all(f);
	}
}
