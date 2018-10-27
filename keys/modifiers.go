package keys

import (
	"bytes"
	"runtime"
)

// Possible Modifiers values.
const (
	CapsLockModifier Modifiers = 1 << iota
	ShiftModifier
	ControlModifier
	OptionModifier
	CommandModifier
	NonStickyModifiers = ShiftModifier | ControlModifier | OptionModifier | CommandModifier
	AllModifiers       = CapsLockModifier | NonStickyModifiers
)

// Modifiers contains flags indicating which modifier keys were down when an event occurred.
type Modifiers int

// CapsLockDown returns true if the caps lock key is being pressed.
func (m Modifiers) CapsLockDown() bool {
	return m&CapsLockModifier == CapsLockModifier
}

// ShiftDown returns true if the shift key is being pressed.
func (m Modifiers) ShiftDown() bool {
	return m&ShiftModifier == ShiftModifier
}

// ControlDown returns true if the control key is being pressed.
func (m Modifiers) ControlDown() bool {
	return m&ControlModifier == ControlModifier
}

// OptionDown returns true if the option/alt key is being pressed.
func (m Modifiers) OptionDown() bool {
	return m&OptionModifier == OptionModifier
}

// CommandDown returns true if the command/meta key is being pressed.
func (m Modifiers) CommandDown() bool {
	return m&CommandModifier == CommandModifier
}

// PlatformMenuModifier returns the platform's standard menu command key modifier.
func PlatformMenuModifier() Modifiers {
	if runtime.GOOS == "darwin" {
		return CommandModifier
	}
	return ControlModifier
}

// PlatformMenuModifierDown returns true if the platform's menu command key is being pressed.
func (m Modifiers) PlatformMenuModifierDown() bool {
	mask := PlatformMenuModifier()
	return m&mask == mask
}

// String implements the fmt.Stringer interface.
func (m Modifiers) String() string {
	if m == 0 {
		return ""
	}
	var buffer bytes.Buffer
	m.append(&buffer, ControlModifier, ControlName)
	m.append(&buffer, OptionModifier, OptionName)
	m.append(&buffer, ShiftModifier, ShiftName)
	m.append(&buffer, CapsLockModifier, CapsLockName)
	if runtime.GOOS == "darwin" {
		m.append(&buffer, CommandModifier, CommandName)
	} else {
		m.append(&buffer, CommandModifier, WindowsName)
	}
	return buffer.String()
}

func (m Modifiers) append(buffer *bytes.Buffer, modifier Modifiers, name string) {
	if m&modifier == modifier {
		buffer.WriteString(name)
	}
}
