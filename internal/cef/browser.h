#ifndef BROWSER_H_
#define BROWSER_H_
#pragma once

#include "include/capi/cef_browser_capi.h"

cef_browser_host_t *get_cef_browser_host(cef_browser_t *browser);
cef_frame_t *get_cef_focused_frame(cef_browser_t *browser);

#endif // BROWSER_H_