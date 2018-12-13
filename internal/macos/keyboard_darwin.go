package macos

import "github.com/richardwilkes/cef/cef"

func (d *driver) OnPreKeyEvent(event *cef.KeyEvent, is_keyboard_shortcut *int32) int32 {
	return 0
}

func (d *driver) OnKeyEvent(event *cef.KeyEvent) int32 {
	return 0
}
