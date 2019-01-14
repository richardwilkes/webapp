package windows

import (
	"github.com/richardwilkes/cef/cef"
	"github.com/richardwilkes/webapp/keys"
)

type keyboardProxy struct {
}

func (d *driver) OnPreKeyEvent(event *cef.KeyEvent, isKeyboardShortcut *int32) int32 {
	lookup := d.refreshMenuKeyForWindow(d.KeyWindow())
	mods := eventModifiers(event)
	if k, ok := keys.ByWinCode[int(event.WindowsKeyCode)]; ok {
		if _, ok := lookup[mods.String()+k.Name]; ok {
			*isKeyboardShortcut = 1
		}
	}
	return 0
}

func (d *driver) OnKeyEvent(event *cef.KeyEvent) int32 {
	lookup := d.refreshMenuKeyForWindow(d.KeyWindow())
	mods := eventModifiers(event)
	if k, ok := keys.ByWinCode[int(event.WindowsKeyCode)]; ok {
		if item, ok := lookup[mods.String()+k.Name]; ok {
			if item.validator == nil || item.validator() {
				if item.handler != nil {
					item.handler()
				}
			}
			return 1
		}
	}
	return 0
}

func eventModifiers(event *cef.KeyEvent) keys.Modifiers {
	var km keys.Modifiers
	mods := cef.EventFlags(event.Modifiers)
	if mods&cef.EventflagControlDown != 0 {
		km |= keys.ControlModifier
	}
	if mods&cef.EventflagShiftDown != 0 {
		km |= keys.ShiftModifier
	}
	if mods&cef.EventflagAltDown != 0 {
		km |= keys.OptionModifier
	}
	return km
}
