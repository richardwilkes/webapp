#if !defined(PLATFORM_WINDOWS_H)
#define PLATFORM_WINDOWS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#if defined(PLATFORM_WINDOWS_IMPL)
#define PLATFORM_WINDOWS_EXPORT __declspec(dllexport)
#else
#define PLATFORM_WINDOWS_EXPORT __declspec(dllimport)
#endif

	struct _menu;

	typedef void(*voidCallback)(void);
	typedef unsigned char(*boolCallback)(void);
	typedef int32_t(*int32Callback)(void);
	typedef void(*voidStringCallback)(const char *);
	typedef void (*voidMenuCallback)(struct _menu *);

	typedef struct _platform {
		void *impl;
	} platform;

	typedef struct _window {
		void *impl;
	} window;

	typedef struct _menu {
		void *impl;
	} menu;

	typedef struct _menuItem {
		void *impl;
	} menuItem;

	PLATFORM_WINDOWS_EXPORT void windowsNewPlatform(platform *platform, voidCallback willFinishStartup, voidCallback didFinishStartup, voidCallback willActivate, voidCallback didActivate, voidCallback willDeactivate, voidCallback didDeactivate, boolCallback quitAfterLastWindowClosed, int32Callback checkQuit);
	PLATFORM_WINDOWS_EXPORT void windowsPlatformSetMenuBar(platform *platform, menu *menu);
	PLATFORM_WINDOWS_EXPORT void windowsNewWindow(platform *platform, window *window, int width, int height, const char *url);
	PLATFORM_WINDOWS_EXPORT void windowsWindowSetTitle(window *window, const char *title);
	PLATFORM_WINDOWS_EXPORT void windowsNewMenuBar(menu *menu);
	PLATFORM_WINDOWS_EXPORT void windowsNewMenu(menu *menu, const char *title);
	PLATFORM_WINDOWS_EXPORT int windowsMenuGetCount(menu *menu);
	PLATFORM_WINDOWS_EXPORT void windowsMenuInsertItem(menu *menu, menuItem *menuItem, int index);
	PLATFORM_WINDOWS_EXPORT void windowsNewMenuItem(menuItem *menuItem, const char *title);
	PLATFORM_WINDOWS_EXPORT void windowsNewMenuItemSeparator(menuItem *menuItem);
	PLATFORM_WINDOWS_EXPORT void windowsMenuItemHack(menuItem *menuItem, menu *menu);

	void windowsInit(platform *platform);

#ifdef __cplusplus
}
#endif

#endif // PLATFORM_WINDOWS_H
