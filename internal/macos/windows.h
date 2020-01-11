// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

typedef void *CWindowPtr;
typedef void *CViewPtr;

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
int themeIsDark(CWindowPtr wnd);
