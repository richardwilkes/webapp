package macos

import (
	// // Note: Only one file per package needs the #cgo directives.
	// //       Imports are still needed on a per-file basis.
	//
	// #cgo CFLAGS: -x objective-c
	// #cgo LDFLAGS: -framework Cocoa
	// #import "app.h"
	// #import "menus.h"
	// #import "windows.h"
	"C"

	"github.com/richardwilkes/cef/cef"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/webapp"
)

var (
	_ webapp.Driver                  = &driver{}
	_ webapp.AppVisibilityController = &driver{}
)

type driver struct {
	menubar            *webapp.MenuBar
	menus              map[C.CMenuPtr]*webapp.Menu
	menuItemValidators map[int]func() bool
	menuItemHandlers   map[int]func()
	windows            map[C.CWindowPtr]*webapp.Window
}

var drv = &driver{
	menus:              make(map[C.CMenuPtr]*webapp.Menu),
	menuItemValidators: make(map[int]func() bool),
	menuItemHandlers:   make(map[int]func()),
	windows:            make(map[C.CWindowPtr]*webapp.Window),
}

// Driver returns the macOS implementation of the driver.
func Driver() *driver {
	return drv
}

func (d *driver) PrepareForStart() error {
	C.prepareForStart()
	return nil
}

func (d *driver) PrepareForEventLoop() {
	// Nothing to do
}

func (d *driver) RunEventLoop() {
	cef.RunMessageLoop()
	cef.Shutdown()
	atexit.Exit(0)
}

func (d *driver) AttemptQuit() {
	C.attemptQuit()
}

func (d *driver) MayQuitNow(quit bool) {
	var mayQuit C.int
	if quit {
		mayQuit = 1
	}
	C.mayQuitNow(mayQuit)
}

func (d *driver) HideApp() {
	C.hideApp()
}

func (d *driver) HideOtherApps() {
	C.hideOtherApps()
}

func (d *driver) ShowAllApps() {
	C.showAllApps()
}

//export willFinishStartupCallback
func willFinishStartupCallback() {
	webapp.WillFinishStartupCallback()
}

//export didFinishStartupCallback
func didFinishStartupCallback() {
	webapp.DidFinishStartupCallback()
}

//export willActivateCallback
func willActivateCallback() {
	webapp.WillActivateCallback()
}

//export didActivateCallback
func didActivateCallback() {
	webapp.DidActivateCallback()
}

//export willDeactivateCallback
func willDeactivateCallback() {
	webapp.WillDeactivateCallback()
}

//export didDeactivateCallback
func didDeactivateCallback() {
	webapp.DidDeactivateCallback()
}

//export quitAfterLastWindowClosedCallback
func quitAfterLastWindowClosedCallback() bool {
	return webapp.QuitAfterLastWindowClosedCallback()
}

//export checkQuitCallback
func checkQuitCallback() int {
	return int(webapp.CheckQuitCallback())
}

//export quittingCallback
func quittingCallback() {
	webapp.QuittingCallback()
	cef.QuitMessageLoop()
	cef.Shutdown()
	jot.Flush()
}
