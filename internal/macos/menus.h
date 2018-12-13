#import <stdbool.h>

typedef void *CMenuPtr;
typedef void *CMenuItemPtr;

typedef struct {
	CMenuPtr owner;
	int      index;
	int      id;
	char *   title;
	CMenuPtr subMenu;
} CMenuItemInfo;

void setMenuBar(CMenuPtr bar);
void setServicesMenu(CMenuPtr menu);
void setWindowMenu(CMenuPtr menu);
void setHelpMenu(CMenuPtr menu);

CMenuPtr newMenu(const char *title);
void disposeMenu(CMenuPtr menu);
int menuItemCount(CMenuPtr menu);
CMenuItemPtr menuItemAtIndex(CMenuPtr menu, int index);
void insertMenuItem(CMenuPtr menu, CMenuItemPtr item, int index);
void removeMenuItem(CMenuPtr menu, int index);

CMenuItemPtr newMenuSeparator();
CMenuItemPtr newMenuItem(int cmdID, const char *title, const char *selector, const char *key, int modifiers, bool needDelegate);
CMenuPtr subMenu(CMenuItemPtr item);
void setSubMenu(CMenuItemPtr item, CMenuPtr subMenu);
void disposeMenuItem(CMenuItemPtr item);
void setMenuItemTitle(CMenuItemPtr item, const char *title);
CMenuItemInfo *menuItemInfo(CMenuItemPtr item);
void disposeMenuItemInfo(CMenuItemInfo *info);
