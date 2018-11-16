#import "driver.h"
#import "AppDelegate.h"
#import "MainApp.h"
#import "KeyableWindow.h"
#import "WindowDelegate.h"
#import "MenuItemDelegate.h"
#import "_cgo_export.h"

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

void setMenuBar(CMenuPtr bar) {
	[NSApp setMainMenu:(NSMenu *)bar];
}

void setServicesMenu(CMenuPtr menu) {
	[NSApp setServicesMenu:(NSMenu *)menu];
}

void setWindowMenu(CMenuPtr menu) {
	[NSApp setWindowsMenu:(NSMenu *)menu];
}

void setHelpMenu(CMenuPtr menu) {
	[NSApp setHelpMenu:(NSMenu *)menu];
}

CMenuPtr newMenu(const char *title) {
	return (CMenuPtr)[[[NSMenu alloc] initWithTitle:[NSString stringWithUTF8String:title]] retain];
}

void disposeMenu(CMenuPtr menu) {
	[(NSMenu *)menu release];
}

int menuItemCount(CMenuPtr menu) {
	return [(NSMenu *)menu numberOfItems];
}

CMenuItemPtr menuItem(CMenuPtr menu, int index) {
	return (index < 0 || index >= menuItemCount(menu)) ?  nil : [(NSMenu *)menu itemAtIndex:index];
}

void insertMenuItem(CMenuPtr menu, CMenuItemPtr item, int index) {
	NSMenu *m = (NSMenu *)menu;
	if (index < 0) {
		index = [m numberOfItems];
	}
	[m insertItem:item atIndex:index];
}

void removeMenuItem(CMenuPtr menu, int index) {
	[(NSMenu *)menu removeItemAtIndex:index];
}

CMenuItemPtr newMenuSeparator() {
	return (CMenuItemPtr)[[NSMenuItem separatorItem] retain];
}

static MenuItemDelegate *menuItemDelegate = nil;

CMenuItemPtr newMenuItem(int tag, const char *title, const char *selector, const char *key, int modifiers, bool needDelegate) {
	NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:[NSString stringWithUTF8String:title] action:NSSelectorFromString([NSString stringWithUTF8String:selector]) keyEquivalent:[NSString stringWithUTF8String:key]] retain];
	item.tag = tag;
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
	[item setKeyEquivalentModifierMask:modifiers << 16];
	if (needDelegate) {
		if (!menuItemDelegate) {
			menuItemDelegate = [MenuItemDelegate new];
		}
		[item setTarget:menuItemDelegate];
	}
	return (CMenuItemPtr)item;
}

CMenuPtr subMenu(CMenuItemPtr item) {
	NSMenuItem *mitem = (NSMenuItem *)item;
	if ([mitem hasSubmenu]) {
		return (CMenuPtr)[mitem submenu];
	}
	return nil;
}

void setSubMenu(CMenuItemPtr item, CMenuPtr subMenu) {
	[(NSMenuItem *)item setSubmenu: (NSMenu *)subMenu];
}

void disposeMenuItem(CMenuItemPtr item) {
	[(NSMenuItem *)item release];
}

Display *displays(unsigned long *qty) {
	[WebApplication sharedApplication];
	NSArray *s = [NSScreen screens];
	unsigned long count = [s count];
	*qty = count;
	Display *d = (Display *)malloc(sizeof(Display) * count);
	for (unsigned int i = 0; i < count; i++) {
		NSScreen *screen = [s objectAtIndex: i];
		CGDirectDisplayID dID = (CGDirectDisplayID)[[[screen deviceDescription] objectForKey:@"NSScreenNumber"] unsignedIntValue];
		d[i].bounds = CGDisplayBounds(dID);
		d[i].usableBounds = [screen visibleFrame];
		d[i].isMain = CGDisplayIsMain(dID);
		CGRect b = [screen frame];
		d[i].usableBounds.origin.y = d[i].bounds.origin.y + (b.origin.y + b.size.height - (d[i].usableBounds.origin.y + d[i].usableBounds.size.height));
	}
	return d;
}

CWindowPtr getKeyWindow() {
	return [NSApp keyWindow];
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
