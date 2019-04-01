package windows

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/win32"
)

const windowClassName = "wndClass"

func (d *driver) wndProc(wnd win32.HWND, msg uint32, wparam win32.WPARAM, lparam win32.LPARAM) win32.LRESULT {
	switch msg {
	case win32.WM_COMMAND:
		if mi, ok := d.menuitems[int(wparam)]; ok {
			if mi.handler != nil {
				mi.handler()
			}
		}
		return 0
	case win32.WM_SIZE:
		if w, ok := d.windows[wnd]; ok && w.Browser != nil {
			size := d.WindowContentSize(w)
			win32.SetWindowPos(win32.HWND(w.Browser.GetHost().GetWindowHandle()), 0, 0, 0, int32(size.Width), int32(size.Height), win32.SWP_NOZORDER)
		}
		return 0
	case win32.WM_CLOSE:
		if w, ok := d.windows[wnd]; ok {
			w.AttemptClose()
		} else {
			win32.DestroyWindow(wnd)
		}
		if len(d.windows) == 0 && webapp.QuitAfterLastWindowClosedCallback() {
			webapp.AttemptQuit()
		}
		return 0
	case win32.WM_DESTROY:
		win32.PostQuitMessage(0)
		return 0
	case win32.WM_ACTIVATE:
		if w, ok := d.windows[wnd]; ok {
			if wparam&(win32.WA_ACTIVE|win32.WA_CLICKACTIVE) != 0 {
				w.GainedFocus()
				if child := win32.GetWindow(wnd, win32.GW_CHILD); child != win32.NULL {
					win32.SetFocus(child)
				}
				return 0
			}
			w.LostFocus()
		}
	case win32.WM_INITMENUPOPUP:
		if menu, ok := d.menus[win32.HMENU(wparam)]; ok {
			for i := menu.Count() - 1; i >= 0; i-- {
				state := win32.MF_ENABLED
				if item := menu.ItemAtIndex(i); item.ID != 0 {
					if info, exists := d.menuitems[item.ID]; exists && info.validator != nil {
						if !info.validator() {
							state = win32.MF_DISABLED
						}
					}
				}
				win32.EnableMenuItem(win32.HMENU(wparam), i, state|win32.MF_BYPOSITION)
			}
		}
	}
	return win32.DefWindowProc(wnd, msg, wparam, lparam)
}

func (d *driver) KeyWindow() *webapp.Window {
	return d.windows[win32.GetForegroundWindow()]
}

func (d *driver) BringAllWindowsToFront() {
	list := make([]*webapp.Window, 0)
	win32.EnumWindows(func(wnd win32.HWND, data win32.LPARAM) win32.BOOL {
		if one, ok := d.windows[wnd]; ok {
			list = append(list, one)
		}
		return 1
	}, 0)
	for i, one := range list {
		after := win32.HWND_TOP
		flags := uint32(win32.SWP_NOMOVE | win32.SWP_NOSIZE)
		if i != 0 {
			if w, ok := list[i-1].PlatformData.(win32.HWND); ok {
				flags |= win32.SWP_NOACTIVATE
				after = w
			}
		}
		if w, ok := one.PlatformData.(win32.HWND); ok {
			win32.SetWindowPos(w, after, 0, 0, 0, 0, flags)
		}
	}
}

func (d *driver) WindowInit(wnd *webapp.Window, style webapp.StyleMask, bounds geom.Rect, title string) error {
	w := win32.CreateWindowExS(0, windowClassName, title, win32.WS_OVERLAPPEDWINDOW|win32.WS_CLIPCHILDREN, int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), win32.NULL, win32.NULL, d.instance, win32.NULL)
	if w == win32.NULL {
		return errs.New("unable to create window")
	}
	wnd.PlatformData = w
	d.windows[w] = wnd
	return nil
}

func (d *driver) WindowBrowserParent(wnd *webapp.Window) unsafe.Pointer {
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		return unsafe.Pointer(w)
	}
	return nil
}

func (d *driver) WindowClose(wnd *webapp.Window) {
	wnd.WillCloseCallback()
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		win32.DestroyWindow(w)
		delete(d.windows, w)
	}
}

func (d *driver) WindowSetTitle(wnd *webapp.Window, title string) {
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		win32.SetWindowTextS(w, title)
	}
}

func (d *driver) WindowBounds(wnd *webapp.Window) geom.Rect {
	var bounds geom.Rect
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		var rect win32.RECT
		win32.GetWindowRect(w, &rect)
		bounds.X = float64(rect.Left)
		bounds.Y = float64(rect.Top)
		bounds.Width = float64(rect.Right - rect.Left)
		bounds.Height = float64(rect.Bottom - rect.Top)
	}
	return bounds
}

func (d *driver) WindowContentSize(wnd *webapp.Window) geom.Size {
	var rect win32.RECT
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		win32.GetClientRect(w, &rect)
	}
	return geom.Size{
		Width:  float64(rect.Right - rect.Left),
		Height: float64(rect.Bottom - rect.Top),
	}
}

func (d *driver) WindowSetBounds(wnd *webapp.Window, bounds geom.Rect) {
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		win32.MoveWindow(w, int32(bounds.X), int32(bounds.Y), int32(bounds.Width), int32(bounds.Height), true)
	}
}

func (d *driver) WindowToFront(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		win32.ShowWindow(w, win32.SW_SHOWNORMAL)
		win32.DrawMenuBar(w)
		win32.SetActiveWindow(w)
	}
}

func (d *driver) WindowMinimize(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		win32.ShowWindow(w, win32.SW_MINIMIZE)
	}
}

func (d *driver) WindowZoom(wnd *webapp.Window) {
	if w, ok := wnd.PlatformData.(win32.HWND); ok {
		win32.ShowWindow(w, win32.SW_MAXIMIZE)
	}
}

func (d *driver) WindowThemeIsDark(wnd *webapp.Window) bool {
	// TODO: Implement
	// if w, ok := wnd.PlatformData.(win32.HWND); ok {
	// 	return ...
	// }
	return false
}
