package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/internal/cef"
	"github.com/richardwilkes/webapp/internal/windows/constants/swp"
	"github.com/richardwilkes/webapp/internal/windows/constants/wm"
)

// WNDCLASSEXW is defined here:
// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagwndclassexw
type WNDCLASSEXW struct {
	Size       uint32
	Style      uint32
	WndProc    uintptr
	ClsExtra   int32
	WndExtra   int32
	Instance   syscall.Handle
	Icon       syscall.Handle
	Cursor     syscall.Handle
	Background syscall.Handle
	MenuName   *uint16
	ClassName  *uint16
	IconSm     syscall.Handle
}

// WndProc provides standard handling of window messages.
func WndProc(wnd syscall.Handle, msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case wm.SIZE:
		if w, ok := drv.windows[wnd]; ok {
			size := drv.WindowContentSize(w)
			SetWindowPos(syscall.Handle(unsafe.Pointer(cef.GetWindowHandle(cef.GetBrowserHost(w.Browser)))), 0, 0, 0, int32(size.Width), int32(size.Height), swp.NOZORDER)
		}
	case wm.CLOSE:
		if w, ok := drv.windows[wnd]; ok {
			w.AttemptClose()
		} else {
			if err := DestroyWindow(wnd); err != nil {
				jot.Error(err)
			}
		}
		if len(drv.windows) == 0 && webapp.QuitAfterLastWindowClosedCallback() {
			webapp.AttemptQuit()
		}
	case wm.DESTROY:
		PostQuitMessage(0)
	default:
		return DefWindowProcW(wnd, msg, wparam, lparam)
	}
	return 0
}
