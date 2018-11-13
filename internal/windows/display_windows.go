package windows

// Constants that may be passed as the modeNum param to EnumDisplaySettingsExW
const (
	ENUM_CURRENT_SETTINGS  = 0xFFFFFFFF
	ENUM_REGISTRY_SETTINGS = 0xFFFFFFFE
)

// Constatns that may be passed as the flags param to EnumDisplaySettingsExW
const (
	EDS_ROTATEDMODE = 0x00000002
	EDS_RAWMODE     = 0x00000004
)

// Constatns that may be passed as the index param to GetDeviceCaps
const (
	HORZRES        = 8
	VERTRES        = 10
	LOGPIXELSX     = 88
	LOGPIXELSY     = 90
	SCALINGFACTORX = 114
	SCALINGFACTORY = 115
)

// DISPLAY_DEVICEW is defined here:
// https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-_display_devicew
type DISPLAY_DEVICEW struct {
	Size         uint32
	DeviceName   [32]uint16
	DeviceString [128]uint16
	Flags        uint32
	DeviceID     [128]uint16
	DeviceKey    [128]uint16
}

// DEVMODEW is defined here:
// https://docs.microsoft.com/en-us/windows/desktop/api/wingdi/ns-wingdi-_devicemodew
type DEVMODEW struct {
	DeviceName    [32]uint16
	SpecVersion   uint16
	DriverVersion uint16
	Size          uint16
	DriverExtra   uint16
	Fields        uint32
	X             int32
	Y             int32
	Orientation   uint32
	FixedOutput   uint32
	Color         int16
	Duplex        int16
	YResolution   int16
	TTOption      int16
	Collate       int16
	FormName      [32]uint16
	LogPixels     uint16
	BitsPerPixel  uint32
	PelsWidth     uint32
	PelsHeight    uint32
	Flags         uint32
	Frequency     uint32
	ICMMethod     uint32
	ICMIntent     uint32
	MediaType     uint32
	DitherType    uint32
	Reserved1     uint32
	Reserved2     uint32
	PanningWidth  uint32
	PanningHeight uint32
}
