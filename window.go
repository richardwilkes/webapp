package webapp

import (
	"github.com/richardwilkes/cef"
	"github.com/richardwilkes/toolbox/xmath/geom"
)

// StyleMask controls the look and capabilities of a window.
type StyleMask int

// Possible values for the StyleMask.
const (
	TitledWindowMask StyleMask = 1 << iota
	ClosableWindowMask
	MinimizableWindowMask
	ResizableWindowMask
	BorderlessWindowMask = 0
	StdWindowMask        = TitledWindowMask | ClosableWindowMask | MinimizableWindowMask | ResizableWindowMask
)

// Window holds window information.
type Window struct {
	PlatformPtr uintptr
	Browser     *cef.Browser
	style       StyleMask
	title       string
	// MayCloseCallback is called when the user has requested that the window
	// be closed. Return true to permit it, false to cancel the operation.
	// Defaults to always returning true.
	MayCloseCallback func() bool
	// WillCloseCallback is called just prior to the window closing.
	WillCloseCallback func()
	// GainedFocus is called when the keyboard focus is gained on this window.
	GainedFocus func()
	// LostFocus is called when the keyboard focus is lost from this window.
	LostFocus func()
}

var windowList = make([]*Window, 0)

// Count returns the number of windows that are open.
func Count() int {
	return len(windowList)
}

// Windows returns a slice containing the current set of open windows.
func Windows() []*Window {
	list := make([]*Window, 0, len(windowList))
	copy(list, windowList)
	return list
}

// KeyWindow returns the window that currently has the keyboard focus, or nil
// if none of your application's windows has the keyboard focus.
func KeyWindow() *Window {
	return driver.KeyWindow()
}

// AllWindowsToFront attempts to bring all of the application's windows to the
// foreground.
func AllWindowsToFront() {
	driver.BringAllWindowsToFront()
}

// NewWindow creates a new window with a webview as its content.
func NewWindow(style StyleMask, bounds geom.Rect, title, url string, clientProxy cef.ClientProxy, browserSettings *cef.BrowserSettings, requestContext *cef.RequestContext) (*Window, error) {
	window := &Window{
		style:             style,
		MayCloseCallback:  func() bool { return true },
		WillCloseCallback: func() {},
		GainedFocus:       func() {},
		LostFocus:         func() {},
	}
	if err := driver.WindowInit(window, style, bounds, title); err != nil {
		return nil, err
	}
	bounds.Size = window.WindowContentSize()
	bounds.X = 0
	bounds.Y = 0
	if browserSettings == nil {
		browserSettings = cef.NewBrowserSettings()
	}
	window.Browser = cef.BrowserHostCreateBrowserSync(cef.NewWindowInfo(driver.WindowBrowserParent(window), bounds), cef.NewClient(clientProxy), url, browserSettings, requestContext)
	windowList = append(windowList, window)
	return window, nil
}

func (window *Window) String() string {
	return window.title
}

// AttemptClose closes the window if permitted.
func (window *Window) AttemptClose() {
	if window.MayCloseCallback() {
		window.Dispose()
	}
}

// Dispose of the window.
func (window *Window) Dispose() {
	for i, wnd := range windowList {
		if wnd == window {
			copy(windowList[i:], windowList[i+1:])
			count := len(windowList) - 1
			windowList[count] = nil
			windowList = windowList[:count]
			break
		}
	}
	driver.WindowClose(window)
}

// Title returns the title of this window.
func (window *Window) Title() string {
	return window.title
}

// SetTitle sets the title of this window.
func (window *Window) SetTitle(title string) {
	window.title = title
	driver.WindowSetTitle(window, title)
}

// Bounds returns the boundaries in display coordinates of the frame of this
// window (i.e. the area that includes both the content and its border and
// window controls).
func (window *Window) Bounds() geom.Rect {
	return driver.WindowBounds(window)
}

// SetBounds sets the boundaries of the frame of this window.
func (window *Window) SetBounds(bounds geom.Rect) {
	driver.WindowSetBounds(window, bounds)
}

// WindowContentSize returns the size of the windoe's content area.
func (window *Window) WindowContentSize() geom.Size {
	return driver.WindowContentSize(window)
}

// Focused returns true if the window has the current keyboard focus.
func (window *Window) Focused() bool {
	return window == KeyWindow()
}

// ToFront attempts to bring the window to the foreground and give it the
// keyboard focus.
func (window *Window) ToFront() {
	driver.WindowToFront(window)
}

// Minimize performs the platform's minimize function on the window.
func (window *Window) Minimize() {
	driver.WindowMinimize(window)
}

// Zoom performs the platform's zoom function on the window.
func (window *Window) Zoom() {
	driver.WindowZoom(window)
}

// Closable returns true if the window was created with the
// ClosableWindowMask.
func (window *Window) Closable() bool {
	return window.style&ClosableWindowMask != 0
}

// Minimizable returns true if the window was created with the
// MinimizableWindowMask.
func (window *Window) Minimizable() bool {
	return window.style&MinimizableWindowMask != 0
}

// Resizable returns true if the window was created with the
// ResizableWindowMask.
func (window *Window) Resizable() bool {
	return window.style&ResizableWindowMask != 0
}
