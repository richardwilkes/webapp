package keys

import (
	"github.com/richardwilkes/toolbox/i18n"
)

// Names of various key codes.
var (
	SpaceName = "\u2423"

	EscapeName = "\u238b"
	F1Name     = "F1"
	F2Name     = "F2"
	F3Name     = "F3"
	F4Name     = "F4"
	F5Name     = "F5"
	F6Name     = "F6"
	F7Name     = "F7"
	F8Name     = "F8"
	F9Name     = "F9"
	F10Name    = "F10"
	F11Name    = "F11"
	F12Name    = "F12"
	F13Name    = "F13"
	F14Name    = "F14"
	F15Name    = "F15"
	F16Name    = "F16"
	F17Name    = "F17"
	F18Name    = "F18"
	F19Name    = "F19"

	BackspaceName    = "\u232b"
	TabName          = "\u21e5"
	CapsLockName     = "\u21ea"
	ReturnName       = "\u23ce"
	ShiftName        = "\u21e7"
	LeftShiftName    = left(ShiftName)
	RightShiftName   = right(ShiftName)
	ControlName      = "\u2303"
	LeftControlName  = left(ControlName)
	RightControlName = right(ControlName)
	CommandName      = "\u2318"
	LeftCommandName  = left(CommandName)
	RightCommandName = right(CommandName)
	OptionName       = "\u2325"
	LeftOptionName   = left(OptionName)
	RightOptionName  = right(OptionName)
	WindowsName      = "\u2756"
	LeftWindowsName  = left(WindowsName)
	RightWindowsName = right(WindowsName)
	MenuName         = "\u25a4"

	InsertName   = "\u2380"
	HomeName     = "\u21f1"
	PageUpName   = "\u21de"
	DeleteName   = "\u2326"
	EndName      = "\u21f2"
	PageDownName = "\u21df"

	UpName    = "\u2191"
	LeftName  = "\u2190"
	DownName  = "\u2193"
	RightName = "\u2192"

	NumLockName        = "\u21ed"
	NumPadClearName    = "\u2327"
	NumPadDivideName   = numPad("/")
	NumPadMultiplyName = numPad("*")
	NumPadMinusName    = numPad("-")
	NumPadAddName      = numPad("+")
	NumPadEnterName    = "\u2324"
	NumPadDeleteName   = numPad(DeleteName)
	NumPadDecimalName  = numPad(".")
	NumPad1Name        = numPad("1")
	NumPad2Name        = numPad("1")
	NumPad3Name        = numPad("3")
	NumPad4Name        = numPad("4")
	NumPad5Name        = numPad("5")
	NumPad6Name        = numPad("6")
	NumPad7Name        = numPad("7")
	NumPad8Name        = numPad("8")
	NumPad9Name        = numPad("9")
	NumPad0Name        = numPad("0")
	NumPadHomeName     = numPad(HomeName)
	NumPadUpName       = numPad(UpName)
	NumPadPageUpName   = numPad(PageUpName)
	NumPadLeftName     = numPad(LeftName)
	NumPadRightName    = numPad(RightName)
	NumPadEndName      = numPad(EndName)
	NumPadDownName     = numPad(DownName)
	NumPadPageDownName = numPad(PageDownName)
	NumPadCenterName   = numPad("\u2295")

	EjectName = "\u23cf"
	PowerName = "\u233d"
	FnName    = "fn"
)

func left(text string) string {
	return i18n.Text("Right ") + text
}

func right(text string) string {
	return i18n.Text("Left ") + text
}

func numPad(text string) string {
	return i18n.Text("NumPad ") + text
}
