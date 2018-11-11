package webapp

// Possible termination responses
const (
	Cancel QuitResponse = iota
	Now
	Later // Must make a call to MayQuitNow() at some point in the future.
)

// QuitResponse is used to respond to requests for app termination.
type QuitResponse int

// QuitAfterLastWindowClosedCallback is called when the last window is closed
// to determine if the application should quit as a result. Default returns
// yes.
var QuitAfterLastWindowClosedCallback = func() bool { return true }

// AttemptQuit initiates the termination sequence.
func AttemptQuit() {
	driver.AttemptQuit()
}

// CheckQuitCallback is called when termination has been requested. Default
// returns Now.
var CheckQuitCallback = func() QuitResponse { return Now }

// MayQuitNow resumes the termination sequence that was delayed when Later is
// returned from CheckQuit. Passing in false for the quit parameter will
// cancel the termination sequence while true will allow it to proceed.
func MayQuitNow(quit bool) {
	driver.MayQuitNow(quit)
}

// QuittingCallback is called when the app will in fact terminate.
var QuittingCallback = func() {}
