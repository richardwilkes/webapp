package webapp

import (
	"runtime"

	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/webapp/internal/cef"
)

var driver Driver

// Initialize the library. This should be called on the main OS thread as
// early as possible -- typically before parsing command-line arguments, as
// this may be a request to start a sub-process and not the primary
// application.
//
// Note that starting goroutines may cause your code to start executing on a
// secondary thread, so either make this call prior to starting goroutines, or
// explicitly call runtime.LockOSThread() before doing so.
//
// This function will not return if the executable is being started as a CEF
// sub-process.
func Initialize(platformDriver Driver) error {
	if platformDriver == nil {
		return errs.New("platformDriver may not be nil")
	}
	driver = platformDriver
	runtime.LockOSThread()
	cef.EnableHighResolutionSupport()
	return driver.Initialize()
}

// Start the user interface. This and most other functions in this package
// should only be called from the main (UI) thread.
func Start() error {
	if driver == nil {
		return errs.New("webapp.Initialize(driver) must be called first")
	}
	if err := driver.PrepareForStart(); err != nil {
		return err
	}
	if err := cef.Initialize(cef.NewSettings()); err != nil {
		return err
	}
	driver.PrepareForEventLoop()
	cef.RunMessageLoop()
	cef.Shutdown()
	atexit.Exit(0)
	return nil
}

// WillFinishStartupCallback is called right before application startup has
// completed. This is a good point to create any windows your app wants to
// display.
var WillFinishStartupCallback = func() {}

// DidFinishStartupCallback is called once application startup has completed.
var DidFinishStartupCallback = func() {}

// WillActivateCallback is called right before the application is
// activated.
var WillActivateCallback = func() {}

// DidActivateCallback is called once the application is activated.
var DidActivateCallback = func() {}

// WillDeactivateCallback is called right before the application is
// deactivated.
var WillDeactivateCallback = func() {}

// DidDeactivateCallback is called once the application is deactivated.
var DidDeactivateCallback = func() {}
