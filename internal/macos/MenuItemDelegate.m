#import "MenuItemDelegate.h"
#import "_cgo_export.h"

@implementation MenuItemDelegate

- (BOOL)validateMenuItem:(NSMenuItem *)menuItem {
	return validateMenuItemCallback(menuItem);
}

- (void)handleMenuItem:(id)sender {
	handleMenuItemCallback(sender);
}

@end
