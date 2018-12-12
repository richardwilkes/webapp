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

// Modifiers contains flags indicating which modifier keys were down when an
// event occurred.
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

// PlatformMenuModifier returns the platform's standard menu command key
// modifier.
func PlatformMenuModifier() Modifiers {
	if runtime.GOOS == "darwin" {
		return CommandModifier
	}
	return ControlModifier
}

// PlatformMenuModifierDown returns true if the platform's menu command key is
// being pressed.
func (m Modifiers) PlatformMenuModifierDown() bool {
	mask := PlatformMenuModifier()
	return m&mask == mask
}

// String returns a text representation of these modifiers. If any modifiers
// are present, then the string will end with a '+'.
func (m Modifiers) String() string {
	if m == 0 {
		return ""
	}
	var buffer bytes.Buffer
	if m&ControlModifier == ControlModifier {
		buffer.WriteString("Ctrl+")
	}
	if m&OptionModifier == OptionModifier {
		if runtime.GOOS == "darwin" {
			buffer.WriteString("Opt+")
		} else {
			buffer.WriteString("Alt+")
		}
	}
	if m&ShiftModifier == ShiftModifier {
		buffer.WriteString("Shift+")
	}
	if m&CapsLockModifier == CapsLockModifier {
		buffer.WriteString("Caps+")
	}
	if m&CommandModifier == CommandModifier {
		if runtime.GOOS == "darwin" {
			buffer.WriteString("Cmd+")
		} else {
			buffer.WriteString("Win+")
		}
	}
	return buffer.String()
}

// SymbolString returns a representation of these modifiers using symbols
// rather than names. It will not have a trailing '+' symbol.
func (m Modifiers) SymbolString() string {
	if m == 0 {
		return ""
	}
	var buffer bytes.Buffer
	if m&ControlModifier == ControlModifier {
		buffer.WriteString("\u2303")
	}
	if m&OptionModifier == OptionModifier {
		buffer.WriteString("\u2325")
	}
	if m&ShiftModifier == ShiftModifier {
		buffer.WriteString("\u21e7")
	}
	if m&CapsLockModifier == CapsLockModifier {
		buffer.WriteString("\u21ea")
	}
	if m&CommandModifier == CommandModifier {
		if runtime.GOOS == "darwin" {
			buffer.WriteString("\u2318")
		} else {
			buffer.WriteString("\u2756")
		}
	}
	return buffer.String()
}
