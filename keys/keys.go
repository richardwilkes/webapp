package keys

import (
	"fmt"
)

// Key holds cross-platform information about a key.
type Key struct {
	Name     string
	Linux    int
	Win      int
	Mac      int
	MacEquiv string
}

// Data taken from https://chromium.googlesource.com/chromium/src/+/lkgr/ui/events/keycodes/dom/keycode_converter_data.inc
var (
	A              = &Key{Linux: 0x26, Win: 0x41, Mac: 0x00, MacEquiv: "a", Name: "A"}
	B              = &Key{Linux: 0x38, Win: 0x42, Mac: 0x0b, MacEquiv: "b", Name: "B"}
	C              = &Key{Linux: 0x36, Win: 0x43, Mac: 0x08, MacEquiv: "c", Name: "C"}
	D              = &Key{Linux: 0x28, Win: 0x44, Mac: 0x02, MacEquiv: "d", Name: "D"}
	E              = &Key{Linux: 0x1a, Win: 0x45, Mac: 0x0e, MacEquiv: "e", Name: "E"}
	F              = &Key{Linux: 0x29, Win: 0x46, Mac: 0x03, MacEquiv: "f", Name: "F"}
	G              = &Key{Linux: 0x2a, Win: 0x47, Mac: 0x05, MacEquiv: "g", Name: "G"}
	H              = &Key{Linux: 0x2b, Win: 0x48, Mac: 0x04, MacEquiv: "h", Name: "H"}
	I              = &Key{Linux: 0x1f, Win: 0x49, Mac: 0x22, MacEquiv: "i", Name: "I"}
	J              = &Key{Linux: 0x2c, Win: 0x4a, Mac: 0x26, MacEquiv: "j", Name: "J"}
	K              = &Key{Linux: 0x2d, Win: 0x4b, Mac: 0x28, MacEquiv: "k", Name: "K"}
	L              = &Key{Linux: 0x2e, Win: 0x4c, Mac: 0x25, MacEquiv: "l", Name: "L"}
	M              = &Key{Linux: 0x3a, Win: 0x4d, Mac: 0x2e, MacEquiv: "m", Name: "M"}
	N              = &Key{Linux: 0x39, Win: 0x4e, Mac: 0x2d, MacEquiv: "n", Name: "N"}
	O              = &Key{Linux: 0x20, Win: 0x4f, Mac: 0x1f, MacEquiv: "o", Name: "O"}
	P              = &Key{Linux: 0x21, Win: 0x50, Mac: 0x23, MacEquiv: "p", Name: "P"}
	Q              = &Key{Linux: 0x18, Win: 0x51, Mac: 0x0c, MacEquiv: "q", Name: "Q"}
	R              = &Key{Linux: 0x1b, Win: 0x52, Mac: 0x0f, MacEquiv: "r", Name: "R"}
	S              = &Key{Linux: 0x27, Win: 0x53, Mac: 0x01, MacEquiv: "s", Name: "S"}
	T              = &Key{Linux: 0x1c, Win: 0x54, Mac: 0x11, MacEquiv: "t", Name: "T"}
	U              = &Key{Linux: 0x1e, Win: 0x55, Mac: 0x20, MacEquiv: "u", Name: "U"}
	V              = &Key{Linux: 0x37, Win: 0x56, Mac: 0x09, MacEquiv: "v", Name: "V"}
	W              = &Key{Linux: 0x19, Win: 0x57, Mac: 0x0d, MacEquiv: "w", Name: "W"}
	X              = &Key{Linux: 0x35, Win: 0x58, Mac: 0x07, MacEquiv: "x", Name: "X"}
	Y              = &Key{Linux: 0x1d, Win: 0x59, Mac: 0x10, MacEquiv: "y", Name: "Y"}
	Z              = &Key{Linux: 0x34, Win: 0x5a, Mac: 0x06, MacEquiv: "z", Name: "Z"}
	One            = &Key{Linux: 0x0a, Win: 0x31, Mac: 0x12, MacEquiv: "1", Name: "1"}
	Two            = &Key{Linux: 0x0b, Win: 0x32, Mac: 0x13, MacEquiv: "2", Name: "2"}
	Three          = &Key{Linux: 0x0c, Win: 0x33, Mac: 0x14, MacEquiv: "3", Name: "3"}
	Four           = &Key{Linux: 0x0d, Win: 0x34, Mac: 0x15, MacEquiv: "4", Name: "4"}
	Five           = &Key{Linux: 0x0e, Win: 0x35, Mac: 0x17, MacEquiv: "5", Name: "5"}
	Six            = &Key{Linux: 0x0f, Win: 0x36, Mac: 0x16, MacEquiv: "6", Name: "6"}
	Seven          = &Key{Linux: 0x10, Win: 0x37, Mac: 0x1a, MacEquiv: "7", Name: "7"}
	Eight          = &Key{Linux: 0x11, Win: 0x38, Mac: 0x1c, MacEquiv: "8", Name: "8"}
	Nine           = &Key{Linux: 0x12, Win: 0x39, Mac: 0x19, MacEquiv: "9", Name: "9"}
	Zero           = &Key{Linux: 0x13, Win: 0x30, Mac: 0x1d, MacEquiv: "0", Name: "0"}
	Return         = &Key{Linux: 0x24, Win: 0x0d, Mac: 0x24, MacEquiv: fmt.Sprintf("%c", 0x0d), Name: "Return"}
	Escape         = &Key{Linux: 0x09, Win: 0x1b, Mac: 0x35, MacEquiv: fmt.Sprintf("%c", 0x1b), Name: "Escape"}
	Backspace      = &Key{Linux: 0x16, Win: 0x08, Mac: 0x33, MacEquiv: fmt.Sprintf("%c", 0x08), Name: "Backspace"}
	Tab            = &Key{Linux: 0x17, Win: 0x09, Mac: 0x30, MacEquiv: fmt.Sprintf("%c", 0x09), Name: "Tab"}
	Space          = &Key{Linux: 0x41, Win: 0x20, Mac: 0x31, MacEquiv: " ", Name: "Space"}
	Minus          = &Key{Linux: 0x14, Win: 0xbd, Mac: 0x1b, MacEquiv: "-", Name: "Minus"}
	Equal          = &Key{Linux: 0x15, Win: 0xbb, Mac: 0x18, MacEquiv: "=", Name: "="}
	LeftBracket    = &Key{Linux: 0x22, Win: 0xdb, Mac: 0x21, MacEquiv: "[", Name: "["}
	RightBracket   = &Key{Linux: 0x23, Win: 0xdd, Mac: 0x1e, MacEquiv: "]", Name: "]"}
	Backslash      = &Key{Linux: 0x33, Win: 0xde, Mac: 0x2a, MacEquiv: "\\", Name: "\\"}
	Semicolon      = &Key{Linux: 0x2f, Win: 0xba, Mac: 0x29, MacEquiv: ";", Name: ";"}
	Quote          = &Key{Linux: 0x30, Win: 0xc0, Mac: 0x27, MacEquiv: "'", Name: "'"}
	Backquote      = &Key{Linux: 0x31, Win: 0xdc, Mac: 0x32, MacEquiv: "`", Name: "`"}
	Comma          = &Key{Linux: 0x3b, Win: 0xbc, Mac: 0x2b, MacEquiv: ",", Name: ","}
	Period         = &Key{Linux: 0x3c, Win: 0xbe, Mac: 0x2f, MacEquiv: ".", Name: "."}
	Slash          = &Key{Linux: 0x3d, Win: 0xbf, Mac: 0x2c, MacEquiv: "/", Name: "/"}
	F1             = &Key{Linux: 0x43, Win: 0x70, Mac: 0x7a, MacEquiv: fmt.Sprintf("%c", 0xf704), Name: "F1"}
	F2             = &Key{Linux: 0x44, Win: 0x71, Mac: 0x78, MacEquiv: fmt.Sprintf("%c", 0xf705), Name: "F2"}
	F3             = &Key{Linux: 0x45, Win: 0x72, Mac: 0x63, MacEquiv: fmt.Sprintf("%c", 0xf706), Name: "F3"}
	F4             = &Key{Linux: 0x46, Win: 0x73, Mac: 0x76, MacEquiv: fmt.Sprintf("%c", 0xf707), Name: "F4"}
	F5             = &Key{Linux: 0x47, Win: 0x74, Mac: 0x60, MacEquiv: fmt.Sprintf("%c", 0xf708), Name: "F5"}
	F6             = &Key{Linux: 0x48, Win: 0x75, Mac: 0x61, MacEquiv: fmt.Sprintf("%c", 0xf709), Name: "F6"}
	F7             = &Key{Linux: 0x49, Win: 0x76, Mac: 0x62, MacEquiv: fmt.Sprintf("%c", 0xf70a), Name: "F7"}
	F8             = &Key{Linux: 0x4a, Win: 0x77, Mac: 0x64, MacEquiv: fmt.Sprintf("%c", 0xf70b), Name: "F8"}
	F9             = &Key{Linux: 0x4b, Win: 0x78, Mac: 0x65, MacEquiv: fmt.Sprintf("%c", 0xf70c), Name: "F9"}
	F10            = &Key{Linux: 0x4c, Win: 0x79, Mac: 0x6d, MacEquiv: fmt.Sprintf("%c", 0xf70d), Name: "F10"}
	F11            = &Key{Linux: 0x5f, Win: 0x7a, Mac: 0x67, MacEquiv: fmt.Sprintf("%c", 0xf70e), Name: "F11"}
	F12            = &Key{Linux: 0x60, Win: 0x7b, Mac: 0x6f, MacEquiv: fmt.Sprintf("%c", 0xf70f), Name: "F12"}
	F13            = &Key{Linux: 0xbf, Win: 0x7c, Mac: 0x69, MacEquiv: fmt.Sprintf("%c", 0xf710), Name: "F13"}
	F14            = &Key{Linux: 0xc0, Win: 0x7d, Mac: 0x6b, MacEquiv: fmt.Sprintf("%c", 0xf711), Name: "F14"}
	F15            = &Key{Linux: 0xc1, Win: 0x7e, Mac: 0x71, MacEquiv: fmt.Sprintf("%c", 0xf712), Name: "F15"}
	Delete         = &Key{Linux: 0x77, Win: 0x2e, Mac: 0x75, MacEquiv: fmt.Sprintf("%c", 0xf728), Name: "Delete"}
	Home           = &Key{Linux: 0x6e, Win: 0x24, Mac: 0x73, MacEquiv: fmt.Sprintf("%c", 0xf729), Name: "Home"}
	End            = &Key{Linux: 0x73, Win: 0x23, Mac: 0x77, MacEquiv: fmt.Sprintf("%c", 0xf72b), Name: "End"}
	PageUp         = &Key{Linux: 0x70, Win: 0x21, Mac: 0x74, MacEquiv: fmt.Sprintf("%c", 0xf72c), Name: "PageUp"}
	PageDown       = &Key{Linux: 0x75, Win: 0x22, Mac: 0x79, MacEquiv: fmt.Sprintf("%c", 0xf72d), Name: "PageDown"}
	Left           = &Key{Linux: 0x71, Win: 0x25, Mac: 0x7b, MacEquiv: fmt.Sprintf("%c", 0xf702), Name: "Left"}
	Up             = &Key{Linux: 0x6f, Win: 0x26, Mac: 0x7e, MacEquiv: fmt.Sprintf("%c", 0xf700), Name: "Up"}
	Right          = &Key{Linux: 0x72, Win: 0x27, Mac: 0x7c, MacEquiv: fmt.Sprintf("%c", 0xf703), Name: "Right"}
	Down           = &Key{Linux: 0x74, Win: 0x28, Mac: 0x7d, MacEquiv: fmt.Sprintf("%c", 0xf701), Name: "Down"}
	Clear          = &Key{Linux: 0x4d, Win: 0x90, Mac: 0x47, MacEquiv: "", Name: "Clear"}
	NumpadDivide   = &Key{Linux: 0x6a, Win: 0x6f, Mac: 0x4b, MacEquiv: "/", Name: "/"}
	NumpadMultiply = &Key{Linux: 0x3f, Win: 0x6a, Mac: 0x43, MacEquiv: "*", Name: "*"}
	NumpadAdd      = &Key{Linux: 0x56, Win: 0x6b, Mac: 0x45, MacEquiv: "+", Name: "+"}
	NumpadSubtract = &Key{Linux: 0x52, Win: 0x6c, Mac: 0x4e, MacEquiv: "-", Name: "-"}
	NumpadDecimal  = &Key{Linux: 0x5b, Win: 0x6d, Mac: 0x41, MacEquiv: ".", Name: "."}
	NumpadEnter    = &Key{Linux: 0x68, Win: 0x0d, Mac: 0x4c, MacEquiv: fmt.Sprintf("%c", 0x0d), Name: "Enter"}
	Numpad1        = &Key{Linux: 0x57, Win: 0x61, Mac: 0x53, MacEquiv: "1", Name: "1"}
	Numpad2        = &Key{Linux: 0x58, Win: 0x62, Mac: 0x54, MacEquiv: "2", Name: "2"}
	Numpad3        = &Key{Linux: 0x59, Win: 0x63, Mac: 0x55, MacEquiv: "3", Name: "3"}
	Numpad4        = &Key{Linux: 0x53, Win: 0x64, Mac: 0x56, MacEquiv: "4", Name: "4"}
	Numpad5        = &Key{Linux: 0x54, Win: 0x65, Mac: 0x57, MacEquiv: "5", Name: "5"}
	Numpad6        = &Key{Linux: 0x55, Win: 0x66, Mac: 0x58, MacEquiv: "6", Name: "6"}
	Numpad7        = &Key{Linux: 0x4f, Win: 0x67, Mac: 0x59, MacEquiv: "7", Name: "7"}
	Numpad8        = &Key{Linux: 0x50, Win: 0x68, Mac: 0x5b, MacEquiv: "8", Name: "8"}
	Numpad9        = &Key{Linux: 0x51, Win: 0x69, Mac: 0x5c, MacEquiv: "9", Name: "9"}
	Numpad0        = &Key{Linux: 0x5a, Win: 0x60, Mac: 0x52, MacEquiv: "0", Name: "0"}
)

// Maps to Key by code for each platform
var (
	ByLinuxCode = make(map[int]*Key)
	ByWinCode   = make(map[int]*Key)
	ByMacCode   = make(map[int]*Key)
)

var known = []*Key{
	A,
	B,
	C,
	D,
	E,
	F,
	G,
	H,
	I,
	J,
	K,
	L,
	M,
	N,
	O,
	P,
	Q,
	R,
	S,
	T,
	U,
	V,
	W,
	X,
	Y,
	Z,
	One,
	Two,
	Three,
	Four,
	Five,
	Six,
	Seven,
	Eight,
	Nine,
	Zero,
	Return,
	Escape,
	Backspace,
	Tab,
	Space,
	Minus,
	Equal,
	LeftBracket,
	RightBracket,
	Backslash,
	Semicolon,
	Quote,
	Backquote,
	Comma,
	Period,
	Slash,
	F1,
	F2,
	F3,
	F4,
	F5,
	F6,
	F7,
	F8,
	F9,
	F10,
	F11,
	F12,
	F13,
	F14,
	F15,
	Delete,
	Home,
	End,
	PageUp,
	PageDown,
	Left,
	Up,
	Right,
	Down,
	Clear,
	NumpadDivide,
	NumpadMultiply,
	NumpadAdd,
	NumpadSubtract,
	NumpadDecimal,
	NumpadEnter,
	Numpad1,
	Numpad2,
	Numpad3,
	Numpad4,
	Numpad5,
	Numpad6,
	Numpad7,
	Numpad8,
	Numpad9,
	Numpad0,
}

func init() {
	for _, one := range known {
		if _, exists := ByLinuxCode[one.Linux]; !exists {
			ByLinuxCode[one.Linux] = one
		}
		if _, exists := ByWinCode[one.Win]; !exists {
			ByWinCode[one.Win] = one
		}
		if _, exists := ByMacCode[one.Mac]; !exists {
			ByMacCode[one.Mac] = one
		}
	}
}
