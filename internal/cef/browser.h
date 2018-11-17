#ifndef BROWSER_H_
#define BROWSER_H_
#pragma once

#include "include/capi/cef_browser_capi.h"

cef_browser_host_t *get_cef_browser_host(cef_browser_t *browser);
void cef_browser_cut(cef_browser_t *browser);
void cef_browser_copy(cef_browser_t *browser);
void cef_browser_paste(cef_browser_t *browser);
void cef_browser_delete(cef_browser_t *browser);
void cef_browser_select_all(cef_browser_t *browser);

#endif // BROWSER_H_