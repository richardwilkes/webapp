package webapp

import (
	"unsafe"

	"github.com/richardwilkes/cef/cef"
)

type client struct {
	keyboardHandler *cef.KeyboardHandler
}

func newClient() *cef.Client {
	c := &client{}
	c.keyboardHandler = cef.NewKeyboardHandler(c)
	c.keyboardHandler.Base().AddRef()
	return cef.NewClient(c)
}

func (c *client) GetContextMenuHandler(self *cef.Client) *cef.ContextMenuHandler {
	return nil
}

func (c *client) GetDialogHandler(self *cef.Client) *cef.DialogHandler {
	return nil
}

func (c *client) GetDisplayHandler(self *cef.Client) *cef.DisplayHandler {
	return nil
}

func (c *client) GetDownloadHandler(self *cef.Client) *cef.DownloadHandler {
	return nil
}

func (c *client) GetDragHandler(self *cef.Client) *cef.DragHandler {
	return nil
}

func (c *client) GetFindHandler(self *cef.Client) *cef.FindHandler {
	return nil
}

func (c *client) GetFocusHandler(self *cef.Client) *cef.FocusHandler {
	return nil
}

func (c *client) GetJsdialogHandler(self *cef.Client) *cef.JsdialogHandler {
	return nil
}

func (c *client) GetKeyboardHandler(self *cef.Client) *cef.KeyboardHandler {
	c.keyboardHandler.Base().AddRef()
	return c.keyboardHandler
}

func (c *client) GetLifeSpanHandler(self *cef.Client) *cef.LifeSpanHandler {
	return nil
}

func (c *client) GetLoadHandler(self *cef.Client) *cef.LoadHandler {
	return nil
}

func (c *client) GetRenderHandler(self *cef.Client) *cef.RenderHandler {
	return nil
}

func (c *client) GetRequestHandler(self *cef.Client) *cef.RequestHandler {
	return nil
}

func (c *client) OnProcessMessageReceived(self *cef.Client, browser *cef.Browser, source_process cef.ProcessID, message *cef.ProcessMessage) int32 {
	return 0
}

func (c *client) OnPreKeyEvent(self *cef.KeyboardHandler, browser *cef.Browser, event *cef.KeyEvent, os_event unsafe.Pointer, is_keyboard_shortcut *int32) int32 {
	return driver.OnPreKeyEvent(event, is_keyboard_shortcut)
}

func (c *client) OnKeyEvent(self *cef.KeyboardHandler, browser *cef.Browser, event *cef.KeyEvent, os_event unsafe.Pointer) int32 {
	return driver.OnKeyEvent(event)
}
