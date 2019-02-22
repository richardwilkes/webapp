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
		return 0
	case WM_SIZE:
		if w, ok := d.windows[wnd]; ok && w.Browser != nil {
			size := d.WindowContentSize(w)
			SetWindowPos(HWND(w.Browser.GetHost().GetWindowHandle()), 0, 0, 0, int32(size.Width), int32(size.Height), SWP_NOZORDER)
		}
		return 0
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
		return 0
	case WM_DESTROY:
		PostQuitMessage(0)
		return 0
	case WM_ACTIVATE:
		if w, ok := d.windows[wnd]; ok {
			if wparam&(WA_ACTIVE|WA_CLICKACTIVE) != 0 {
				w.GainedFocus()
				if child := GetWindow(wnd, GW_CHILD); child != NULL {
					SetFocus(child)
				}
				return 0
			}
			w.LostFocus()
		}
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
	}
	return DefWindowProcW(wnd, msg, wparam, lparam)
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
			if w, ok := list[i-1].PlatformData.(HWND); ok {
				flags |= SWP_NOACTIVATE
				after = w
			}
		}
		if w, ok := one.PlatformData.(HWND); ok {
			if err := SetWindowPos(w, after, 0, 0, 0, 0, flags); err != nil {
				jot.Error(err)
			}
		}
	}
}

func (d *driver) WindowInit(wnd *webapp.Window, style webapp.StyleMask, bounds geom.Rect, title string) error {
	w, err := CreateWindowExW(0, windowClassName, title, WS_OVERLAPPEDWINDOW|WS_CLIPCHILDREN, int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), NULL, NULL, d.instance, NULL)
	if err != nil {
		return err
	}
	wnd.PlatformData = w
	d.windows[w] = wnd
	return nil
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) unsafe.Pointer {
	if w, ok := wnd.PlatformData.(HWND); ok {
		return unsafe.Pointer(w)
	}
	return nil
}

func (d *driver) WindowClose(wnd *webapp.Window) {
	wnd.WillCloseCallback()
	if w, ok := wnd.PlatformData.(HWND); ok {
		if err := DestroyWindow(w); err != nil {
			jot.Error(err)
		}
		delete(d.windows, w)
	}
}

func (d *driver) WindowSetTitle(wnd *webapp.Window, title string) {
	if w, ok := wnd.PlatformData.(HWND); ok {
		if err := SetWindowTextW(w, title); err != nil {
			jot.Error(err)
		}
	}
}

func (d *driver) WindowBounds(wnd *webapp.Window) geom.Rect {
	var bounds geom.Rect
	if w, ok := wnd.PlatformData.(HWND); ok {
		var rect RECT
		if err := GetWindowRect(w, &rect); err != nil {
			jot.Error(err)
		}
		bounds.X = float64(rect.Left)
		bounds.Y = float64(rect.Top)
		bounds.Width = float64(rect.Right - rect.Left)
		bounds.Height = float64(rect.Bottom - rect.Top)
	}
	return bounds
}

func (d *driver) WindowContentSize(wnd *webapp.Window) geom.Size {
	var rect RECT
	if w, ok := wnd.PlatformData.(HWND); ok {
		if err := GetClientRect(w, &rect); err != nil {
			jot.Error(err)
		}
	}
	return geom.Size{
		Width:  float64(rect.Right - rect.Left),
		Height: float64(rect.Bottom - rect.Top),
	}
}

func (d *driver) WindowSetBounds(wnd *webapp.Window, bounds geom.Rect) {
	if w, ok := wnd.PlatformData.(HWND); ok {
		if err := MoveWindow(w, int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), true); err != nil {
			jot.Error(err)
		}
	}
}

func (d *driver) WindowToFront(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(HWND); ok {
		ShowWindow(w, SW_SHOWNORMAL)
		DrawMenuBar(w)
		if err := SetActiveWindow(w); err != nil {
			jot.Error(err)
		}
	}
}

func (d *driver) WindowMinimize(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(HWND); ok {
		ShowWindow(w, SW_MINIMIZE)
	}
}

func (d *driver) WindowZoom(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(HWND); ok {
		ShowWindow(w, SW_MAXIMIZE)
	}
}

func (d *driver) WindowThemeIsDark(wnd *webapp.Window) bool {
	// TODO: Implement
	// if w, ok := wnd.PlatformData.(HWND); ok {
	// 	return ...
	// }
	return false
}
