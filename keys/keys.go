package keys

import (
	"fmt"

	"github.com/richardwilkes/toolbox/log/jot"
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
	A              = &Key{Name: "A", Linux: 0x0026, Win: 0x001e, Mac: 0x0000, MacEquiv: "a"}
	B              = &Key{Name: "B", Linux: 0x0038, Win: 0x0030, Mac: 0x000b, MacEquiv: "b"}
	C              = &Key{Name: "C", Linux: 0x0036, Win: 0x002e, Mac: 0x0008, MacEquiv: "c"}
	D              = &Key{Name: "D", Linux: 0x0028, Win: 0x0020, Mac: 0x0002, MacEquiv: "d"}
	E              = &Key{Name: "E", Linux: 0x001a, Win: 0x0012, Mac: 0x000e, MacEquiv: "e"}
	F              = &Key{Name: "F", Linux: 0x0029, Win: 0x0021, Mac: 0x0003, MacEquiv: "f"}
	G              = &Key{Name: "G", Linux: 0x002a, Win: 0x0022, Mac: 0x0005, MacEquiv: "g"}
	H              = &Key{Name: "H", Linux: 0x002b, Win: 0x0023, Mac: 0x0004, MacEquiv: "h"}
	I              = &Key{Name: "I", Linux: 0x001f, Win: 0x0017, Mac: 0x0022, MacEquiv: "i"}
	J              = &Key{Name: "J", Linux: 0x002c, Win: 0x0024, Mac: 0x0026, MacEquiv: "j"}
	K              = &Key{Name: "K", Linux: 0x002d, Win: 0x0025, Mac: 0x0028, MacEquiv: "k"}
	L              = &Key{Name: "L", Linux: 0x002e, Win: 0x0026, Mac: 0x0025, MacEquiv: "l"}
	M              = &Key{Name: "M", Linux: 0x003a, Win: 0x0032, Mac: 0x002e, MacEquiv: "m"}
	N              = &Key{Name: "N", Linux: 0x0039, Win: 0x0031, Mac: 0x002d, MacEquiv: "n"}
	O              = &Key{Name: "O", Linux: 0x0020, Win: 0x0018, Mac: 0x001f, MacEquiv: "o"}
	P              = &Key{Name: "P", Linux: 0x0021, Win: 0x0019, Mac: 0x0023, MacEquiv: "p"}
	Q              = &Key{Name: "Q", Linux: 0x0018, Win: 0x0010, Mac: 0x000c, MacEquiv: "q"}
	R              = &Key{Name: "R", Linux: 0x001b, Win: 0x0013, Mac: 0x000f, MacEquiv: "r"}
	S              = &Key{Name: "S", Linux: 0x0027, Win: 0x001f, Mac: 0x0001, MacEquiv: "s"}
	T              = &Key{Name: "T", Linux: 0x001c, Win: 0x0014, Mac: 0x0011, MacEquiv: "t"}
	U              = &Key{Name: "U", Linux: 0x001e, Win: 0x0016, Mac: 0x0020, MacEquiv: "u"}
	V              = &Key{Name: "V", Linux: 0x0037, Win: 0x002f, Mac: 0x0009, MacEquiv: "v"}
	W              = &Key{Name: "W", Linux: 0x0019, Win: 0x0011, Mac: 0x000d, MacEquiv: "w"}
	X              = &Key{Name: "X", Linux: 0x0035, Win: 0x002d, Mac: 0x0007, MacEquiv: "x"}
	Y              = &Key{Name: "Y", Linux: 0x001d, Win: 0x0015, Mac: 0x0010, MacEquiv: "y"}
	Z              = &Key{Name: "Z", Linux: 0x0034, Win: 0x002c, Mac: 0x0006, MacEquiv: "z"}
	One            = &Key{Name: "1", Linux: 0x000a, Win: 0x0002, Mac: 0x0012, MacEquiv: "1"}
	Two            = &Key{Name: "2", Linux: 0x000b, Win: 0x0003, Mac: 0x0013, MacEquiv: "2"}
	Three          = &Key{Name: "3", Linux: 0x000c, Win: 0x0004, Mac: 0x0014, MacEquiv: "3"}
	Four           = &Key{Name: "4", Linux: 0x000d, Win: 0x0005, Mac: 0x0015, MacEquiv: "4"}
	Five           = &Key{Name: "5", Linux: 0x000e, Win: 0x0006, Mac: 0x0017, MacEquiv: "5"}
	Six            = &Key{Name: "6", Linux: 0x000f, Win: 0x0007, Mac: 0x0016, MacEquiv: "6"}
	Seven          = &Key{Name: "7", Linux: 0x0010, Win: 0x0008, Mac: 0x001a, MacEquiv: "7"}
	Eight          = &Key{Name: "8", Linux: 0x0011, Win: 0x0009, Mac: 0x001c, MacEquiv: "8"}
	Nine           = &Key{Name: "9", Linux: 0x0012, Win: 0x000a, Mac: 0x0019, MacEquiv: "9"}
	Zero           = &Key{Name: "0", Linux: 0x0013, Win: 0x000b, Mac: 0x001d, MacEquiv: "0"}
	Return         = &Key{Name: "Return", Linux: 0x0024, Win: 0x001c, Mac: 0x0024, MacEquiv: fmt.Sprintf("%c", 0x0d)}
	Escape         = &Key{Name: "Escape", Linux: 0x0009, Win: 0x0001, Mac: 0x0035, MacEquiv: fmt.Sprintf("%c", 0x1b)}
	Backspace      = &Key{Name: "Backspace", Linux: 0x0016, Win: 0x000e, Mac: 0x0033, MacEquiv: fmt.Sprintf("%c", 0x08)}
	Tab            = &Key{Name: "Tab", Linux: 0x0017, Win: 0x000f, Mac: 0x0030, MacEquiv: fmt.Sprintf("%c", 0x09)}
	Space          = &Key{Name: "Space", Linux: 0x0041, Win: 0x0039, Mac: 0x0031, MacEquiv: " "}
	Minus          = &Key{Name: "Minus", Linux: 0x0014, Win: 0x000c, Mac: 0x001b, MacEquiv: "-"}
	Equal          = &Key{Name: "=", Linux: 0x0015, Win: 0x000d, Mac: 0x0018, MacEquiv: "="}
	LeftBracket    = &Key{Name: "[", Linux: 0x0022, Win: 0x001a, Mac: 0x0021, MacEquiv: "["}
	RightBracket   = &Key{Name: "]", Linux: 0x0023, Win: 0x001b, Mac: 0x001e, MacEquiv: "]"}
	Backslash      = &Key{Name: "\\", Linux: 0x0033, Win: 0x002b, Mac: 0x002a, MacEquiv: "\\"}
	Semicolon      = &Key{Name: ";", Linux: 0x002f, Win: 0x0027, Mac: 0x0029, MacEquiv: ";"}
	Quote          = &Key{Name: "'", Linux: 0x0030, Win: 0x0028, Mac: 0x0027, MacEquiv: "'"}
	Backquote      = &Key{Name: "`", Linux: 0x0031, Win: 0x0029, Mac: 0x0032, MacEquiv: "`"}
	Comma          = &Key{Name: ",", Linux: 0x003b, Win: 0x0033, Mac: 0x002b, MacEquiv: ","}
	Period         = &Key{Name: ".", Linux: 0x003c, Win: 0x0034, Mac: 0x002f, MacEquiv: "."}
	Slash          = &Key{Name: "/", Linux: 0x003d, Win: 0x0035, Mac: 0x002c, MacEquiv: "/"}
	F1             = &Key{Name: "F1", Linux: 0x0043, Win: 0x003b, Mac: 0x007a, MacEquiv: fmt.Sprintf("%c", 0xF704)}
	F2             = &Key{Name: "F2", Linux: 0x0044, Win: 0x003c, Mac: 0x0078, MacEquiv: fmt.Sprintf("%c", 0xF705)}
	F3             = &Key{Name: "F3", Linux: 0x0045, Win: 0x003d, Mac: 0x0063, MacEquiv: fmt.Sprintf("%c", 0xF706)}
	F4             = &Key{Name: "F4", Linux: 0x0046, Win: 0x003e, Mac: 0x0076, MacEquiv: fmt.Sprintf("%c", 0xF707)}
	F5             = &Key{Name: "F5", Linux: 0x0047, Win: 0x003f, Mac: 0x0060, MacEquiv: fmt.Sprintf("%c", 0xF708)}
	F6             = &Key{Name: "F6", Linux: 0x0048, Win: 0x0040, Mac: 0x0061, MacEquiv: fmt.Sprintf("%c", 0xF709)}
	F7             = &Key{Name: "F7", Linux: 0x0049, Win: 0x0041, Mac: 0x0062, MacEquiv: fmt.Sprintf("%c", 0xF70a)}
	F8             = &Key{Name: "F8", Linux: 0x004a, Win: 0x0042, Mac: 0x0064, MacEquiv: fmt.Sprintf("%c", 0xF70b)}
	F9             = &Key{Name: "F9", Linux: 0x004b, Win: 0x0043, Mac: 0x0065, MacEquiv: fmt.Sprintf("%c", 0xF70c)}
	F10            = &Key{Name: "F10", Linux: 0x004c, Win: 0x0044, Mac: 0x006d, MacEquiv: fmt.Sprintf("%c", 0xF70d)}
	F11            = &Key{Name: "F11", Linux: 0x005f, Win: 0x0057, Mac: 0x0067, MacEquiv: fmt.Sprintf("%c", 0xF70e)}
	F12            = &Key{Name: "F12", Linux: 0x0060, Win: 0x0058, Mac: 0x006f, MacEquiv: fmt.Sprintf("%c", 0xF70f)}
	F13            = &Key{Name: "F13", Linux: 0x00bf, Win: 0x0064, Mac: 0x0069, MacEquiv: fmt.Sprintf("%c", 0xF710)}
	F14            = &Key{Name: "F14", Linux: 0x00c0, Win: 0x0065, Mac: 0x006b, MacEquiv: fmt.Sprintf("%c", 0xF711)}
	F15            = &Key{Name: "F15", Linux: 0x00c1, Win: 0x0066, Mac: 0x0071, MacEquiv: fmt.Sprintf("%c", 0xF712)}
	Insert         = &Key{Name: "Insert", Linux: 0x0076, Win: 0xe052, Mac: 0x0072, MacEquiv: fmt.Sprintf("%c", 0xF727)}
	Home           = &Key{Name: "Home", Linux: 0x006e, Win: 0xe047, Mac: 0x0073, MacEquiv: fmt.Sprintf("%c", 0xF729)}
	PageUp         = &Key{Name: "PageUp", Linux: 0x0070, Win: 0xe049, Mac: 0x0074, MacEquiv: fmt.Sprintf("%c", 0xF72C)}
	Delete         = &Key{Name: "Delete", Linux: 0x0077, Win: 0xe053, Mac: 0x0075, MacEquiv: fmt.Sprintf("%c", 0xF728)}
	End            = &Key{Name: "End", Linux: 0x0073, Win: 0xe04f, Mac: 0x0077, MacEquiv: fmt.Sprintf("%c", 0xF72B)}
	PageDown       = &Key{Name: "PageDown", Linux: 0x0075, Win: 0xe051, Mac: 0x0079, MacEquiv: fmt.Sprintf("%c", 0xF72D)}
	Right          = &Key{Name: "Right", Linux: 0x0072, Win: 0xe04d, Mac: 0x007c, MacEquiv: fmt.Sprintf("%c", 0xF703)}
	Left           = &Key{Name: "Left", Linux: 0x0071, Win: 0xe04b, Mac: 0x007b, MacEquiv: fmt.Sprintf("%c", 0xF702)}
	Down           = &Key{Name: "Down", Linux: 0x0074, Win: 0xe050, Mac: 0x007d, MacEquiv: fmt.Sprintf("%c", 0xF701)}
	Up             = &Key{Name: "Up", Linux: 0x006f, Win: 0xe048, Mac: 0x007e, MacEquiv: fmt.Sprintf("%c", 0xF700)}
	Clear          = &Key{Name: "Clear", Linux: 0x004d, Win: 0xe045, Mac: 0x0047, MacEquiv: "" /* Don't know what this should be */}
	NumpadDivide   = &Key{Name: "/", Linux: 0x006a, Win: 0xe035, Mac: 0x004b, MacEquiv: "/"}
	NumpadMultiply = &Key{Name: "*", Linux: 0x003f, Win: 0x0037, Mac: 0x0043, MacEquiv: "*"}
	NumpadSubtract = &Key{Name: "-", Linux: 0x0052, Win: 0x004a, Mac: 0x004e, MacEquiv: "-"}
	NumpadAdd      = &Key{Name: "+", Linux: 0x0056, Win: 0x004e, Mac: 0x0045, MacEquiv: "+"}
	NumpadEnter    = &Key{Name: "Enter", Linux: 0x0068, Win: 0xe01c, Mac: 0x004c, MacEquiv: fmt.Sprintf("%c", 0x0d)}
	Numpad1        = &Key{Name: "1", Linux: 0x0057, Win: 0x004f, Mac: 0x0053, MacEquiv: "1"}
	Numpad2        = &Key{Name: "2", Linux: 0x0058, Win: 0x0050, Mac: 0x0054, MacEquiv: "2"}
	Numpad3        = &Key{Name: "3", Linux: 0x0059, Win: 0x0051, Mac: 0x0055, MacEquiv: "3"}
	Numpad4        = &Key{Name: "4", Linux: 0x0053, Win: 0x004b, Mac: 0x0056, MacEquiv: "4"}
	Numpad5        = &Key{Name: "5", Linux: 0x0054, Win: 0x004c, Mac: 0x0057, MacEquiv: "5"}
	Numpad6        = &Key{Name: "6", Linux: 0x0055, Win: 0x004d, Mac: 0x0058, MacEquiv: "6"}
	Numpad7        = &Key{Name: "7", Linux: 0x004f, Win: 0x0047, Mac: 0x0059, MacEquiv: "7"}
	Numpad8        = &Key{Name: "8", Linux: 0x0050, Win: 0x0048, Mac: 0x005b, MacEquiv: "8"}
	Numpad9        = &Key{Name: "9", Linux: 0x0051, Win: 0x0049, Mac: 0x005c, MacEquiv: "9"}
	Numpad0        = &Key{Name: "0", Linux: 0x005a, Win: 0x0052, Mac: 0x0052, MacEquiv: "0"}
	NumpadDecimal  = &Key{Name: ".", Linux: 0x005b, Win: 0x0053, Mac: 0x0041, MacEquiv: "."}
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
	Insert,
	Home,
	PageUp,
	Delete,
	End,
	PageDown,
	Right,
	Left,
	Down,
	Up,
	Clear,
	NumpadDivide,
	NumpadMultiply,
	NumpadSubtract,
	NumpadAdd,
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
	NumpadDecimal,
}

func init() {
	for _, one := range known {
		if _, exists := ByLinuxCode[one.Linux]; exists {
			jot.Fatalf(1, "Linux code %x already exists", one.Linux)
		}
		if _, exists := ByWinCode[one.Win]; exists {
			jot.Fatalf(1, "Win code %x already exists", one.Win)
		}
		if _, exists := ByMacCode[one.Mac]; exists {
			jot.Fatalf(1, "Mac code %x already exists", one.Mac)
		}
		ByLinuxCode[one.Linux] = one
		ByWinCode[one.Win] = one
		ByMacCode[one.Mac] = one
	}
}
