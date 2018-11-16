#import "MenuItemDelegate.h"
#import "_cgo_export.h"

@implementation MenuItemDelegate

- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
	return validateMenuItemCallback(menuItem.tag);
}

- (void)handleMenuItem:(id)sender {
	handleMenuItemCallback(((NSMenuItem *)sender).tag);
}

@end
