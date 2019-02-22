#import "KeyableWindow.h"
#import "WindowDelegate.h"
#import "windows.h"

CWindowPtr getKeyWindow() {
	CWindowPtr wnd = [NSApp keyWindow];
	if (!wnd) {
		wnd = [NSApp mainWindow];
	}
	return wnd;
}

void bringAllWindowsToFront() {
	[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
}

CWindowPtr newWindow(int styleMask, double x, double y, double width, double height, const char *title) {
	// The styleMask bits match those that Mac OS uses
	NSRect contentRect = [NSWindow contentRectForFrameRect:NSMakeRect(x, [[NSScreen mainScreen] visibleFrame].size.height - (y + height), width, height) styleMask:styleMask];
	KeyableWindow *window = [[KeyableWindow alloc] initWithContentRect:contentRect styleMask:styleMask backing:NSBackingStoreBuffered defer:YES];
	[window setDelegate: [WindowDelegate new]];
	[window setTitle:[NSString stringWithUTF8String:title]];
	return (CWindowPtr)window;
}

CViewPtr contentView(CWindowPtr wnd) {
	return (CViewPtr)[(NSWindow *)wnd contentView];
}

void closeWindow(CWindowPtr wnd) {
	[(NSWindow *)wnd close];
}

void setWindowTitle(CWindowPtr wnd, const char *title) {
	[(NSWindow *)wnd setTitle:[NSString stringWithUTF8String:title]];
}

void getWindowBounds(CWindowPtr wnd, double *x, double *y, double *width, double *height) {
	CGRect frame = [(NSWindow *)wnd frame];
	*x = frame.origin.x;
	*y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	*width = frame.size.width;
	*height = frame.size.height;
}

void setWindowBounds(CWindowPtr wnd, double x, double y, double width, double height) {
	NSWindow *win = (NSWindow *)wnd;
	CGRect frame = [win frame];
	[win setFrame:NSMakeRect(x, [[NSScreen mainScreen] visibleFrame].size.height - (y + height), width, height) display:YES];
}

void getWindowContentSize(CWindowPtr wnd, double *width, double *height) {
	CGRect frame = [[(NSWindow *)wnd contentView] frame];
	*width = frame.size.width;
	*height = frame.size.height;
}

void bringWindowToFront(CWindowPtr wnd) {
	[(NSWindow *)wnd makeKeyAndOrderFront:nil];
}

void minimizeWindow(CWindowPtr wnd) {
	[(NSWindow *)wnd performMiniaturize:nil];
}

void zoomWindow(CWindowPtr wnd) {
	[(NSWindow *)wnd performZoom:nil];
}

int themeIsDark(CWindowPtr wnd) {
	if (@available(macOS 10.14, *)) {
		NSAppearanceName basicAppearance = [[(NSWindow *)wnd contentView].effectiveAppearance bestMatchFromAppearancesWithNames:@[ NSAppearanceNameAqua, NSAppearanceNameDarkAqua ]];
		return [basicAppearance isEqualToString:NSAppearanceNameDarkAqua];
	} else {
		return 0;
	}
}
