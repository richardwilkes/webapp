package dpi

// From https://docs.microsoft.com/en-us/windows/desktop/api/windef/ne-windef-dpi_awareness
const (
	AWARENESS_INVALID = iota
	AWARENESS_UNAWARE
	AWARENESS_SYSTEM_AWARE
	AWARENESS_PER_MONITOR_AWARE
)

// From https://docs.microsoft.com/en-us/windows/desktop/hidpi/dpi-awareness-context
const (
	AWARENESS_CONTEXT_UNAWARE              = 0xFFFFFFFF
	AWARENESS_CONTEXT_SYSTEM_AWARE         = 0xFFFFFFFE
	AWARENESS_CONTEXT_PER_MONITOR_AWARE    = 0xFFFFFFFD
	AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2 = 0xFFFFFFFC
	AWARENESS_CONTEXT_UNAWARE_GDISCALED    = 0xFFFFFFFB
)
