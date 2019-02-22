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
