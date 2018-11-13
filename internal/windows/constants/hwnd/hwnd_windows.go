package hwnd

import "syscall"

const (
	NOTOPMOST = syscall.Handle(0xFFFFFFFE)
	TOPMOST   = syscall.Handle(0xFFFFFFFF)
	TOP       = syscall.Handle(0)
	BOTTOM    = syscall.Handle(1)
)
