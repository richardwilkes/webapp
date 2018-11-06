#include <windows.h>

#include "platform_windows.h"
#include "_cgo_export.h"

void windowsPrepareForStart(platform *platform) {
	CoInitializeEx(NULL, COINIT_APARTMENTTHREADED);
	windowsNewPlatform(platform, willFinishStartupCallback, didFinishStartupCallback, willActivateCallback, didActivateCallback, willDeactivateCallback, didDeactivateCallback, quitAfterLastWindowClosedCallback, checkQuitCallback, handleMenuItemCallback);
}
