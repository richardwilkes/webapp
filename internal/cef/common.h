#ifndef PLATFORM_COMMON_H_
#define PLATFORM_COMMON_H_
#pragma once

#include <stdlib.h>
#include <string.h>
#include <stdio.h>
#include <stdatomic.h>
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"

cef_base_ref_counted_t *refcnt_alloc(int size);
cef_string_t *new_cef_string_from_utf8(const char *str);
cef_settings_t *new_cef_settings();
cef_browser_settings_t *new_cef_browser_settings();
cef_client_t *new_cef_client();
cef_window_info_t *new_cef_window_info(cef_window_handle_t parent, int x, int y, int width, int height);
cef_window_handle_t get_cef_window_handle(cef_browser_host_t *host);
cef_task_t *new_cef_task(int id);

#endif // PLATFORM_COMMON_H_
