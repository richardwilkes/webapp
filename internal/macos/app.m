#import "AppDelegate.h"
#import "MainApp.h"
#import "app.h"

void prepareForStart() {
	[WebApplication sharedApplication];
	[NSApp setDelegate:[AppDelegate new]];
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
