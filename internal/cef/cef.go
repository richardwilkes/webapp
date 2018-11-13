package cef

import (
	// #cgo CFLAGS: -I ${SRCDIR}/../../cef
	// #cgo darwin LDFLAGS: -framework Cocoa -F ${SRCDIR}/../../cef/Release -framework "Chromium Embedded Framework"
	// #cgo windows LDFLAGS: -L${SRCDIR}/../../cef/Release -lcef
	// #include "common.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/xmath/geom"
)

// Settings is an alias for the CEF settings type.
type Settings *C.cef_settings_t

// NewSettings creates a new default Settings instance.
func NewSettings() Settings {
	return Settings(C.new_cef_settings())
}

// Client is an alias for the CEF client type.
type Client *C.cef_client_t

// NewClient creates a new default Client instance.
func NewClient() Client {
	return Client(C.new_cef_client())
}

// WindowHandle is an alias for the CEF window handle type.
type WindowHandle C.cef_window_handle_t

// WindowInfo is an alias for the CEF window info type.
type WindowInfo *C.cef_window_info_t

// NewWindowInfo creates a new default WindowInfo instance.
func NewWindowInfo(parent WindowHandle, bounds geom.Rect) WindowInfo {
	return WindowInfo(C.new_cef_window_info((C.cef_window_handle_t)(parent), C.int(bounds.X), C.int(bounds.Y), C.int(bounds.Width), C.int(bounds.Height)))
}

// BrowserSettings is an alias for the CEF browser settings type.
type BrowserSettings *C.cef_browser_settings_t

// NewBrowserSettings creates a new default BrowserSettings instance.
func NewBrowserSettings() BrowserSettings {
	return BrowserSettings(C.new_cef_browser_settings())
}

// String is an alias for the CEF string type.
type String *C.cef_string_t

// NewString creates a new default String instance.
func NewString(str string) String {
	s := C.CString(str)
	cs := C.new_cef_string_from_utf8(s)
	C.free(unsafe.Pointer(s))
	return String(cs)
}

// Browser is an alias for the CEF browser type.
type Browser *C.cef_browser_t

// NewBrowser creates a new Browser instance.
func NewBrowser(info WindowInfo, client Client, url string, settings BrowserSettings) Browser {
	return Browser(C.cef_browser_host_create_browser_sync(info, client, NewString(url), settings, nil))
}

// BrowserHost is an alias for the CEF browser host object.
type BrowserHost *C.cef_browser_host_t

// GetBrowserHost retrieves the BrowserHost.
func GetBrowserHost(browser Browser) BrowserHost {
	return BrowserHost(C.get_cef_browser_host((*C.cef_browser_t)(browser)))
}

// GetWindowHandle returns the WindowHandle for the browser content.
func GetWindowHandle(host BrowserHost) WindowHandle {
	return WindowHandle(C.get_cef_window_handle((*C.cef_browser_host_t)(host)))
}

// Initialize CEF.
func Initialize(settings Settings) error {
	if C.cef_initialize((*C.cef_main_args_t)(C.calloc(1, C.sizeof_struct__cef_main_args_t)), settings, nil, nil) != 1 {
		return errs.New("Unable to initialize CEF")
	}
	return nil
}

// RunMessageLoop runs the application event loop.
func RunMessageLoop() {
	C.cef_run_message_loop()
}

// QuitMessageLoop quits the message loop in preparation for exiting.
func QuitMessageLoop() {
	C.cef_quit_message_loop()
}

// Shutdown CEF and the application.
func Shutdown() {
	C.cef_shutdown()
}

// EnableHighResolutionSupport enables CEF's high-resolution support. This
// should be called before any other CEF function if this support is desired.
func EnableHighResolutionSupport() {
	C.cef_enable_highdpi_support()
}
