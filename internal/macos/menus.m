#import "MenuItemDelegate.h"
#import "menus.h"

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

CMenuItemPtr menuItemAtIndex(CMenuPtr menu, int index) {
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

CMenuItemPtr newMenuItem(int id, const char *title, const char *selector, const char *key, int modifiers, bool needDelegate) {
	NSMenuItem *item = [[[NSMenuItem alloc] initWithTitle:[NSString stringWithUTF8String:title] action:NSSelectorFromString([NSString stringWithUTF8String:selector]) keyEquivalent:[NSString stringWithUTF8String:key]] retain];
	[item setTag:id];
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

void setMenuItemTitle(CMenuItemPtr item, const char *title) {
	((NSMenuItem *)item).title = [NSString stringWithUTF8String:title];
}

CMenuItemInfo *menuItemInfo(CMenuItemPtr item) {
	NSMenuItem *mitem = (NSMenuItem *)item;
	CMenuItemInfo *info = (CMenuItemInfo *)calloc(1, sizeof(CMenuItemInfo));
	NSMenu *parent = [mitem menu];
	info->owner = (CMenuPtr)parent;
	info->index = [parent indexOfItem:mitem];
	info->id = [mitem tag];
	info->title = strdup([[mitem title] UTF8String]);
	if ([mitem hasSubmenu]) {
		info->subMenu = (CMenuPtr)[mitem submenu];
	}
	return info;
}

void disposeMenuItemInfo(CMenuItemInfo *info) {
	free(info->title);
	free(info);
}
