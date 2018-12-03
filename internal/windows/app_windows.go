package windows

import (
	"syscall"
	"unsafe"

	"github.com/richardwilkes/cef/cef"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/webapp"
)

type driver struct {
	instance             HINSTANCE
	windows              map[HWND]*webapp.Window
	menubars             map[HMENU]*webapp.MenuBar
	menus                map[HMENU]*webapp.Menu
	menuitems            map[int]*menuItem
	awaitingQuitDecision bool
}

var drv = &driver{
	windows:   make(map[HWND]*webapp.Window),
	menubars:  make(map[HMENU]*webapp.MenuBar),
	menus:     make(map[HMENU]*webapp.Menu),
	menuitems: make(map[int]*menuItem),
}

// Driver returns the Windows implementation of the driver.
func Driver() *driver {
	return drv
}

func (d *driver) PrepareForStart() error {
	var err error
	if d.instance, err = GetModuleHandleW(); err != nil {
		return err
	}
	wcx := WNDCLASSEXW{
		Style:    CS_HREDRAW | CS_VREDRAW,
		WndProc:  syscall.NewCallback(d.wndProc),
		Instance: d.instance,
		// Icon: LoadIcon(hInstance, MAKEINTRESOURCE(IDI_CEFCLIENT)),
		// Background: cCOLOR_WINDOW + 1,
		// MenuName: MAKEINTRESOURCE(IDC_CEFCLIENT),
		// IconSm: LoadIcon(wcex.hInstance, MAKEINTRESOURCE(IDI_SMALL)),
	}
	wcx.Size = uint32(unsafe.Sizeof(wcx))
	if wcx.Cursor, err = LoadCursorW__(NULL, IDC_ARROW); err != nil {
		return err
	}
	if wcx.ClassName, err = syscall.UTF16PtrFromString(windowClassName); err != nil {
		return errs.NewWithCause("Unable to convert className to UTF16", err)
	}
	_, err = RegisterClassExW(&wcx)
	return err
}

func (d *driver) PrepareForEventLoop() {
	webapp.WillFinishStartupCallback()
	webapp.DidFinishStartupCallback()
}

func (d *driver) AttemptQuit() {
	switch webapp.CheckQuitCallback() {
	case webapp.Cancel:
		return
	case webapp.Now:
		d.quit()
	case webapp.Later:
		d.awaitingQuitDecision = true
	}
}

func (d *driver) MayQuitNow(quit bool) {
	if d.awaitingQuitDecision {
		d.awaitingQuitDecision = false
		if quit {
			d.quit()
		}
	} else {
		jot.Error("Call to MayQuitNow without AttemptQuit")
	}
}

func (d *driver) quit() {
	webapp.QuittingCallback()
	PostQuitMessage(0)
	cef.QuitMessageLoop()
}
