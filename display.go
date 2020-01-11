// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package webapp

import (
	"github.com/richardwilkes/toolbox/xmath/geom"
)

// Display holds information about each available active display.
type Display struct {
	Bounds        geom.Rect
	UsableBounds  geom.Rect
	ScalingFactor float64
	IsMain        bool
}

// MainDisplay returns the primary display.
func MainDisplay() *Display {
	for _, d := range Displays() {
		if d.IsMain {
			return d
		}
	}
	return nil
}

// Displays returns all displays.
func Displays() []*Display {
	return driver.Displays()
}
