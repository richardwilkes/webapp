// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#import "WindowDelegate.h"
#import "_cgo_export.h"

@implementation WindowDelegate

-(void)windowDidBecomeKey:(NSNotification *)notification {
	windowGainedKey((CWindowPtr)[notification object]);
}

-(void)windowDidResignKey:(NSNotification *)notification {
	windowLostKey((CWindowPtr)[notification object]);
}

-(BOOL)windowShouldClose:(id)sender {
	return (BOOL)windowShouldClose((CWindowPtr)sender);
}

-(void)windowWillClose:(NSNotification *)notification {
	windowWillClose((CWindowPtr)[notification object]);
}

@end
