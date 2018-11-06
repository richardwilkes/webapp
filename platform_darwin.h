#import <Quartz/Quartz.h>
#import <Cocoa/Cocoa.h>
#import <dispatch/dispatch.h>
#import "include/capi/cef_app_capi.h"
#import "include/capi/cef_client_capi.h"

typedef void *Menu;
typedef void *MenuItem;
typedef void *Window;

typedef struct {
	CGRect bounds;
	CGRect usableBounds;
	double scaleFactor;
	int    isMain;
} Display;

void prepareForStart();
void attemptQuit();
void mayQuitNow(int quit);
void hideApp();
void hideOtherApps();
void showAllApps();
void invoke(unsigned long id);
void invokeAfter(unsigned long id, long afterNanos);

void setMenuBar(Menu bar);
void setServicesMenu(Menu menu);
void setWindowMenu(Menu menu);
void setHelpMenu(Menu menu);

Menu newMenu(const char *title);
void disposeMenu(Menu menu);
int menuItemCount(Menu menu);
MenuItem menuItem(Menu menu, int index);
void insertMenuItem(Menu menu, MenuItem item, int index);
void removeMenuItem(Menu menu, int index);

MenuItem newMenuItem(const char *title, const char *selector, const char *key, int modifiers, bool needDelegate);
Menu subMenu(MenuItem item);
void setSubMenu(MenuItem item, Menu subMenu);
MenuItem newMenuSeparator();
void disposeMenuItem(MenuItem item);

Window getKeyWindow();
void bringAllWindowsToFront();
Window newWindow(int styleMask, double x, double y, double width, double height, const char *url);
void closeWindow(Window window);
void setWindowTitle(Window window, const char *title);
void getWindowBounds(Window window, double *x, double *y, double *width, double *height);
void setWindowBounds(Window window, double x, double y, double width, double height);
void bringWindowToFront(Window window);
void minimizeWindow(Window window);
void zoomWindow(Window window);

Display *displays(unsigned long *qty);
