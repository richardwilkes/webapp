// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

#import "AppDelegate.h"
#import "app.h"

void prepareForStart() {
	AppDelegate *delegate = [AppDelegate new];
	[NSApp setDelegate:delegate];
	[NSDistributedNotificationCenter.defaultCenter addObserver:delegate selector:@selector(themeChanged:) name:@"AppleInterfaceThemeChangedNotification" object: nil];
}

void attemptQuit() {
	[NSApp terminate:nil];
}

void mayQuitNow(int quit) {
	[NSApp replyToApplicationShouldTerminate:quit];
}

void hideApp() {
	[NSApp hide:nil];
}

void hideOtherApps() {
	[NSApp hideOtherApplications:NSApp];
}

void showAllApps() {
	[NSApp unhideAllApplications:NSApp];
}
