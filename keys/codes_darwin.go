package keys

// Taken from the old Events.h header in Carbon.
const (
	vkAnsiA              = 0x00
	vkAnsiS              = 0x01
	vkAnsiD              = 0x02
	vkAnsiF              = 0x03
	vkAnsiH              = 0x04
	vkAnsiG              = 0x05
	vkAnsiZ              = 0x06
	vkAnsiX              = 0x07
	vkAnsiC              = 0x08
	vkAnsiV              = 0x09
	vkAnsiB              = 0x0B
	vkAnsiQ              = 0x0C
	vkAnsiW              = 0x0D
	vkAnsiE              = 0x0E
	vkAnsiR              = 0x0F
	vkAnsiY              = 0x10
	vkAnsiT              = 0x11
	vkAnsi1              = 0x12
	vkAnsi2              = 0x13
	vkAnsi3              = 0x14
	vkAnsi4              = 0x15
	vkAnsi6              = 0x16
	vkAnsi5              = 0x17
	vkAnsiEqual          = 0x18
	vkAnsi9              = 0x19
	vkAnsi7              = 0x1A
	vkAnsiMinus          = 0x1B
	vkAnsi8              = 0x1C
	vkAnsi0              = 0x1D
	vkAnsiRightBracket   = 0x1E
	vkAnsiO              = 0x1F
	vkAnsiU              = 0x20
	vkAnsiLeftBracket    = 0x21
	vkAnsiI              = 0x22
	vkAnsiP              = 0x23
	vkAnsiL              = 0x25
	vkAnsiJ              = 0x26
	vkAnsiQuote          = 0x27
	vkAnsiK              = 0x28
	vkAnsiSemicolon      = 0x29
	vkAnsiBackslash      = 0x2A
	vkAnsiComma          = 0x2B
	vkAnsiSlash          = 0x2C
	vkAnsiN              = 0x2D
	vkAnsiM              = 0x2E
	vkAnsiPeriod         = 0x2F
	vkAnsiGrave          = 0x32
	vkAnsiKeypadDecimal  = 0x41
	vkAnsiKeypadMultiply = 0x43
	vkAnsiKeypadPlus     = 0x45
	vkAnsiKeypadClear    = 0x47
	vkAnsiKeypadDivide   = 0x4B
	vkAnsiKeypadEnter    = 0x4C
	vkAnsiKeypadMinus    = 0x4E
	//vkAnsiKeypadEquals   = 0x51
	vkAnsiKeypad0  = 0x52
	vkAnsiKeypad1  = 0x53
	vkAnsiKeypad2  = 0x54
	vkAnsiKeypad3  = 0x55
	vkAnsiKeypad4  = 0x56
	vkAnsiKeypad5  = 0x57
	vkAnsiKeypad6  = 0x58
	vkAnsiKeypad7  = 0x59
	vkAnsiKeypad8  = 0x5B
	vkAnsiKeypad9  = 0x5C
	vkReturn       = 0x24
	vkTab          = 0x30
	vkSpace        = 0x31
	vkDelete       = 0x33
	vkEscape       = 0x35
	vkRightCommand = 0x36
	vkCommand      = 0x37
	vkShift        = 0x38
	vkCapsLock     = 0x39
	vkOption       = 0x3A
	vkControl      = 0x3B
	vkRightShift   = 0x3C
	vkRightOption  = 0x3D
	vkRightControl = 0x3E
	//vkFunction            = 0x3F
	vkF17 = 0x40
	//vkVolumeUp            = 0x48
	//vkVolumeDown          = 0x49
	//vkMute                = 0x4A
	vkF18 = 0x4F
	vkF19 = 0x50
	//vkF20                 = 0x5A
	vkF5  = 0x60
	vkF6  = 0x61
	vkF7  = 0x62
	vkF3  = 0x63
	vkF8  = 0x64
	vkF9  = 0x65
	vkF11 = 0x67
	vkF13 = 0x69
	vkF16 = 0x6A
	vkF14 = 0x6B
	vkF10 = 0x6D
	vkF12 = 0x6F
	vkF15 = 0x71
	//vkHelp                = 0x72
	vkHome          = 0x73
	vkPageUp        = 0x74
	vkForwardDelete = 0x75
	vkF4            = 0x76
	vkEnd           = 0x77
	vkF2            = 0x78
	vkPageDown      = 0x79
	vkF1            = 0x7A
	vkLeftArrow     = 0x7B
	vkRightArrow    = 0x7C
	vkDownArrow     = 0x7D
	vkUpArrow       = 0x7E
)

func init() {
	InsertMapping(vkUpArrow, &Mapping{KeyCode: VirtualKeyUp, Name: UpName})
	InsertMapping(vkLeftArrow, &Mapping{KeyCode: VirtualKeyLeft, Name: LeftName})
	InsertMapping(vkDownArrow, &Mapping{KeyCode: VirtualKeyDown, Name: DownName})
	InsertMapping(vkRightArrow, &Mapping{KeyCode: VirtualKeyRight, Name: RightName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyInsert, Name: InsertName}) // Not on a Mac keyboard
	InsertMapping(vkHome, &Mapping{KeyCode: VirtualKeyHome, Name: HomeName})
	InsertMapping(vkEnd, &Mapping{KeyCode: VirtualKeyEnd, Name: EndName})
	InsertMapping(vkDelete, &Mapping{KeyCode: VirtualKeyBackspace, Name: BackspaceName})
	InsertMapping(vkTab, &Mapping{KeyCode: VirtualKeyTab, KeyChar: '\t', Name: TabName})
	InsertMapping(vkReturn, &Mapping{KeyCode: VirtualKeyReturn, KeyChar: '\n', Name: ReturnName})
	InsertMapping(vkAnsiKeypadEnter, &Mapping{KeyCode: VirtualKeyNumPadEnter, KeyChar: '\n', Name: NumPadEnterName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyEject, Name: EjectName}) // Swallowed by the system
	InsertMapping(vkShift, &Mapping{KeyCode: VirtualKeyShiftLeft, Name: LeftShiftName})
	InsertMapping(vkRightShift, &Mapping{KeyCode: VirtualKeyShiftRight, Name: RightShiftName})
	InsertMapping(vkControl, &Mapping{KeyCode: VirtualKeyControlLeft, Name: LeftControlName})
	InsertMapping(vkRightControl, &Mapping{KeyCode: VirtualKeyControlRight, Name: RightControlName})
	InsertMapping(vkOption, &Mapping{KeyCode: VirtualKeyOptionLeft, Name: LeftOptionName})
	InsertMapping(vkRightOption, &Mapping{KeyCode: VirtualKeyOptionRight, Name: RightOptionName})
	InsertMapping(vkCommand, &Mapping{KeyCode: VirtualKeyCommandLeft, Name: LeftCommandName})
	InsertMapping(vkRightCommand, &Mapping{KeyCode: VirtualKeyCommandRight, Name: RightCommandName})
	InsertMapping(vkCapsLock, &Mapping{KeyCode: VirtualKeyCapsLock, Name: CapsLockName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyMenu, Name: MenuName}) // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyFn, Name: FnName})     // Swallowed by the system
	InsertMapping(vkEscape, &Mapping{KeyCode: VirtualKeyEscape, KeyChar: '\x1b', Name: EscapeName})
	InsertMapping(vkPageUp, &Mapping{KeyCode: VirtualKeyPageUp, Name: PageUpName})
	InsertMapping(vkPageDown, &Mapping{KeyCode: VirtualKeyPageDown, Name: PageDownName})
	InsertMapping(vkSpace, &Mapping{KeyCode: VirtualKeySpace, KeyChar: ' ', Name: SpaceName})
	insertASCIIKeyCodeMapping(vkAnsiQuote, VirtualKeyQuote)
	insertASCIIKeyCodeMapping(vkAnsiComma, VirtualKeyComma)
	insertASCIIKeyCodeMapping(vkAnsiMinus, VirtualKeyMinus)
	insertASCIIKeyCodeMapping(vkAnsiPeriod, VirtualKeyPeriod)
	insertASCIIKeyCodeMapping(vkAnsiSlash, VirtualKeySlash)
	insertASCIIKeyCodeMapping(vkAnsi0, VirtualKey0)
	insertASCIIKeyCodeMapping(vkAnsi1, VirtualKey1)
	insertASCIIKeyCodeMapping(vkAnsi2, VirtualKey2)
	insertASCIIKeyCodeMapping(vkAnsi3, VirtualKey3)
	insertASCIIKeyCodeMapping(vkAnsi4, VirtualKey4)
	insertASCIIKeyCodeMapping(vkAnsi5, VirtualKey5)
	insertASCIIKeyCodeMapping(vkAnsi6, VirtualKey6)
	insertASCIIKeyCodeMapping(vkAnsi7, VirtualKey7)
	insertASCIIKeyCodeMapping(vkAnsi8, VirtualKey8)
	insertASCIIKeyCodeMapping(vkAnsi9, VirtualKey9)
	insertASCIIKeyCodeMapping(vkAnsiSemicolon, VirtualKeySemiColon)
	insertASCIIKeyCodeMapping(vkAnsiEqual, VirtualKeyEqual)
	insertASCIIKeyCodeMapping(vkAnsiA, VirtualKeyA)
	insertASCIIKeyCodeMapping(vkAnsiB, VirtualKeyB)
	insertASCIIKeyCodeMapping(vkAnsiC, VirtualKeyC)
	insertASCIIKeyCodeMapping(vkAnsiD, VirtualKeyD)
	insertASCIIKeyCodeMapping(vkAnsiE, VirtualKeyE)
	insertASCIIKeyCodeMapping(vkAnsiF, VirtualKeyF)
	insertASCIIKeyCodeMapping(vkAnsiG, VirtualKeyG)
	insertASCIIKeyCodeMapping(vkAnsiH, VirtualKeyH)
	insertASCIIKeyCodeMapping(vkAnsiI, VirtualKeyI)
	insertASCIIKeyCodeMapping(vkAnsiJ, VirtualKeyJ)
	insertASCIIKeyCodeMapping(vkAnsiK, VirtualKeyK)
	insertASCIIKeyCodeMapping(vkAnsiL, VirtualKeyL)
	insertASCIIKeyCodeMapping(vkAnsiM, VirtualKeyM)
	insertASCIIKeyCodeMapping(vkAnsiN, VirtualKeyN)
	insertASCIIKeyCodeMapping(vkAnsiO, VirtualKeyO)
	insertASCIIKeyCodeMapping(vkAnsiP, VirtualKeyP)
	insertASCIIKeyCodeMapping(vkAnsiQ, VirtualKeyQ)
	insertASCIIKeyCodeMapping(vkAnsiR, VirtualKeyR)
	insertASCIIKeyCodeMapping(vkAnsiS, VirtualKeyS)
	insertASCIIKeyCodeMapping(vkAnsiT, VirtualKeyT)
	insertASCIIKeyCodeMapping(vkAnsiU, VirtualKeyU)
	insertASCIIKeyCodeMapping(vkAnsiV, VirtualKeyV)
	insertASCIIKeyCodeMapping(vkAnsiW, VirtualKeyW)
	insertASCIIKeyCodeMapping(vkAnsiX, VirtualKeyX)
	insertASCIIKeyCodeMapping(vkAnsiY, VirtualKeyY)
	insertASCIIKeyCodeMapping(vkAnsiZ, VirtualKeyZ)
	insertASCIIKeyCodeMapping(vkAnsiLeftBracket, VirtualKeyLeftBracket)
	insertASCIIKeyCodeMapping(vkAnsiBackslash, VirtualKeyBackSlash)
	insertASCIIKeyCodeMapping(vkAnsiRightBracket, VirtualKeyRightBracket)
	insertASCIIKeyCodeMapping(vkAnsiGrave, VirtualKeyBacktick)
	InsertMapping(vkForwardDelete, &Mapping{KeyCode: VirtualKeyDelete, Name: DeleteName})
	InsertMapping(vkAnsiKeypad0, &Mapping{KeyCode: VirtualKeyNumPad0, KeyChar: '0', Name: NumPad0Name})
	InsertMapping(vkAnsiKeypad1, &Mapping{KeyCode: VirtualKeyNumPad1, KeyChar: '1', Name: NumPad1Name})
	InsertMapping(vkAnsiKeypad2, &Mapping{KeyCode: VirtualKeyNumPad2, KeyChar: '2', Name: NumPad2Name})
	InsertMapping(vkAnsiKeypad3, &Mapping{KeyCode: VirtualKeyNumPad3, KeyChar: '3', Name: NumPad3Name})
	InsertMapping(vkAnsiKeypad4, &Mapping{KeyCode: VirtualKeyNumPad4, KeyChar: '4', Name: NumPad4Name})
	InsertMapping(vkAnsiKeypad5, &Mapping{KeyCode: VirtualKeyNumPad5, KeyChar: '5', Name: NumPad5Name})
	InsertMapping(vkAnsiKeypad6, &Mapping{KeyCode: VirtualKeyNumPad6, KeyChar: '6', Name: NumPad6Name})
	InsertMapping(vkAnsiKeypad7, &Mapping{KeyCode: VirtualKeyNumPad7, KeyChar: '7', Name: NumPad7Name})
	InsertMapping(vkAnsiKeypad8, &Mapping{KeyCode: VirtualKeyNumPad8, KeyChar: '8', Name: NumPad8Name})
	InsertMapping(vkAnsiKeypad9, &Mapping{KeyCode: VirtualKeyNumPad9, KeyChar: '9', Name: NumPad9Name})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumLock, Name: NumLockName})           // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadUp, Name: NumPadUpName})         // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadLeft, Name: NumPadLeftName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadDown, Name: NumPadDownName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadRight, Name: NumPadRightName})   // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadCenter, Name: NumPadCenterName}) // Not on a Mac keyboard
	InsertMapping(vkAnsiKeypadClear, &Mapping{KeyCode: VirtualKeyNumPadClear, Name: NumPadClearName})
	InsertMapping(vkAnsiKeypadDivide, &Mapping{KeyCode: VirtualKeyNumPadDivide, KeyChar: '/', Name: NumPadDivideName})
	InsertMapping(vkAnsiKeypadMultiply, &Mapping{KeyCode: VirtualKeyNumPadMultiply, KeyChar: '*', Name: NumPadMultiplyName})
	InsertMapping(vkAnsiKeypadMinus, &Mapping{KeyCode: VirtualKeyNumPadMinus, KeyChar: '-', Name: NumPadMinusName})
	InsertMapping(vkAnsiKeypadPlus, &Mapping{KeyCode: VirtualKeyNumPadAdd, KeyChar: '+', Name: NumPadAddName})
	InsertMapping(vkAnsiKeypadDecimal, &Mapping{KeyCode: VirtualKeyNumPadDecimal, KeyChar: '.', Name: NumPadDecimalName})
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadDelete, Name: NumPadDeleteName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadHome, Name: NumPadHomeName})         // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadEnd, Name: NumPadEndName})           // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadPageUp, Name: NumPadPageUpName})     // Not on a Mac keyboard
	insertKeyCodeMapping(&Mapping{KeyCode: VirtualKeyNumPadPageDown, Name: NumPadPageDownName}) // Not on a Mac keyboard
	InsertMapping(vkF1, &Mapping{KeyCode: VirtualKeyF1, Name: F1Name})
	InsertMapping(vkF2, &Mapping{KeyCode: VirtualKeyF2, Name: F2Name})
	InsertMapping(vkF3, &Mapping{KeyCode: VirtualKeyF3, Name: F3Name})
	InsertMapping(vkF4, &Mapping{KeyCode: VirtualKeyF4, Name: F4Name})
	InsertMapping(vkF5, &Mapping{KeyCode: VirtualKeyF5, Name: F5Name})
	InsertMapping(vkF6, &Mapping{KeyCode: VirtualKeyF6, Name: F6Name})
	InsertMapping(vkF7, &Mapping{KeyCode: VirtualKeyF7, Name: F7Name})
	InsertMapping(vkF8, &Mapping{KeyCode: VirtualKeyF8, Name: F8Name})
	InsertMapping(vkF9, &Mapping{KeyCode: VirtualKeyF9, Name: F9Name})
	InsertMapping(vkF10, &Mapping{KeyCode: VirtualKeyF10, Name: F10Name})
	InsertMapping(vkF11, &Mapping{KeyCode: VirtualKeyF11, Name: F11Name})
	InsertMapping(vkF12, &Mapping{KeyCode: VirtualKeyF12, Name: F12Name})
	InsertMapping(vkF13, &Mapping{KeyCode: VirtualKeyF13, Name: F13Name})
	InsertMapping(vkF14, &Mapping{KeyCode: VirtualKeyF14, Name: F14Name})
	InsertMapping(vkF15, &Mapping{KeyCode: VirtualKeyF15, Name: F15Name})
	InsertMapping(vkF16, &Mapping{KeyCode: VirtualKeyF16, Name: F16Name})
	InsertMapping(vkF17, &Mapping{KeyCode: VirtualKeyF17, Name: F17Name})
	InsertMapping(vkF18, &Mapping{KeyCode: VirtualKeyF18, Name: F18Name})
	InsertMapping(vkF19, &Mapping{KeyCode: VirtualKeyF19, Name: F19Name})
}

func insertASCIIKeyCodeMapping(scanCode int, keyCode int) {
	InsertMapping(scanCode, &Mapping{KeyCode: keyCode, KeyChar: rune(keyCode), Dynamic: true, Name: string(rune(keyCode))})
}
