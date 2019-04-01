package windows

import (
	"unsafe"

	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/win32"
)

func (d *driver) Displays() []*webapp.Display {
	result := make([]*webapp.Display, 0)
	win32.EnumDisplayMonitors(0, nil, func(monitor win32.HMONITOR, dc win32.HDC, rect *win32.RECT, param win32.LPARAM) win32.BOOL {
		d := &webapp.Display{}
		var info win32.MONITORINFO
		info.Size = win32.DWORD(unsafe.Sizeof(info))
		if win32.GetMonitorInfo(monitor, &info) {
			var dpiX, dpiY uint32
			if !win32.GetDpiForMonitor(monitor, win32.MDT_EFFECTIVE_DPI, &dpiX, &dpiY) {
				// Windows 7 fallback
				overallX := win32.GetDeviceCaps(dc, win32.LOGPIXELSX)
				overallY := win32.GetDeviceCaps(dc, win32.LOGPIXELSY)
				if overallX > 0 && overallY > 0 {
					dpiX = uint32(overallX)
					dpiY = uint32(overallY)
				}
			}
			d.Bounds.X = float64(info.MonitorBounds.Left)
			d.Bounds.Y = float64(info.MonitorBounds.Top)
			d.Bounds.Width = float64(info.MonitorBounds.Right - info.MonitorBounds.Left)
			d.Bounds.Height = float64(info.MonitorBounds.Bottom - info.MonitorBounds.Top)
			d.UsableBounds.X = float64(info.WorkBounds.Left)
			d.UsableBounds.Y = float64(info.WorkBounds.Top)
			d.UsableBounds.Width = float64(info.WorkBounds.Right - info.WorkBounds.Left)
			d.UsableBounds.Height = float64(info.WorkBounds.Bottom - info.WorkBounds.Top)
			d.ScalingFactor = float64(dpiX) / 96
			d.IsMain = info.Flags&win32.MONITORINFOF_PRIMARY != 0
			result = append(result, d)
		}
		return 1
	}, 0)
	return result
}
