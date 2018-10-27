package keys

import (
	"fmt"
)

// These are platform-independent key codes used by this framework. Keys in the printable space map
// to their ASCII equivalent. Should the framework encounter a scan code not present in this list,
// it will create a new key code for it outside the range 0-255.
const (
	VirtualKeyUp             = 1
	VirtualKeyLeft           = 2
	VirtualKeyDown           = 3
	VirtualKeyRight          = 4
	VirtualKeyInsert         = 5
	VirtualKeyHome           = 6
	VirtualKeyEnd            = 7
	VirtualKeyBackspace      = 8
	VirtualKeyTab            = 9
	VirtualKeyReturn         = 10
	VirtualKeyNumPadEnter    = 13
	VirtualKeyEject          = 15
	VirtualKeyShiftLeft      = 16
	VirtualKeyShiftRight     = 17
	VirtualKeyControlLeft    = 18
	VirtualKeyControlRight   = 19
	VirtualKeyOptionLeft     = 20
	VirtualKeyOptionRight    = 21
	VirtualKeyCommandLeft    = 22
	VirtualKeyCommandRight   = 23
	VirtualKeyCapsLock       = 24
	VirtualKeyMenu           = 25
	VirtualKeyFn             = 26
	VirtualKeyEscape         = 27
	VirtualKeyPageUp         = 28
	VirtualKeyPageDown       = 29
	VirtualKeySpace          = 32
	VirtualKeyQuote          = 39
	VirtualKeyComma          = 44
	VirtualKeyMinus          = 45
	VirtualKeyPeriod         = 46
	VirtualKeySlash          = 47
	VirtualKey0              = 48
	VirtualKey1              = 49
	VirtualKey2              = 50
	VirtualKey3              = 51
	VirtualKey4              = 52
	VirtualKey5              = 53
	VirtualKey6              = 54
	VirtualKey7              = 55
	VirtualKey8              = 56
	VirtualKey9              = 57
	VirtualKeySemiColon      = 59
	VirtualKeyEqual          = 61
	VirtualKeyA              = 65
	VirtualKeyB              = 66
	VirtualKeyC              = 67
	VirtualKeyD              = 68
	VirtualKeyE              = 69
	VirtualKeyF              = 70
	VirtualKeyG              = 71
	VirtualKeyH              = 72
	VirtualKeyI              = 73
	VirtualKeyJ              = 74
	VirtualKeyK              = 75
	VirtualKeyL              = 76
	VirtualKeyM              = 77
	VirtualKeyN              = 78
	VirtualKeyO              = 79
	VirtualKeyP              = 80
	VirtualKeyQ              = 81
	VirtualKeyR              = 82
	VirtualKeyS              = 83
	VirtualKeyT              = 84
	VirtualKeyU              = 85
	VirtualKeyV              = 86
	VirtualKeyW              = 87
	VirtualKeyX              = 88
	VirtualKeyY              = 89
	VirtualKeyZ              = 90
	VirtualKeyLeftBracket    = 91
	VirtualKeyBackSlash      = 92
	VirtualKeyRightBracket   = 93
	VirtualKeyBacktick       = 96
	VirtualKeyDelete         = 127
	VirtualKeyNumPad0        = 130
	VirtualKeyNumPad1        = 131
	VirtualKeyNumPad2        = 132
	VirtualKeyNumPad3        = 133
	VirtualKeyNumPad4        = 134
	VirtualKeyNumPad5        = 135
	VirtualKeyNumPad6        = 136
	VirtualKeyNumPad7        = 137
	VirtualKeyNumPad8        = 138
	VirtualKeyNumPad9        = 139
	VirtualKeyNumLock        = 140
	VirtualKeyNumPadUp       = 141
	VirtualKeyNumPadLeft     = 142
	VirtualKeyNumPadDown     = 143
	VirtualKeyNumPadRight    = 144
	VirtualKeyNumPadCenter   = 145
	VirtualKeyNumPadClear    = 146
	VirtualKeyNumPadDivide   = 147
	VirtualKeyNumPadMultiply = 148
	VirtualKeyNumPadMinus    = 149
	VirtualKeyNumPadAdd      = 150
	VirtualKeyNumPadDecimal  = 151
	VirtualKeyNumPadDelete   = 152
	VirtualKeyNumPadHome     = 153
	VirtualKeyNumPadEnd      = 154
	VirtualKeyNumPadPageUp   = 155
	VirtualKeyNumPadPageDown = 156
	VirtualKeyF1             = 201
	VirtualKeyF2             = 202
	VirtualKeyF3             = 203
	VirtualKeyF4             = 204
	VirtualKeyF5             = 205
	VirtualKeyF6             = 206
	VirtualKeyF7             = 207
	VirtualKeyF8             = 208
	VirtualKeyF9             = 209
	VirtualKeyF10            = 210
	VirtualKeyF11            = 211
	VirtualKeyF12            = 212
	VirtualKeyF13            = 213
	VirtualKeyF14            = 214
	VirtualKeyF15            = 215
	VirtualKeyF16            = 216
	VirtualKeyF17            = 217
	VirtualKeyF18            = 218
	VirtualKeyF19            = 219
)

// Mapping provides a mapping between key codes and the rune they represent, if any.
type Mapping struct {
	KeyCode int
	KeyChar rune
	// Dynamic means the KeyChar value is only one of multiple possibilities.
	Dynamic bool
	Name    string
}

var (
	scanCodeToMapping = make(map[int]*Mapping)
	keyCodeToMapping  = make(map[int]*Mapping)
)

// InsertMapping inserts a mapping for the specified scanCode into the map used by
// MappingForScanCode and MappingForKeyCode.
func InsertMapping(scanCode int, mapping *Mapping) {
	scanCodeToMapping[scanCode] = mapping
	insertKeyCodeMapping(mapping)
}

func insertKeyCodeMapping(mapping *Mapping) {
	keyCodeToMapping[mapping.KeyCode] = mapping
}

// MappingForScanCode returns the mapping for the specified scanCode, or nil.
func MappingForScanCode(scanCode int) *Mapping {
	mapping, ok := scanCodeToMapping[scanCode]
	if !ok {
		mapping = &Mapping{KeyCode: scanCode << 8, Dynamic: true, Name: fmt.Sprintf("ScanCode %d", scanCode)}
		InsertMapping(scanCode, mapping)
	}
	return mapping
}

// MappingForKeyCode returns the mapping for the specified keyCode, or nil.
func MappingForKeyCode(keyCode int) *Mapping {
	if mapping, ok := keyCodeToMapping[keyCode]; ok {
		return mapping
	}
	return nil
}

// IsControlAction returns true if the keyCode should trigger a control, such as a button, that is
// focused.
func IsControlAction(keyCode int) bool {
	return keyCode == VirtualKeyReturn || keyCode == VirtualKeyNumPadEnter || keyCode == VirtualKeySpace
}

// Transform a scan code into a key code and character. If the scan code has no
// mapping or the mapping is dynamic, the first character of chars will be
// returned as the character, if available.
func Transform(scanCode int, chars string) (code int, ch rune) {
	extract := true
	if mapping := MappingForScanCode(scanCode); mapping != nil {
		code = mapping.KeyCode
		if !mapping.Dynamic {
			ch = mapping.KeyChar
			extract = false
		}
	} else {
		code = scanCode
	}
	if extract && chars != "" {
		ch = (([]rune)(chars))[0]
	}
	return
}
