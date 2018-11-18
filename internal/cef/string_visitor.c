#include "string_visitor.h"
#import "_cgo_export.h"

void async_string_visitor_callback(cef_string_visitor_t *self, const cef_string_t *str) {
	cef_string_userfree_utf8_t conv = cef_string_userfree_utf8_alloc();
	cef_string_to_utf8(str->str, str->length, (cef_string_utf8_t *)conv);
	asyncStringVisitorCallback(((async_string_visitor_t *)self)->id, conv->str);
	cef_string_userfree_utf8_free(conv);
}

int async_string_visitor_refcnt_release_callback(cef_base_ref_counted_t *self) {
	async_string_visitor_t *visitor = (async_string_visitor_t *)self;
	int32_t id = visitor->id;
	int result = visitor->release(self);
	if (result) {
		asyncStringVisitorDisposedCallback(id);
	}
	return result;
}

cef_string_visitor_t *new_async_string_visitor(int32_t id) {
	async_string_visitor_t *visitor = (async_string_visitor_t *)refcnt_alloc(sizeof(async_string_visitor_t));
	visitor->id = id;
	visitor->visitor.visit = async_string_visitor_callback;
	visitor->release = visitor->visitor.base.release;
	visitor->visitor.base.release = async_string_visitor_refcnt_release_callback;
	return &visitor->visitor;
}
