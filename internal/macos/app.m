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
