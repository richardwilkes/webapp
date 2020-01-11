// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

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
