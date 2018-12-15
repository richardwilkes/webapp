package windows

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
)

const windowClassName = "wndClass"

func (d *driver) wndProc(wnd HWND, msg uint32, wparam WPARAM, lparam LPARAM) LRESULT {
	switch msg {
	case WM_COMMAND:
		if mi, ok := d.menuitems[int(wparam)]; ok {
			if mi.handler != nil {
				mi.handler()
			}
		}
	case WM_SIZE:
		if w, ok := d.windows[wnd]; ok {
			size := d.WindowContentSize(w)
			SetWindowPos(HWND(unsafe.Pointer(w.Browser.GetHost().GetWindowHandle())), 0, 0, 0, int32(size.Width), int32(size.Height), SWP_NOZORDER)
		}
	case WM_CLOSE:
		if w, ok := d.windows[wnd]; ok {
			w.AttemptClose()
		} else {
			if err := DestroyWindow(wnd); err != nil {
				jot.Error(err)
			}
		}
		if len(d.windows) == 0 && webapp.QuitAfterLastWindowClosedCallback() {
			webapp.AttemptQuit()
		}
	case WM_DESTROY:
		PostQuitMessage(0)
	case WM_ACTIVATE:
		if w, ok := d.windows[wnd]; ok {
			if wparam&(WA_ACTIVE|WA_CLICKACTIVE) != 0 {
				w.GainedFocus()
			} else {
				w.LostFocus()
			}
		}
		return DefWindowProcW(wnd, msg, wparam, lparam)
	case WM_INITMENUPOPUP:
		if menu, ok := d.menus[HMENU(wparam)]; ok {
			for i := menu.Count() - 1; i >= 0; i-- {
				state := MF_ENABLED
				if item := menu.ItemAtIndex(i); item.ID != 0 {
					if info, exists := d.menuitems[item.ID]; exists && info.validator != nil {
						if !info.validator() {
							state = MF_DISABLED
						}
					}
				}
				EnableMenuItem(HMENU(wparam), i, state|MF_BYPOSITION)
			}
		}
		return DefWindowProcW(wnd, msg, wparam, lparam)
	default:
		return DefWindowProcW(wnd, msg, wparam, lparam)
	}
	return 0
}

func (d *driver) KeyWindow() *webapp.Window {
	return d.windows[GetForegroundWindow()]
}

func (d *driver) BringAllWindowsToFront() {
	list := make([]*webapp.Window, 0)
	if err := EnumWindows(func(wnd HWND, data LPARAM) BOOL {
		if one, ok := d.windows[wnd]; ok {
			list = append(list, one)
		}
		return 1
	}, 0); err != nil {
		jot.Error(err)
		return
	}
	for i, one := range list {
		after := HWND_TOP
		flags := uint32(SWP_NOMOVE | SWP_NOSIZE)
		if i != 0 {
			flags |= SWP_NOACTIVATE
			after = HWND(list[i-1].PlatformPtr)
		}
		if err := SetWindowPos(HWND(one.PlatformPtr), after, 0, 0, 0, 0, flags); err != nil {
			jot.Error(err)
		}
	}
}

func (d *driver) WindowInit(wnd *webapp.Window, style webapp.StyleMask, bounds geom.Rect, title string) error {
	w, err := CreateWindowExW(0, windowClassName, title, WS_OVERLAPPEDWINDOW|WS_CLIPCHILDREN, int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), NULL, NULL, d.instance, NULL)
	if err != nil {
		return err
	}
	wnd.PlatformPtr = uintptr(w)
	d.windows[w] = wnd
	return nil
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) unsafe.Pointer {
	return unsafe.Pointer(wnd.PlatformPtr)
}

func (d *driver) WindowClose(wnd *webapp.Window) {
	wnd.WillCloseCallback()
	hwnd := HWND(wnd.PlatformPtr)
	if err := DestroyWindow(hwnd); err != nil {
		jot.Error(err)
	}
	delete(d.windows, hwnd)
}

func (d *driver) WindowSetTitle(wnd *webapp.Window, title string) {
	if err := SetWindowTextW(HWND(wnd.PlatformPtr), title); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowBounds(wnd *webapp.Window) geom.Rect {
	var rect RECT
	if err := GetWindowRect(HWND(wnd.PlatformPtr), &rect); err != nil {
		jot.Error(err)
	}
	var bounds geom.Rect
	bounds.X = float64(rect.Left)
	bounds.Y = float64(rect.Top)
	bounds.Width = float64(rect.Right - rect.Left)
	bounds.Height = float64(rect.Bottom - rect.Top)
	return bounds
}

func (d *driver) WindowContentSize(wnd *webapp.Window) geom.Size {
	var rect RECT
	if err := GetClientRect(HWND(wnd.PlatformPtr), &rect); err != nil {
		jot.Error(err)
	}
	return geom.Size{
		Width:  float64(rect.Right - rect.Left),
		Height: float64(rect.Bottom - rect.Top),
	}
}

func (d *driver) WindowSetBounds(wnd *webapp.Window, bounds geom.Rect) {
	if err := MoveWindow(HWND(wnd.PlatformPtr), int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), true); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowToFront(wnd *webapp.Window) {
	w := HWND(wnd.PlatformPtr)
	ShowWindow(w, SW_SHOWNORMAL)
	DrawMenuBar(w)
	if err := SetActiveWindow(w); err != nil {
		jot.Error(err)
	}
}

func (d *driver) WindowMinimize(wnd *webapp.Window) {
	ShowWindow(HWND(wnd.PlatformPtr), SW_MINIMIZE)
}

func (d *driver) WindowZoom(wnd *webapp.Window) {
	ShowWindow(HWND(wnd.PlatformPtr), SW_MAXIMIZE)
}
