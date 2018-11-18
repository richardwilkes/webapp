package cef

import (
	// // Note: Only one file per package needs the #cgo directives.
	// //       Includes are still needed on a per-file basis.
	//
	// #cgo CFLAGS: -I ${SRCDIR}/../../cef
	// #cgo darwin LDFLAGS: -framework Cocoa -F ${SRCDIR}/../../cef/Release -framework "Chromium Embedded Framework"
	// #cgo windows LDFLAGS: -L${SRCDIR}/../../cef/Release -lcef -Wl,--subsystem,windows
	// #include <stdlib.h>
	// #include "include/capi/cef_app_capi.h"
	// #include "include/capi/cef_client_capi.h"
	// #include "ref.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/errs"
)

// Various CEF type aliases
type (
	Settings        *C.cef_settings_t
	Client          *C.cef_client_t
	BrowserSettings *C.cef_browser_settings_t
)

// NewSettings creates a new default Settings instance.
func NewSettings() Settings {
	settings := (*C.cef_settings_t)(C.calloc(1, C.sizeof_struct__cef_settings_t))
	settings.size = C.sizeof_struct__cef_settings_t
	settings.no_sandbox = 1
	settings.command_line_args_disabled = 1
	return Settings(settings)
}

// NewClient creates a new default Client instance.
func NewClient() Client {
	return Client(unsafe.Pointer(C.gocef_refcnt_alloc(C.sizeof_struct__cef_client_t)))
}

// NewBrowserSettings creates a new default BrowserSettings instance.
func NewBrowserSettings() BrowserSettings {
	settings := (*C.cef_browser_settings_t)(C.calloc(1, C.sizeof_struct__cef_browser_settings_t))
	settings.size = C.sizeof_struct__cef_browser_settings_t
	return BrowserSettings(settings)
}

// NewBrowser creates a new Browser instance.
func NewBrowser(info *WindowInfo, client Client, url string, settings BrowserSettings) *Browser {
	return &Browser{
		native: C.cef_browser_host_create_browser_sync(info.native, client, newCEFStr(url), settings, nil),
	}
}

// ExecuteProcess is used to start the secondary CEF processes. If this is
// the main process, this call will do nothing and return. If it is a
// secondary process, the call will not return.
func ExecuteProcess() error {
	args, err := mainArgs()
	if err != nil {
		return err
	}
	if code := C.cef_execute_process(args, nil, nil); code >= 0 {
		atexit.Exit(int(code))
	}
	return nil
}

// Initialize CEF.
func Initialize(settings Settings) error {
	args, err := mainArgs()
	if err != nil {
		return err
	}
	if C.cef_initialize(args, settings, nil, nil) != 1 {
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
