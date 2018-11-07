#include "platform_common.h"

typedef struct _cef_refcnt_t {
	volatile atomic_int_fast32_t count;
	int32_t dummy; // Just here to pad it out to 8 bytes
} cef_refcnt_t;

void refcnt_add(cef_base_ref_counted_t *p) {
	cef_refcnt_t *ref = (cef_refcnt_t *)(((void *)p) - sizeof(cef_refcnt_t));
	atomic_fetch_add(&ref->count, 1);
}

int refcnt_release(cef_base_ref_counted_t *p) {
	cef_refcnt_t *ref = (cef_refcnt_t *)(((void *)p) - sizeof(cef_refcnt_t));
	if (atomic_fetch_sub(&ref->count, 1) == 0) {
		free(ref);
		return 1;
	}
	return 0;
}

int refcnt_has_one(cef_base_ref_counted_t *p) {
	cef_refcnt_t *ref = (cef_refcnt_t *)(((void *)p) - sizeof(cef_refcnt_t));
	return atomic_load(&ref->count) == 1;
}

int refcnt_has_at_least_one(cef_base_ref_counted_t *p) {
	cef_refcnt_t *ref = (cef_refcnt_t *)(((void *)p) - sizeof(cef_refcnt_t));
	return atomic_load(&ref->count) > 0;
}

cef_base_ref_counted_t *refcnt_alloc(int size) {
	void *p = calloc(1, size + sizeof(cef_refcnt_t));
	cef_base_ref_counted_t *ret = (cef_base_ref_counted_t *)(p + sizeof(cef_refcnt_t));
	ret->size = size;
	ret->add_ref = refcnt_add;
	ret->release = refcnt_release;
	ret->has_one_ref = refcnt_has_one;
	ret->has_at_least_one_ref = refcnt_has_at_least_one;
	return ret;
}

cef_string_t *new_cef_string_from_utf8(const char *str) {
	cef_string_t *s = (cef_string_t *)calloc(1, sizeof(cef_string_t));
	cef_string_from_utf8(str, strlen(str), s);
	return s;
}

cef_browser_settings_t *new_cef_browser_settings() {
	cef_browser_settings_t *settings = (cef_browser_settings_t *)calloc(1, sizeof(cef_browser_settings_t));
	settings->size = sizeof(cef_browser_settings_t);
	return settings;
}

cef_client_t *new_cef_client() {
	return (cef_client_t *)refcnt_alloc(sizeof(cef_client_t));
}

cef_window_info_t *new_cef_window_info() {
	return (cef_window_info_t *)calloc(1, sizeof(cef_window_info_t));
}
