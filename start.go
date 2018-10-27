package webapp

import (
	"runtime"
)

// Start the user interface. Locks the calling goroutine to its current OS
// thread. Does not return. This should be called as early as possible in your
// main(), before any goroutines are started.
func Start() {
	runtime.LockOSThread()
	platformStart()
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
