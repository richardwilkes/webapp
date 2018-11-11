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
