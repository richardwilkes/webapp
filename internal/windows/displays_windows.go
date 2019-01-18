package windows

import (
	"unsafe"

	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/webapp"
)

func (d *driver) Displays() []*webapp.Display {
	result := make([]*webapp.Display, 0)
	if err := EnumDisplayMonitors(0, nil, func(monitor HMONITOR, dc HDC, rect *RECT, param LPARAM) BOOL {
		d := &webapp.Display{}
		var info MONITORINFO
		info.Size = DWORD(unsafe.Sizeof(info))
		if err := GetMonitorInfoW(monitor, &info); err != nil {
			jot.Error(err)
		} else {
			var dpiX, dpiY uint32
			if err = GetDpiForMonitor(monitor, MDT_EFFECTIVE_DPI, &dpiX, &dpiY); err != nil {
				// Windows 7 fallback
				overallX := GetDeviceCaps(dc, LOGPIXELSX)
				overallY := GetDeviceCaps(dc, LOGPIXELSY)

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
			d.IsMain = info.Flags&MONITORINFOF_PRIMARY != 0
			result = append(result, d)
		}
		return 1
	}, 0); err != nil {
		jot.Error(err)
	}
	return result
}
