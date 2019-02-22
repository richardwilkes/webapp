package webapp

import (
	"runtime"

	"github.com/richardwilkes/cef/cef"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/errs"
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
func Initialize(platformDriver Driver) (*cef.MainArgs, error) {
	if platformDriver == nil {
		return nil, errs.New("platformDriver may not be nil")
	}
	driver = platformDriver
	runtime.LockOSThread()
	cef.EnableHighdpiSupport()
	args := cef.NewMainArgs()
	if code := cef.ExecuteProcess(args, nil, nil); code >= 0 {
		atexit.Exit(int(code))
	}
	return args, nil
}

// Start the user interface. This and most other functions in this package
// should only be called from the main (UI) thread.
func Start(args *cef.MainArgs, settings *cef.Settings, application *cef.App) error {
	if driver == nil {
		return errs.New("webapp.Initialize(driver) must be called first")
	}
	cef.InstantiateApplication()
	if err := driver.PrepareForStart(); err != nil {
		return err
	}
	if settings == nil {
		settings = cef.NewSettings()
	}
	if cef.Initialize(args, settings, application, nil) == 0 {
		return errs.New("Unable to initialize CEF")
	}
	driver.PrepareForEventLoop()
	driver.RunEventLoop()
	return nil // Never reaches here
}

// PlatformDriver returns the underlying driver.
func PlatformDriver() Driver {
	return driver
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

// ThemeChangedCallback is called when the theme is changed.
var ThemeChangedCallback = func() {}
