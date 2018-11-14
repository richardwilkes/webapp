package cef

import (
	// #cgo CFLAGS: -I ${SRCDIR}/../../cef
	// #cgo darwin LDFLAGS: -framework Cocoa -F ${SRCDIR}/../../cef/Release -framework "Chromium Embedded Framework"
	// #cgo windows LDFLAGS: -L${SRCDIR}/../../cef/Release -lcef
	// #include "common.h"
	"C"
	"sync"
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/xmath/geom"
)

var (
	nextTaskID     C.int = 1
	nextTaskIDLock sync.Mutex
	taskMap        = make(map[C.int]func())
)

// Various CEF type aliases
type (
	String          *C.cef_string_t
	Settings        *C.cef_settings_t
	Client          *C.cef_client_t
	WindowHandle    C.cef_window_handle_t
	WindowInfo      *C.cef_window_info_t
	BrowserSettings *C.cef_browser_settings_t
	Browser         *C.cef_browser_t
	BrowserHost     *C.cef_browser_host_t
)

// NewSettings creates a new default Settings instance.
func NewSettings() Settings {
	return Settings(C.new_cef_settings())
}

// NewClient creates a new default Client instance.
func NewClient() Client {
	return Client(C.new_cef_client())
}

// NewWindowInfo creates a new default WindowInfo instance.
func NewWindowInfo(parent WindowHandle, bounds geom.Rect) WindowInfo {
	return WindowInfo(C.new_cef_window_info((C.cef_window_handle_t)(parent), C.int(bounds.X), C.int(bounds.Y), C.int(bounds.Width), C.int(bounds.Height)))
}

// NewBrowserSettings creates a new default BrowserSettings instance.
func NewBrowserSettings() BrowserSettings {
	return BrowserSettings(C.new_cef_browser_settings())
}

// NewString creates a new default String instance.
func NewString(str string) String {
	s := C.CString(str)
	cs := C.new_cef_string_from_utf8(s)
	C.free(unsafe.Pointer(s))
	return String(cs)
}

// NewBrowser creates a new Browser instance.
func NewBrowser(info WindowInfo, client Client, url string, settings BrowserSettings) Browser {
	return Browser(C.cef_browser_host_create_browser_sync(info, client, NewString(url), settings, nil))
}

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

// PostUITask schedules a task for execution on the UI thread.
func PostUITask(task func()) {
	nextTaskIDLock.Lock()
	id := nextTaskID
	nextTaskID++
	taskMap[id] = task
	nextTaskIDLock.Unlock()
	C.cef_post_task(C.TID_UI, C.new_cef_task(id))
}

//export taskCallback
func taskCallback(id C.int) {
	nextTaskIDLock.Lock()
	task := taskMap[id]
	delete(taskMap, id)
	nextTaskIDLock.Unlock()
	if task != nil {
		task()
	}
}
