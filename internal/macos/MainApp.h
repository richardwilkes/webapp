#import <Cocoa/Cocoa.h>
#import "include/capi/cef_app_capi.h"
#import "include/cef_application_mac.h"

@interface WebApplication : NSApplication<CefAppProtocol>
{
	BOOL handlingSendEvent;
}
@end
