package webapp

import (
	"os"
	"runtime"

	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/webapp/internal/cef"
)

var driver Driver

// Start the user interface. This should only be called on the main OS thread.
// Starting goroutines may cause your code to start executing on a secondary
// thread, so either make this call prior to starting goroutines, or
// explicitly call runtime.LockOSThread() as the first thing in main().
// Only returns if an error occurs during initialization. No other functions
// within this package should be called before this method.
func Start(platformDriver Driver) error {
	if platformDriver == nil {
		return errs.New("platformDriver may not be nil")
	}
	driver = platformDriver
	runtime.LockOSThread()
	if err := driver.PrepareForStart(); err != nil {
		return err
	}
	if err := cef.Initialize(cef.NewMainArgs(os.Args), cef.NewSettings()); err != nil {
		return err
	}
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
