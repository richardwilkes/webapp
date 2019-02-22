#import "AppDelegate.h"
#import "_cgo_export.h"

@implementation AppDelegate

- (void)applicationWillFinishLaunching:(NSNotification *)aNotification {
	willFinishStartupCallback();
}

- (void)applicationDidFinishLaunching:(NSNotification *)aNotification {
	didFinishStartupCallback();
}

- (NSApplicationTerminateReply)applicationShouldTerminate:(NSApplication *)sender {
	// The Mac response codes map to the same values we use
	return checkQuitCallback();
}

- (BOOL)applicationShouldTerminateAfterLastWindowClosed:(NSApplication *)theApplication {
	return quitAfterLastWindowClosedCallback();
}

- (void)applicationWillTerminate:(NSNotification *)aNotification {
	quittingCallback();
}

- (void)applicationWillBecomeActive:(NSNotification *)aNotification {
	willActivateCallback();
}

- (void)applicationDidBecomeActive:(NSNotification *)aNotification {
	didActivateCallback();
}

- (void)applicationWillResignActive:(NSNotification *)aNotification {
	willDeactivateCallback();
}

- (void)applicationDidResignActive:(NSNotification *)aNotification {
	didDeactivateCallback();
}

- (void)themeChanged:(NSNotification *)aNotification {
	themeChangedCallback();
}

@end
