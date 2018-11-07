#ifndef PLATFORM_COMMON_H_
#define PLATFORM_COMMON_H_
#pragma once

#include <stdlib.h>
#include <string.h>
#include <stdatomic.h>
#include "include/capi/cef_app_capi.h"
#include "include/capi/cef_client_capi.h"

cef_base_ref_counted_t *refcnt_alloc(int size);
cef_string_t *new_cef_string_from_utf8(const char *str);
cef_browser_settings_t *new_cef_browser_settings();
cef_client_t *new_cef_client();
cef_window_info_t *new_cef_window_info();

#endif // PLATFORM_COMMON_H_
