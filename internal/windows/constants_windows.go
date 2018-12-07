package windows

// Miscellaneous
const (
	NULL          = 0
	CW_USEDEFAULT = 0x80000000
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-loadcursora
const (
	IDC_APPSTARTING = 32650
	IDC_ARROW       = 32512
	IDC_CROSS       = 32515
	IDC_HAND        = 32649
	IDC_HELP        = 32651
	IDC_IBEAM       = 32513
	IDC_ICON        = 32641
	IDC_NO          = 32648
	IDC_SIZE        = 32640
	IDC_SIZEALL     = 32646
	IDC_SIZENESW    = 32643
	IDC_SIZENS      = 32645
	IDC_SIZENWSE    = 32642
	IDC_SIZEWE      = 32644
	IDC_UPARROW     = 32516
	IDC_WAIT        = 32514
)

// https://docs.microsoft.com/en-us/windows/desktop/winmsg/window-class-styles
const (
	CS_VREDRAW         = 0x0001
	CS_HREDRAW         = 0x0002
	CS_DBLCLKS         = 0x0008
	CS_OWNDC           = 0x0020
	CS_CLASSDC         = 0x0040
	CS_PARENTDC        = 0x0080
	CS_NOCLOSE         = 0x0200
	CS_SAVEBITS        = 0x0800
	CS_BYTEALIGNCLIENT = 0x1000
	CS_BYTEALIGNWINDOW = 0x2000
	CS_GLOBALCLASS     = 0x4000
	CS_DROPSHADOW      = 0x00020000
)

// http://www.pinvoke.net/default.aspx/Enums/DisplayDeviceStateFlags.html
const (
	DISPLAY_DEVICE_ACTIVE           = 0x00000001
	DISPLAY_DEVICE_MULTIDRIVER      = 0x00000002
	DISPLAY_DEVICE_PRIMARY_DEVICE   = 0x00000004
	DISPLAY_DEVICE_MIRRORING_DRIVER = 0x00000008
	DISPLAY_DEVICE_VGA_COMPATIBLE   = 0x00000010
	DISPLAY_DEVICE_REMOVABLE        = 0x00000020
	DISPLAY_DEVICE_DISCONNECTED     = 0x20000000
	DISPLAY_DEVICE_REMOTE           = 0x40000000
	DISPLAY_DEVICE_MODESPRUNED      = 0x80000000
)

// https://docs.microsoft.com/en-us/windows/desktop/api/windef/ne-windef-dpi_awareness
const (
	DPI_AWARENESS_INVALID = iota
	DPI_AWARENESS_UNAWARE
	DPI_AWARENESS_SYSTEM_AWARE
	DPI_AWARENESS_PER_MONITOR_AWARE
)

// https://docs.microsoft.com/en-us/windows/desktop/hidpi/dpi-awareness-context
const (
	DPI_AWARENESS_CONTEXT_UNAWARE              = ^DPI_AWARENESS_CONTEXT(0) // -1
	DPI_AWARENESS_CONTEXT_SYSTEM_AWARE         = ^DPI_AWARENESS_CONTEXT(1) // -2
	DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE    = ^DPI_AWARENESS_CONTEXT(2) // -3
	DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2 = ^DPI_AWARENESS_CONTEXT(3) // -4
	DPI_AWARENESS_CONTEXT_UNAWARE_GDISCALED    = ^DPI_AWARENESS_CONTEXT(4) // -5
)

// Predefined window handles
const (
	HWND_BROADCAST = HWND(0xFFFF)
	HWND_BOTTOM    = HWND(1)
	HWND_TOP       = HWND(0)
	HWND_DESKTOP   = HWND(0)
	HWND_TOPMOST   = ^HWND(0) // -1
	HWND_NOTOPMOST = ^HWND(1) // -2
	HWND_MESSAGE   = ^HWND(2) // -3
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-deletemenu
const (
	MF_BYCOMMAND  = 0x00000000
	MF_BYPOSITION = 0x00000400
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmenuiteminfow
const (
	MFS_ENABLED   = 0x00000000
	MFS_UNCHECKED = MFS_ENABLED
	MFS_UNHILITE  = MFS_ENABLED
	MFS_DISABLED  = 0x00000003
	MFS_GRAYED    = MFS_DISABLED
	MFS_CHECKED   = 0x00000008
	MFS_HILITE    = 0x00000080
	MFS_DEFAULT   = 0x00001000
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmenuiteminfow
const (
	MFT_STRING       = 0x00000000
	MFT_BITMAP       = 0x00000004
	MFT_MENUBARBREAK = 0x00000020
	MFT_MENUBREAK    = 0x00000040
	MFT_OWNERDRAW    = 0x00000100
	MFT_RADIOCHECK   = 0x00000200
	MFT_SEPARATOR    = 0x00000800
	MFT_RIGHTORDER   = 0x00002000
	MFT_RIGHTJUSTIFY = 0x00004000
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/ns-winuser-tagmenuiteminfow
const (
	MIIM_STATE      = 0x00000001
	MIIM_ID         = 0x00000002
	MIIM_SUBMENU    = 0x00000004
	MIIM_CHECKMARKS = 0x00000008
	MIIM_TYPE       = 0x00000010
	MIIM_DATA       = 0x00000020
	MIIM_STRING     = 0x00000040
	MIIM_BITMAP     = 0x00000080
	MIIM_FTYPE      = 0x00000100
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-showwindow
const (
	SW_HIDE            = 0
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = SW_MAXIMIZE
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-setwindowpos
const (
	SWP_NOSIZE         = 0x0001
	SWP_NOMOVE         = 0x0002
	SWP_NOZORDER       = 0x0004
	SWP_NOREDRAW       = 0x0008
	SWP_NOACTIVATE     = 0x0010
	SWP_DRAWFRAME      = 0x0020
	SWP_FRAMECHANGED   = SWP_DRAWFRAME
	SWP_SHOWWINDOW     = 0x0040
	SWP_HIDEWINDOW     = 0x0080
	SWP_NOCOPYBITS     = 0x0100
	SWP_NOOWNERZORDER  = 0x0200
	SWP_NOREPOSITION   = SWP_NOOWNERZORDER
	SWP_NOSENDCHANGING = 0x0400
	SWP_DEFERERASE     = 0x2000
	SWP_ASYNCWINDOWPOS = 0x4000
)

// Windows message constants
const (
	WM_NULL                        = 0x0000
	WM_CREATE                      = 0x0001
	WM_DESTROY                     = 0x0002
	WM_MOVE                        = 0x0003
	WM_SIZE                        = 0x0005
	WM_ACTIVATE                    = 0x0006
	WM_SETFOCUS                    = 0x0007
	WM_KILLFOCUS                   = 0x0008
	WM_ENABLE                      = 0x000A
	WM_SETREDRAW                   = 0x000B
	WM_SETTEXT                     = 0x000C
	WM_GETTEXT                     = 0x000D
	WM_GETTEXTLENGTH               = 0x000E
	WM_PAINT                       = 0x000F
	WM_CLOSE                       = 0x0010
	WM_QUERYENDSESSION             = 0x0011
	WM_QUIT                        = 0x0012
	WM_QUERYOPEN                   = 0x0013
	WM_ERASEBKGND                  = 0x0014
	WM_SYSCOLORCHANGE              = 0x0015
	WM_ENDSESSION                  = 0x0016
	WM_SHOWWINDOW                  = 0x0018
	WM_WININICHANGE                = 0x001A
	WM_SETTINGCHANGE               = WM_WININICHANGE
	WM_DEVMODECHANGE               = 0x001B
	WM_ACTIVATEAPP                 = 0x001C
	WM_FONTCHANGE                  = 0x001D
	WM_TIMECHANGE                  = 0x001E
	WM_CANCELMODE                  = 0x001F
	WM_SETCURSOR                   = 0x0020
	WM_MOUSEACTIVATE               = 0x0021
	WM_CHILDACTIVATE               = 0x0022
	WM_QUEUESYNC                   = 0x0023
	WM_GETMINMAXINFO               = 0x0024
	WM_PAINTICON                   = 0x0026
	WM_ICONERASEBKGND              = 0x0027
	WM_NEXTDLGCTL                  = 0x0028
	WM_SPOOLERSTATUS               = 0x002A
	WM_DRAWITEM                    = 0x002B
	WM_MEASUREITEM                 = 0x002C
	WM_DELETEITEM                  = 0x002D
	WM_VKEYTOITEM                  = 0x002E
	WM_CHARTOITEM                  = 0x002F
	WM_SETFONT                     = 0x0030
	WM_GETFONT                     = 0x0031
	WM_SETHOTKEY                   = 0x0032
	WM_GETHOTKEY                   = 0x0033
	WM_QUERYDRAGICON               = 0x0037
	WM_COMPAREITEM                 = 0x0039
	WM_GETOBJECT                   = 0x003D
	WM_COMPACTING                  = 0x0041
	WM_COMMNOTIFY                  = 0x0044
	WM_WINDOWPOSCHANGING           = 0x0046
	WM_WINDOWPOSCHANGED            = 0x0047
	WM_POWER                       = 0x0048
	WM_COPYDATA                    = 0x004A
	WM_CANCELJOURNAL               = 0x004B
	WM_NOTIFY                      = 0x004E
	WM_INPUTLANGCHANGEREQUEST      = 0x0050
	WM_INPUTLANGCHANGE             = 0x0051
	WM_TCARD                       = 0x0052
	WM_HELP                        = 0x0053
	WM_USERCHANGED                 = 0x0054
	WM_NOTIFYFORMAT                = 0x0055
	WM_CONTEXTMENU                 = 0x007B
	WM_STYLECHANGING               = 0x007C
	WM_STYLECHANGED                = 0x007D
	WM_DISPLAYCHANGE               = 0x007E
	WM_GETICON                     = 0x007F
	WM_SETICON                     = 0x0080
	WM_NCCREATE                    = 0x0081
	WM_NCDESTROY                   = 0x0082
	WM_NCCALCSIZE                  = 0x0083
	WM_NCHITTEST                   = 0x0084
	WM_NCPAINT                     = 0x0085
	WM_NCACTIVATE                  = 0x0086
	WM_GETDLGCODE                  = 0x0087
	WM_SYNCPAINT                   = 0x0088
	WM_NCMOUSEMOVE                 = 0x00A0
	WM_NCLBUTTONDOWN               = 0x00A1
	WM_NCLBUTTONUP                 = 0x00A2
	WM_NCLBUTTONDBLCLK             = 0x00A3
	WM_NCRBUTTONDOWN               = 0x00A4
	WM_NCRBUTTONUP                 = 0x00A5
	WM_NCRBUTTONDBLCLK             = 0x00A6
	WM_NCMBUTTONDOWN               = 0x00A7
	WM_NCMBUTTONUP                 = 0x00A8
	WM_NCMBUTTONDBLCLK             = 0x00A9
	WM_NCXBUTTONDOWN               = 0x00AB
	WM_NCXBUTTONUP                 = 0x00AC
	WM_NCXBUTTONDBLCLK             = 0x00AD
	WM_INPUT_DEVICE_CHANGE         = 0x00FE
	WM_INPUT                       = 0x00FF
	WM_KEYDOWN                     = 0x0100
	WM_KEYUP                       = 0x0101
	WM_CHAR                        = 0x0102
	WM_DEADCHAR                    = 0x0103
	WM_SYSKEYDOWN                  = 0x0104
	WM_SYSKEYUP                    = 0x0105
	WM_SYSCHAR                     = 0x0106
	WM_SYSDEADCHAR                 = 0x0107
	WM_UNICHAR                     = 0x0109
	WM_KEYFIRST                    = WM_KEYDOWN
	WM_KEYLAST                     = WM_UNICHAR
	WM_IME_STARTCOMPOSITION        = 0x010D
	WM_IME_ENDCOMPOSITION          = 0x010E
	WM_IME_COMPOSITION             = 0x010F
	WM_IME_KEYLAST                 = WM_IME_COMPOSITION
	WM_INITDIALOG                  = 0x0110
	WM_COMMAND                     = 0x0111
	WM_SYSCOMMAND                  = 0x0112
	WM_TIMER                       = 0x0113
	WM_HSCROLL                     = 0x0114
	WM_VSCROLL                     = 0x0115
	WM_INITMENU                    = 0x0116
	WM_INITMENUPOPUP               = 0x0117
	WM_MENUSELECT                  = 0x011F
	WM_MENUCHAR                    = 0x0120
	WM_ENTERIDLE                   = 0x0121
	WM_MENURBUTTONUP               = 0x0122
	WM_MENUDRAG                    = 0x0123
	WM_MENUGETOBJECT               = 0x0124
	WM_UNINITMENUPOPUP             = 0x0125
	WM_MENUCOMMAND                 = 0x0126
	WM_CHANGEUISTATE               = 0x0127
	WM_UPDATEUISTATE               = 0x0128
	WM_QUERYUISTATE                = 0x0129
	WM_CTLCOLORMSGBOX              = 0x0132
	WM_CTLCOLOREDIT                = 0x0133
	WM_CTLCOLORLISTBOX             = 0x0134
	WM_CTLCOLORBTN                 = 0x0135
	WM_CTLCOLORDLG                 = 0x0136
	WM_CTLCOLORSCROLLBAR           = 0x0137
	WM_CTLCOLORSTATIC              = 0x0138
	WM_MN_GETHMENU                 = 0x01E1
	WM_MOUSEMOVE                   = 0x0200
	WM_LBUTTONDOWN                 = 0x0201
	WM_LBUTTONUP                   = 0x0202
	WM_LBUTTONDBLCLK               = 0x0203
	WM_RBUTTONDOWN                 = 0x0204
	WM_RBUTTONUP                   = 0x0205
	WM_RBUTTONDBLCLK               = 0x0206
	WM_MBUTTONDOWN                 = 0x0207
	WM_MBUTTONUP                   = 0x0208
	WM_MBUTTONDBLCLK               = 0x0209
	WM_MOUSEWHEEL                  = 0x020A
	WM_XBUTTONDOWN                 = 0x020B
	WM_XBUTTONUP                   = 0x020C
	WM_XBUTTONDBLCLK               = 0x020D
	WM_MOUSEHWHEEL                 = 0x020E
	WM_MOUSEFIRST                  = WM_MOUSEMOVE
	WM_PARENTNOTIFY                = 0x0210
	WM_ENTERMENULOOP               = 0x0211
	WM_EXITMENULOOP                = 0x0212
	WM_NEXTMENU                    = 0x0213
	WM_SIZING                      = 0x0214
	WM_CAPTURECHANGED              = 0x0215
	WM_MOVING                      = 0x0216
	WM_POWERBROADCAST              = 0x0218
	WM_DEVICECHANGE                = 0x0219
	WM_MDICREATE                   = 0x0220
	WM_MDIDESTROY                  = 0x0221
	WM_MDIACTIVATE                 = 0x0222
	WM_MDIRESTORE                  = 0x0223
	WM_MDINEXT                     = 0x0224
	WM_MDIMAXIMIZE                 = 0x0225
	WM_MDITILE                     = 0x0226
	WM_MDICASCADE                  = 0x0227
	WM_MDIICONARRANGE              = 0x0228
	WM_MDIGETACTIVE                = 0x0229
	WM_MDISETMENU                  = 0x0230
	WM_ENTERSIZEMOVE               = 0x0231
	WM_EXITSIZEMOVE                = 0x0232
	WM_DROPFILES                   = 0x0233
	WM_MDIREFRESHMENU              = 0x0234
	WM_IME_SETCONTEXT              = 0x0281
	WM_IME_NOTIFY                  = 0x0282
	WM_IME_CONTROL                 = 0x0283
	WM_IME_COMPOSITIONFULL         = 0x0284
	WM_IME_SELECT                  = 0x0285
	WM_IME_CHAR                    = 0x0286
	WM_IME_REQUEST                 = 0x0288
	WM_IME_KEYDOWN                 = 0x0290
	WM_IME_KEYUP                   = 0x0291
	WM_MOUSEHOVER                  = 0x02A1
	WM_MOUSELEAVE                  = 0x02A3
	WM_NCMOUSEHOVER                = 0x02A0
	WM_NCMOUSELEAVE                = 0x02A2
	WM_WTSSESSION_CHANGE           = 0x02B1
	WM_TABLET_FIRST                = 0x02c0
	WM_TABLET_LAST                 = 0x02df
	WM_CUT                         = 0x0300
	WM_COPY                        = 0x0301
	WM_PASTE                       = 0x0302
	WM_CLEAR                       = 0x0303
	WM_UNDO                        = 0x0304
	WM_RENDERFORMAT                = 0x0305
	WM_RENDERALLFORMATS            = 0x0306
	WM_DESTROYCLIPBOARD            = 0x0307
	WM_DRAWCLIPBOARD               = 0x0308
	WM_PAINTCLIPBOARD              = 0x0309
	WM_VSCROLLCLIPBOARD            = 0x030A
	WM_SIZECLIPBOARD               = 0x030B
	WM_ASKCBFORMATNAME             = 0x030C
	WM_CHANGECBCHAIN               = 0x030D
	WM_HSCROLLCLIPBOARD            = 0x030E
	WM_QUERYNEWPALETTE             = 0x030F
	WM_PALETTEISCHANGING           = 0x0310
	WM_PALETTECHANGED              = 0x0311
	WM_HOTKEY                      = 0x0312
	WM_PRINT                       = 0x0317
	WM_PRINTCLIENT                 = 0x0318
	WM_APPCOMMAND                  = 0x0319
	WM_THEMECHANGED                = 0x031A
	WM_CLIPBOARDUPDATE             = 0x031D
	WM_DWMCOMPOSITIONCHANGED       = 0x031E
	WM_DWMNCRENDERINGCHANGED       = 0x031F
	WM_DWMCOLORIZATIONCOLORCHANGED = 0x0320
	WM_DWMWINDOWMAXIMIZEDCHANGE    = 0x0321
	WM_GETTITLEBARINFOEX           = 0x033F
	WM_HANDHELDFIRST               = 0x0358
	WM_HANDHELDLAST                = 0x035F
	WM_AFXFIRST                    = 0x0360
	WM_AFXLAST                     = 0x037F
	WM_PENWINFIRST                 = 0x0380
	WM_PENWINLAST                  = 0x038F
	WM_USER                        = 0x0400
	WM_REFLECT                     = WM_USER + 0x1C00
	WM_APP                         = 0x8000
)

// https://docs.microsoft.com/en-us/windows/desktop/winmsg/window-styles
const (
	WS_OVERLAPPED       = 0x00000000
	WS_MAXIMIZEBOX      = 0x00010000 // Has maximize control
	WS_MINIMIZEBOX      = 0x00020000 // Has minimize control
	WS_THICKFRAME       = 0x00040000 // Has a sizing border
	WS_SYSMENU          = 0x00080000 // Has a window menu on title bar
	WS_HSCROLL          = 0x00100000 // Has a horizontal scroll bar
	WS_VSCROLL          = 0x00200000 // Has a vertical scroll bar
	WS_DLGFRAME         = 0x00400000 // Dialog border
	WS_BORDER           = 0x00800000 // Thin line border
	WS_CAPTION          = 0x00C00000 // Has title bar
	WS_MAXIMIZE         = 0x01000000 // Initially maximized
	WS_CLIPCHILDREN     = 0x02000000
	WS_CLIPSIBLINGS     = 0x04000000
	WS_DISABLED         = 0x08000000
	WS_VISIBLE          = 0x10000000 // Initially visible
	WS_MINIMIZE         = 0x20000000 // Initially minimized
	WS_CHILDWINDOW      = 0x40000000
	WS_POPUP            = 0x80000000
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_POPUPWINDOW      = WS_POPUP | WS_BORDER | WS_SYSMENU
)

// Constants that may be passed as the modeNum param to EnumDisplaySettingsExW
const (
	ENUM_CURRENT_SETTINGS  = 0xFFFFFFFF
	ENUM_REGISTRY_SETTINGS = 0xFFFFFFFE
)

// Constants that may be passed as the flags param to EnumDisplaySettingsExW
const (
	EDS_ROTATEDMODE = 0x00000002
	EDS_RAWMODE     = 0x00000004
)

// Constants that may be passed as the index param to GetDeviceCaps
const (
	HORZRES        = 8
	VERTRES        = 10
	LOGPIXELSX     = 88
	LOGPIXELSY     = 90
	SCALINGFACTORX = 114
	SCALINGFACTORY = 115
)

// Constants returned in the Flags of MONITORINFO
const (
	MONITORINFOF_PRIMARY = 0x00000001
)

// https://docs.microsoft.com/en-us/windows/desktop/api/shellscalingapi/ne-shellscalingapi-monitor_dpi_type
const (
	MDT_EFFECTIVE_DPI = iota
	MDT_ANGULAR_DPI
	MDT_RAW_DPI
	MDT_DEFAULT
)

// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-getsystemmetrics
const (
	SM_CXSCREEN                    = 0
	SM_CYSCREEN                    = 1
	SM_CXVSCROLL                   = 2
	SM_CYHSCROLL                   = 3
	SM_CYCAPTION                   = 4
	SM_CXBORDER                    = 5
	SM_CYBORDER                    = 6
	SM_CXDLGFRAME                  = 7
	SM_CYDLGFRAME                  = 8
	SM_CXFIXEDFRAME                = SM_CXDLGFRAME
	SM_CYFIXEDFRAME                = SM_CYDLGFRAME
	SM_CYVTHUMB                    = 9
	SM_CXHTHUMB                    = 10
	SM_CXICON                      = 11
	SM_CYICON                      = 12
	SM_CXCURSOR                    = 13
	SM_CYCURSOR                    = 14
	SM_CYMENU                      = 15
	SM_CXFULLSCREEN                = 16
	SM_CYFULLSCREEN                = 17
	SM_CYKANJIWINDOW               = 18
	SM_MOUSEPRESENT                = 19
	SM_CYVSCROLL                   = 20
	SM_CXHSCROLL                   = 21
	SM_DEBUG                       = 22
	SM_SWAPBUTTON                  = 23
	SM_CXMIN                       = 28
	SM_CYMIN                       = 29
	SM_CXSIZE                      = 30
	SM_CYSIZE                      = 31
	SM_CXFRAME                     = 32
	SM_CXSIZEFRAME                 = SM_CXFRAME
	SM_CYFRAME                     = 33
	SM_CYSIZEFRAME                 = SM_CYFRAME
	SM_CXMINTRACK                  = 34
	SM_CYMINTRACK                  = 35
	SM_CXDOUBLECLK                 = 36
	SM_CYDOUBLECLK                 = 37
	SM_CXICONSPACING               = 38
	SM_CYICONSPACING               = 39
	SM_MENUDROPALIGNMENT           = 40
	SM_PENWINDOWS                  = 41
	SM_DBCSENABLED                 = 42
	SM_CMOUSEBUTTONS               = 43
	SM_SECURE                      = 44
	SM_CXEDGE                      = 45
	SM_CYEDGE                      = 46
	SM_CXMINSPACING                = 47
	SM_CYMINSPACING                = 48
	SM_CXSMICON                    = 49
	SM_CYSMICON                    = 50
	SM_CYSMCAPTION                 = 51
	SM_CXSMSIZE                    = 52
	SM_CYSMSIZE                    = 53
	SM_CXMENUSIZE                  = 54
	SM_CYMENUSIZE                  = 55
	SM_ARRANGE                     = 56
	SM_CXMINIMIZED                 = 57
	SM_CYMINIMIZED                 = 58
	SM_CXMAXTRACK                  = 59
	SM_CYMAXTRACK                  = 60
	SM_CXMAXIMIZED                 = 61
	SM_CYMAXIMIZED                 = 62
	SM_NETWORK                     = 63
	SM_CLEANBOOT                   = 67
	SM_CXDRAG                      = 68
	SM_CYDRAG                      = 69
	SM_SHOWSOUNDS                  = 70
	SM_CXMENUCHECK                 = 71
	SM_CYMENUCHECK                 = 72
	SM_SLOWMACHINE                 = 73
	SM_MIDEASTENABLED              = 74
	SM_MOUSEWHEELPRESENT           = 75
	SM_XVIRTUALSCREEN              = 76
	SM_YVIRTUALSCREEN              = 77
	SM_CXVIRTUALSCREEN             = 78
	SM_CYVIRTUALSCREEN             = 79
	SM_CMONITORS                   = 80
	SM_SAMEDISPLAYFORMAT           = 81
	SM_IMMENABLED                  = 82
	SM_CXFOCUSBORDER               = 83
	SM_CYFOCUSBORDER               = 84
	SM_TABLETPC                    = 86
	SM_MEDIACENTER                 = 87
	SM_STARTER                     = 88
	SM_SERVERR2                    = 89
	SM_MOUSEHORIZONTALWHEELPRESENT = 91
	SM_CXPADDEDBORDER              = 92
	SM_DIGITIZER                   = 94
	SM_MAXIMUMTOUCHES              = 95
	SM_REMOTESESSION               = 0x1000
	SM_SHUTTINGDOWN                = 0x2000
	SM_REMOTECONTROL               = 0x2001
	SM_CONVERTIBLESLATEMODE        = 0x2003
	SM_SYSTEMDOCKED                = 0x2004
)
