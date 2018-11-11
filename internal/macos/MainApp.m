#import "MainApp.h"

@implementation WebApplication

- (BOOL)isHandlingSendEvent {
	return handlingSendEvent;
}

- (void)setHandlingSendEvent:(BOOL)handling {
	handlingSendEvent = handling;
}

- (void)sendEvent:(NSEvent*)event {
	BOOL handling = handlingSendEvent;
	handlingSendEvent = YES;
	[super sendEvent:event];
	handlingSendEvent = handling;
}

@end
