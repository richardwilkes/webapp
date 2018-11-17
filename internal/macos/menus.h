#import <stdbool.h>

typedef void *CMenuPtr;
typedef void *CMenuItemPtr;

typedef struct {
	int      tag;
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
CMenuItemPtr menuItemWithTag(CMenuPtr menu, int tag);
void insertMenuItem(CMenuPtr menu, CMenuItemPtr item, int index);
void removeMenuItem(CMenuPtr menu, int index);

CMenuItemPtr newMenuSeparator();
CMenuItemPtr newMenuItem(int tag, const char *title, const char *selector, const char *key, int modifiers, bool needDelegate);
CMenuPtr subMenu(CMenuItemPtr item);
void setSubMenu(CMenuItemPtr item, CMenuPtr subMenu);
void disposeMenuItem(CMenuItemPtr item);

CMenuItemInfo *menuItemInfo(CMenuItemPtr item);
void disposeMenuItemInfo(CMenuItemInfo *info);