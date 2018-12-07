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
