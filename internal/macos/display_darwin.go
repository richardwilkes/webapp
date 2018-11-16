package macos

import (
	// #import "displays.h"
	"C"
	"unsafe"

	"github.com/richardwilkes/webapp"
)

func (d *driver) Displays() []*webapp.Display {
	var count C.ulong
	ptr := unsafe.Pointer(C.displays(&count))
	displays := (*[1 << 30]C.Display)(ptr)
	result := make([]*webapp.Display, count)
	for i := range result {
		d := &webapp.Display{}
		d.Bounds.X = float64(displays[i].bounds.origin.x)
		d.Bounds.Y = float64(displays[i].bounds.origin.y)
		d.Bounds.Width = float64(displays[i].bounds.size.width)
		d.Bounds.Height = float64(displays[i].bounds.size.height)
		d.UsableBounds.X = float64(displays[i].usableBounds.origin.x)
		d.UsableBounds.Y = float64(displays[i].usableBounds.origin.y)
		d.UsableBounds.Width = float64(displays[i].usableBounds.size.width)
		d.UsableBounds.Height = float64(displays[i].usableBounds.size.height)
		d.IsMain = displays[i].isMain != 0
		result[i] = d
	}
	C.free(ptr)
	return result
}
