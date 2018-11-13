#import <Quartz/Quartz.h>
#import <Cocoa/Cocoa.h>
#import <dispatch/dispatch.h>

typedef void *CMenuPtr;
typedef void *CMenuItemPtr;
typedef void *CWindowPtr;
typedef void *CViewPtr;

typedef struct {
	CGRect bounds;
	CGRect usableBounds;
	double scaleFactor;
	int    isMain;
} Display;

void prepareForStart();

void attemptQuit();
void mayQuitNow(int quit);

void invoke(unsigned long id);
void invokeAfter(unsigned long id, long afterNanos);

void hideApp();
void hideOtherApps();
void showAllApps();

void setMenuBar(CMenuPtr bar);
void setServicesMenu(CMenuPtr menu);
void setWindowMenu(CMenuPtr menu);
void setHelpMenu(CMenuPtr menu);

CMenuPtr newMenu(const char *title);
void disposeMenu(CMenuPtr menu);
int menuItemCount(CMenuPtr menu);
CMenuItemPtr menuItem(CMenuPtr menu, int index);
void insertMenuItem(CMenuPtr menu, CMenuItemPtr item, int index);
void removeMenuItem(CMenuPtr menu, int index);

CMenuItemPtr newMenuSeparator();
CMenuItemPtr newMenuItem(const char *title, const char *selector, const char *key, int modifiers, bool needDelegate);
CMenuPtr subMenu(CMenuItemPtr item);
void setSubMenu(CMenuItemPtr item, CMenuPtr subMenu);
void disposeMenuItem(CMenuItemPtr item);

Display *displays(unsigned long *qty);
CWindowPtr getKeyWindow();
void bringAllWindowsToFront();

CWindowPtr newWindow(int styleMask, double x, double y, double width, double height, const char *title);
CViewPtr contentView(CWindowPtr wnd);
void closeWindow(CWindowPtr wnd);
void setWindowTitle(CWindowPtr wnd, const char *title);
void getWindowBounds(CWindowPtr wnd, double *x, double *y, double *width, double *height);
void setWindowBounds(CWindowPtr wnd, double x, double y, double width, double height);
void getWindowContentSize(CWindowPtr wnd, double *width, double *height);
void bringWindowToFront(CWindowPtr wnd);
void minimizeWindow(CWindowPtr wnd);
void zoomWindow(CWindowPtr wnd);
