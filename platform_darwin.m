#import "platform_darwin.h"
#import "_cgo_export.h"

// ----- App section -----

@interface appDelegate : NSObject<NSApplicationDelegate>
@end

@implementation appDelegate

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

@end

void prepareForStart() {
	[[NSAutoreleasePool alloc] init];
	[NSApplication sharedApplication];
	// Required for apps without bundle & Info.plist
	[NSApp setActivationPolicy:NSApplicationActivationPolicyRegular];
	[NSApp setDelegate:[appDelegate new]];
	// Required to use 'NSApplicationActivateIgnoringOtherApps' otherwise our windows end up in the background.
	[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
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

void invoke(unsigned long id) {
	dispatch_async_f(dispatch_get_main_queue(), (void *)id, (dispatch_function_t)dispatchUITaskCallback);
}

void invokeAfter(unsigned long id, long afterNanos) {
	dispatch_after_f(dispatch_time(DISPATCH_TIME_NOW, afterNanos), dispatch_get_main_queue(), (void *)id, (dispatch_function_t)dispatchUITaskCallback);
}

// ----- Menu section -----

@interface ItemDelegate : NSObject
@end

@implementation ItemDelegate

- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
	return validateMenuItemCallback(menuItem);
}

- (void)handleMenuItem:(id)sender {
	handleMenuItemCallback(sender);
}

@end

static ItemDelegate *itemDelegate = nil;

void setMenuBar(Menu bar) {
	[NSApp setMainMenu:bar];
}

void setServicesMenu(Menu menu) {
	[NSApp setServicesMenu:menu];
}

void setWindowMenu(Menu menu) {
	[NSApp setWindowsMenu:menu];
}

void setHelpMenu(Menu menu) {
	[NSApp setHelpMenu:menu];
}

Menu newMenu(const char *title) {
	return [[[NSMenu alloc] initWithTitle:[NSString stringWithUTF8String:title]] retain];
}

void disposeMenu(Menu menu) {
	[((NSMenu *)menu) release];
}

int menuItemCount(Menu menu) {
	return [((NSMenu *)menu) numberOfItems];
}

MenuItem menuItem(Menu menu, int index) {
	return (index < 0 || index >= menuItemCount(menu)) ?  nil : [((NSMenu *)menu) itemAtIndex:index];
}

void insertMenuItem(Menu menu, MenuItem item, int index) {
	[((NSMenu *)menu) insertItem:item atIndex:index];
}

void removeMenuItem(Menu menu, int index) {
	[((NSMenu *)menu) removeItemAtIndex:index];
}

MenuItem newMenuItem(const char *title, const char *selector, const char *key, int modifiers, bool needDelegate) {
	NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:[NSString stringWithUTF8String:title] action:NSSelectorFromString([NSString stringWithUTF8String:selector]) keyEquivalent:[NSString stringWithUTF8String:key]] retain];
	// macOS uses the same modifier mask bit order as we do, but it is shifted up by 16 bits
	[item setKeyEquivalentModifierMask:modifiers << 16];
	if (needDelegate) {
		if (!itemDelegate) {
			itemDelegate = [ItemDelegate new];
		}
		[item setTarget:itemDelegate];
	}
	return item;
}

Menu subMenu(MenuItem item) {
	NSMenuItem *mitem = (NSMenuItem *)item;
	if ([mitem hasSubmenu]) {
		return [mitem submenu];
	}
	return nil;
}

void setSubMenu(MenuItem item, Menu subMenu) {
	[((NSMenuItem *)item) setSubmenu: subMenu];
}

MenuItem newMenuSeparator() {
	return [[NSMenuItem separatorItem] retain];
}

void disposeMenuItem(MenuItem item) {
	[((NSMenuItem *)item) release];
}

// ----- Web view section -----

void add_ref_dummyx(cef_base_ref_counted_t *self) {}
int release_dummyx(cef_base_ref_counted_t *self) { return 1; }
int has_one_ref_dummyx(cef_base_ref_counted_t *self) { return 1; }
int has_at_least_one_ref_dummyx(cef_base_ref_counted_t *self) { return 1; }

void init_dummyx_ref_counter(cef_base_ref_counted_t *h) {
	h->add_ref = add_ref_dummyx;
	h->release = release_dummyx;
	h->has_one_ref = has_one_ref_dummyx;
	h->has_at_least_one_ref = has_at_least_one_ref_dummyx;
}

@interface EmbeddedWebView : NSView
{
	cef_browser_t *webview;
}
@end

@implementation EmbeddedWebView

- (instancetype)initWithFrame: (NSRect)frameRect url:(const char *)url {
	self = [super initWithFrame: frameRect];
	cef_window_info_t *info = (cef_window_info_t *)calloc(1, sizeof(cef_window_info_t));
	info->parent_view = self;

	cef_browser_settings_t *settings = (cef_browser_settings_t *)calloc(1, sizeof(cef_browser_settings_t));
	settings->size = sizeof(cef_browser_settings_t);

	cef_client_t *client = (cef_client_t *)calloc(1, sizeof(cef_client_t));
	client->base.size = sizeof(cef_client_t);
	init_dummyx_ref_counter(&client->base);
	cef_string_t *cefURL = (cef_string_t *)calloc(1, sizeof(cef_string_t));
	cef_string_from_utf8(url, strlen(url), cefURL);
	self->webview = cef_browser_host_create_browser_sync(info, client, cefURL, settings, nil);
	return self;
}

@end

// ----- Window section -----

@interface KeyableWindow : NSWindow
@end

@implementation KeyableWindow

-(BOOL)canBecomeKeyWindow {
	return YES;
}

@end

@interface WindowDelegate : NSObject<NSWindowDelegate>
@end

@implementation WindowDelegate

-(void)windowDidBecomeKey:(NSNotification *)notification {
	windowGainedKey((Window)[notification object]);
}

-(void)windowDidResignKey:(NSNotification *)notification {
	windowLostKey((Window)[notification object]);
}

-(BOOL)windowShouldClose:(id)sender {
	return (BOOL)windowShouldClose((Window)sender);
}

-(void)windowWillClose:(NSNotification *)notification {
	windowWillClose((Window)[notification object]);
}

@end

Window getKeyWindow() {
	return (Window)[NSApp keyWindow];
}

void bringAllWindowsToFront() {
	[[NSRunningApplication currentApplication] activateWithOptions:NSApplicationActivateAllWindows | NSApplicationActivateIgnoringOtherApps];
}

Window newWindow(int styleMask, double x, double y, double width, double height, const char *url) {
	// The styleMask bits match those that Mac OS uses
	NSRect contentRect = [NSWindow contentRectForFrameRect:NSMakeRect(x, [[NSScreen mainScreen] visibleFrame].size.height - (y + height), width, height) styleMask:styleMask];
	NSWindow *window = [[KeyableWindow alloc] initWithContentRect:contentRect styleMask:styleMask backing:NSBackingStoreBuffered defer:YES];
	EmbeddedWebView *view = [[EmbeddedWebView alloc] initWithFrame:NSMakeRect(0, 0, 0, 0) url:url];
	[window setContentView:view];
	[window setDelegate: [WindowDelegate new]];
	return (Window)window;
}

void closeWindow(Window window) {
	[((NSWindow *)window) close];
}

void setWindowTitle(Window window, const char *title) {
	[((NSWindow *)window) setTitle:[NSString stringWithUTF8String:title]];
}

void getWindowBounds(Window window, double *x, double *y, double *width, double *height) {
	CGRect frame = [((NSWindow *)window) frame];
	*x = frame.origin.x;
	*y = [[NSScreen mainScreen] visibleFrame].size.height - (frame.origin.y + frame.size.height);
	*width = frame.size.width;
	*height = frame.size.height;
}

void setWindowBounds(Window window, double x, double y, double width, double height) {
	NSWindow *win = (NSWindow *)window;
	CGRect frame = [win frame];
	[win setFrame:NSMakeRect(x, [[NSScreen mainScreen] visibleFrame].size.height - (y + height), width, height) display:YES];
}

void bringWindowToFront(Window window) {
	[((NSWindow *)window) makeKeyAndOrderFront:nil];
}

void minimizeWindow(Window window) {
	[((NSWindow *)window) performMiniaturize:nil];
}

void zoomWindow(Window window) {
	[((NSWindow *)window) performZoom:nil];
}

// ----- Displays section -----

Display *displays(unsigned long *qty) {
	[NSApplication sharedApplication];
	NSArray *s = [NSScreen screens];
	unsigned long count = [s count];
	*qty = count;
	Display *d = (Display *)malloc(sizeof(Display) * count);
	for (unsigned int i = 0; i < count; i++) {
		NSScreen *screen = [s objectAtIndex: i];
		d[i].scaleFactor = [screen backingScaleFactor];
		CGDirectDisplayID dID = (CGDirectDisplayID)[[[screen deviceDescription] objectForKey:@"NSScreenNumber"] unsignedIntValue];
		d[i].bounds = CGDisplayBounds(dID);
		d[i].usableBounds = [screen visibleFrame];
		d[i].isMain = CGDisplayIsMain(dID);
		CGRect b = [screen frame];
		d[i].usableBounds.origin.y = d[i].bounds.origin.y + (b.origin.y + b.size.height - (d[i].usableBounds.origin.y + d[i].usableBounds.size.height));
	}
	return d;
}
