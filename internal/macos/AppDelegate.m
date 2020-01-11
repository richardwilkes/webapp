// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
