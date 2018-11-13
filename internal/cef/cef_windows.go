package cef

import (
	// #cgo CFLAGS: -I ${SRCDIR}/../../cef
	// #cgo LDFLAGS: -L${SRCDIR}/../../cef/Release -lcef
	// #include "common.h"
	"C"
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/atexit"
)

// ExecuteProcess is used to start the secondary CEF processes. If this is
// the main process, this call will do nothing and return. If it is a
// secondary process, the call will not return.
func ExecuteProcess(instance syscall.Handle) {
	args := (*C.cef_main_args_t)(C.calloc(1, C.sizeof_struct__cef_main_args_t))
	args.instance = C.HINSTANCE(unsafe.Pointer(instance))
	if code := C.cef_execute_process(args, nil, nil); code >= 0 {
		atexit.Exit(int(code))
	}
}
