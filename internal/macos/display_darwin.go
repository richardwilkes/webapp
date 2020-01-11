// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package macos

import (
	// #import "displays.h"
	"C"
)

import (
	"unsafe"

	"github.com/richardwilkes/webapp"
)

func (d *driver) Displays() []*webapp.Display {
	var count C.ulong
	ptr := unsafe.Pointer(C.displays(&count))
	displays := (*[1 << 30]C.Display)(ptr)
	result := make([]*webapp.Display, count)
	for i := range result {
		dsp := &webapp.Display{}
		dsp.Bounds.X = float64(displays[i].bounds.origin.x)
		dsp.Bounds.Y = float64(displays[i].bounds.origin.y)
		dsp.Bounds.Width = float64(displays[i].bounds.size.width)
		dsp.Bounds.Height = float64(displays[i].bounds.size.height)
		dsp.UsableBounds.X = float64(displays[i].usableBounds.origin.x)
		dsp.UsableBounds.Y = float64(displays[i].usableBounds.origin.y)
		dsp.UsableBounds.Width = float64(displays[i].usableBounds.size.width)
		dsp.UsableBounds.Height = float64(displays[i].usableBounds.size.height)
		dsp.ScalingFactor = float64(displays[i].scalingFactor)
		dsp.IsMain = displays[i].isMain != 0
		result[i] = dsp
	}
	C.free(ptr)
	return result
}
